package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
	"github.com/ttdung/channel_v0.46/testutil/sample"
)

func TestMsgSenderwithdrawhashlock_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgSenderwithdrawhashlock
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgSenderwithdrawhashlock{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgSenderwithdrawhashlock{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
