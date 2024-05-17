package keeper

import (
	"context"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ttdung/channel_v0.46/x/channel/types"
)

func (k msgServer) Sendercommit(goCtx context.Context, msg *types.MsgSendercommit) (*types.MsgSendercommitResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	val, found := k.Keeper.GetChannel(ctx, msg.Channelid)
	if !found {
		return nil, fmt.Errorf(fmt.Sprintf("ChannelID %v is not existing", msg.Channelid))
	}

	var partnerAddr, creatorAddr string
	if msg.SenderAddr == val.PartA {
		partnerAddr = val.PartB
		creatorAddr = val.PartA
	} else {
		partnerAddr = val.PartA
		creatorAddr = val.PartB
	}

	toSender, err := sdk.AccAddressFromBech32(msg.SenderAddr)
	if err != nil {
		return nil, err
	}

	fromMultisig, err := sdk.AccAddressFromBech32(msg.MultisigAddr)
	if err != nil {
		return nil, err
	}

	// Send coin to creator of commitment
	for _, coin := range msg.Cointosender {
		if coin.Amount.IsPositive() {
			err = k.bankKeeper.SendCoins(ctx, fromMultisig, toSender, sdk.Coins{*coin})
			if err != nil {
				return nil, err
			}
		}
	}

	// Send to HTLC
	htlcIndex := ""
	CointoHTLC := msg.Cointohtlc
	for _, coin := range CointoHTLC {
		if coin.Amount.IsPositive() {
			err = k.Keeper.bankKeeper.SendCoinsFromAccountToModule(ctx, fromMultisig, types.ModuleName, sdk.Coins{*coin})
			if err != nil {
				return nil, err
			}
		}
	}

	htlcIndex = fmt.Sprintf("%s:%s", msg.Channelid, msg.Hashcodehtlc)

	unlockBlockHtlc := msg.Timelockhtlc + uint64(ctx.BlockHeight())

	commitment := types.Commitment{
		Index:         htlcIndex,
		MultisigAddr:  msg.MultisigAddr,
		Creatoraddr:   creatorAddr,
		Partneraddr:   partnerAddr,
		Hashcode:      msg.Hashcodehtlc,
		Numblock:      unlockBlockHtlc,
		Cointocreator: msg.Cointosender,
		Cointohtlc:    msg.Cointohtlc,
		Channelid:     msg.Channelid,
	}

	k.Keeper.SetCommitment(ctx, commitment)

	// Send to FwdContract
	CointoFC := msg.Cointransfer
	for _, coin := range CointoFC {
		if coin.Amount.IsPositive() {
			err = k.Keeper.bankKeeper.SendCoinsFromAccountToModule(ctx, fromMultisig, types.ModuleName, sdk.Coins{*coin})
			if err != nil {
				return nil, err
			}
		}
	}

	creator := "sender"
	transferIndex := fmt.Sprintf("%s:%s:%s", msg.Channelid, msg.Hashcodedest, creator)

	Timelocksender := msg.Timelocksender

	Timelockreceiver := msg.Timelockreceiver
	Timelockreceiver = Timelockreceiver + uint64(ctx.BlockHeight())

	fwscommitment := types.Fwdcommitment{
		Index:            transferIndex,
		Channelid:        msg.Channelid,
		MultisigAddr:     msg.MultisigAddr,
		SenderAddr:       creatorAddr,
		ReceiverAddr:     partnerAddr,
		Timelockreceiver: Timelockreceiver,
		Timelocksender:   Timelocksender,
		Hashcodehtlc:     msg.Hashcodehtlc,
		Hashcodedest:     msg.Hashcodedest,
		Cointransfer:     msg.Cointransfer,
		Creator:          creator,
	}

	k.Keeper.SetFwdcommitment(ctx, fwscommitment)

	k.Keeper.RemoveChannel(ctx, msg.Channelid)

	return &types.MsgSendercommitResponse{
		HtlcIndex:     htlcIndex,
		TransferIndex: transferIndex,
	}, nil
}
