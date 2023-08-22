package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgOpenchannel{}, "channel/Openchannel", nil)
	cdc.RegisterConcrete(&MsgClosechannel{}, "channel/Closechannel", nil)
	cdc.RegisterConcrete(&MsgCommitment{}, "channel/Commitment", nil)
	cdc.RegisterConcrete(&MsgWithdrawTimelock{}, "channel/WithdrawTimelock", nil)
	cdc.RegisterConcrete(&MsgWithdrawHashlock{}, "channel/WithdrawHashlock", nil)
	cdc.RegisterConcrete(&MsgFund{}, "channel/Fund", nil)
	cdc.RegisterConcrete(&MsgAcceptfund{}, "channel/Acceptfund", nil)
	cdc.RegisterConcrete(&MsgSendercommit{}, "channel/Sendercommit", nil)
	cdc.RegisterConcrete(&MsgReceivercommit{}, "channel/Receivercommit", nil)
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
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
