package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	keepertest "github.com/ttdung/channel_v0.46/testutil/keeper"
	"github.com/ttdung/channel_v0.46/testutil/nullify"
	"github.com/ttdung/channel_v0.46/x/channel/keeper"
	"github.com/ttdung/channel_v0.46/x/channel/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNCommitment(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Commitment {
	items := make([]types.Commitment, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetCommitment(ctx, items[i])
	}
	return items
}

func TestCommitmentGet(t *testing.T) {
	keeper, ctx := keepertest.ChannelKeeper(t)
	items := createNCommitment(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetCommitment(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestCommitmentRemove(t *testing.T) {
	keeper, ctx := keepertest.ChannelKeeper(t)
	items := createNCommitment(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveCommitment(ctx,
			item.Index,
		)
		_, found := keeper.GetCommitment(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestCommitmentGetAll(t *testing.T) {
	keeper, ctx := keepertest.ChannelKeeper(t)
	items := createNCommitment(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllCommitment(ctx)),
	)
}
