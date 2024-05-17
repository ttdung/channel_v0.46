package channel_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "github.com/ttdung/channel_v0.46/testutil/keeper"
	"github.com/ttdung/channel_v0.46/testutil/nullify"
	"github.com/ttdung/channel_v0.46/x/channel"
	"github.com/ttdung/channel_v0.46/x/channel/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		ChannelList: []types.Channel{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		CommitmentList: []types.Commitment{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		FwdcommitmentList: []types.Fwdcommitment{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.ChannelKeeper(t)
	channel.InitGenesis(ctx, *k, genesisState)
	got := channel.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.ChannelList, got.ChannelList)
	require.ElementsMatch(t, genesisState.CommitmentList, got.CommitmentList)
	require.ElementsMatch(t, genesisState.FwdcommitmentList, got.FwdcommitmentList)
	// this line is used by starport scaffolding # genesis/test/assert
}
