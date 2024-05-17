package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgOpenchannel{}, "github.com/ttdung/channel_v0.46/Openchannel", nil)
	cdc.RegisterConcrete(&MsgClosechannel{}, "github.com/ttdung/channel_v0.46/Closechannel", nil)
	cdc.RegisterConcrete(&MsgCommitment{}, "github.com/ttdung/channel_v0.46/Commitment", nil)
	cdc.RegisterConcrete(&MsgWithdrawTimelock{}, "github.com/ttdung/channel_v0.46/WithdrawTimelock", nil)
	cdc.RegisterConcrete(&MsgWithdrawHashlock{}, "github.com/ttdung/channel_v0.46/WithdrawHashlock", nil)
	cdc.RegisterConcrete(&MsgFund{}, "github.com/ttdung/channel_v0.46/Fund", nil)
	cdc.RegisterConcrete(&MsgAcceptfund{}, "github.com/ttdung/channel_v0.46/Acceptfund", nil)
	cdc.RegisterConcrete(&MsgSendercommit{}, "github.com/ttdung/channel_v0.46/Sendercommit", nil)
	cdc.RegisterConcrete(&MsgReceivercommit{}, "github.com/ttdung/channel_v0.46/Receivercommit", nil)
	cdc.RegisterConcrete(&MsgSenderwithdrawtimelock{}, "github.com/ttdung/channel_v0.46/Senderwithdrawtimelock", nil)
	cdc.RegisterConcrete(&MsgSenderwithdrawhashlock{}, "github.com/ttdung/channel_v0.46/Senderwithdrawhashlock", nil)
	cdc.RegisterConcrete(&MsgReceiverwithdraw{}, "github.com/ttdung/channel_v0.46/Receiverwithdraw", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgOpenchannel{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgClosechannel{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCommitment{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgWithdrawTimelock{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgWithdrawHashlock{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgFund{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgAcceptfund{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSendercommit{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgReceivercommit{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSenderwithdrawtimelock{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSenderwithdrawhashlock{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgReceiverwithdraw{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
