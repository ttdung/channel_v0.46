package keeper

import (
	"context"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ttdung/channel_v0.46/x/channel/types"
)

func (k msgServer) Senderwithdrawtimelock(goCtx context.Context, msg *types.MsgSenderwithdrawtimelock) (*types.MsgSenderwithdrawtimelockResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	to, err := sdk.AccAddressFromBech32(msg.To)
	if err != nil {
		return nil, err
	}

	val, found := k.Keeper.GetFwdcommitment(ctx, msg.Transferindex)
	if !found {
		return nil, fmt.Errorf("Transferindex %v is not existing", msg.Transferindex)
	}

	if val.SenderAddr != msg.To {
		return nil, fmt.Errorf("Not matching address!")
	}

	if val.Timelocksender > uint64(ctx.BlockHeight()) {
		return nil, fmt.Errorf("Wait until valid blockheight, expected height %v", ctx.BlockHeight())
	}

	if k.bankKeeper.BlockedAddr(to) {
		err = fmt.Errorf("%s is not allowed to receive funds", msg.To)
	} else {
		for _, coin := range val.Cointransfer {
			if coin.Amount.IsPositive() {
				err = k.Keeper.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, to, sdk.Coins{*coin})
				if err != nil {
					return nil, err
				}
			}
		}
	}

	k.Keeper.RemoveFwdcommitment(ctx, msg.Transferindex)

	return &types.MsgSenderwithdrawtimelockResponse{}, nil
}
