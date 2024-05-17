package keeper

import (
	"context"
	"errors"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ttdung/channel_v0.46/x/channel/types"
)

func (k msgServer) WithdrawTimelock(goCtx context.Context, msg *types.MsgWithdrawTimelock) (*types.MsgWithdrawTimelockResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	val, found := k.Keeper.GetCommitment(ctx, msg.Index)
	if !found {
		return nil, fmt.Errorf("Commitment with index %v is not existing", msg.Index)
	}

	if val.Partneraddr != msg.To {
		return nil, errors.New("not matching receiver address!")
	}

	if val.Numblock > uint64(ctx.BlockHeight()) {
		return nil, errors.New("wait until valid block height")
	}

	to, err := sdk.AccAddressFromBech32(msg.To)
	if err != nil {
		return nil, err
	}

	if k.bankKeeper.BlockedAddr(to) {
		err = fmt.Errorf("%s is not allowed to receive funds", msg.To)
	} else {
		for _, coin := range val.Cointohtlc {
			if coin.Amount.IsPositive() {
				err = k.Keeper.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, to, sdk.Coins{*coin})
				if err != nil {
					return nil, err
				}
			}
		}
	}

	k.Keeper.RemoveCommitment(ctx, msg.Index)

	return &types.MsgWithdrawTimelockResponse{}, nil
}
