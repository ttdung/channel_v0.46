package keeper

import (
	"context"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ttdung/channel_v0.46/x/channel/types"
)

func (k msgServer) Closechannel(goCtx context.Context, msg *types.MsgClosechannel) (*types.MsgClosechannelResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	for _, coin := range msg.CoinA {
		if err := k.bankKeeper.IsSendEnabledCoins(ctx, *coin); err != nil {
			return nil, err
		}
	}

	from, err := sdk.AccAddressFromBech32(msg.MultisigAddr)
	if err != nil {
		return nil, err
	}
	toA, err := sdk.AccAddressFromBech32(msg.PartA)
	if err != nil {
		return nil, err
	}
	toB, err := sdk.AccAddressFromBech32(msg.PartB)
	if err != nil {
		return nil, err
	}

	// todo: Check from_amount > coinA + coin B

	if k.bankKeeper.BlockedAddr(toA) {
		err = sdkerrors.Wrapf(sdkerrors.ErrUnauthorized, "%s is not allowed to receive funds", msg.PartA)
	} else {
		for _, coin := range msg.CoinA {
			if coin.Amount.IsPositive() {
				err = k.bankKeeper.SendCoins(ctx, from, toA, sdk.Coins{*coin})
			}
		}
	}

	if k.bankKeeper.BlockedAddr(toB) {
		err = sdkerrors.Wrapf(sdkerrors.ErrUnauthorized, "%s is not allowed to receive funds", msg.PartB)
	} else {
		for _, coin := range msg.CoinB {
			if coin.Amount.IsPositive() {
				err = k.bankKeeper.SendCoins(ctx, from, toB, sdk.Coins{*coin})
			}
		}
	}

	if err != nil {
		return nil, err
	}

	k.Keeper.RemoveChannel(ctx, msg.Channelid)

	return &types.MsgClosechannelResponse{}, nil
}
