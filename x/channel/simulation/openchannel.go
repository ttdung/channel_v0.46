package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/ttdung/channel_v0.46/x/channel/keeper"
	"github.com/ttdung/channel_v0.46/x/channel/types"
)

func SimulateMsgOpenchannel(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgOpenchannel{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the Openchannel simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "Openchannel simulation not implemented"), nil, nil
	}
}
