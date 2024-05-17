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

func createNFwdcommitment(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Fwdcommitment {
	items := make([]types.Fwdcommitment, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetFwdcommitment(ctx, items[i])
	}
	return items
}

func TestFwdcommitmentGet(t *testing.T) {
	keeper, ctx := keepertest.ChannelKeeper(t)
	items := createNFwdcommitment(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetFwdcommitment(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestFwdcommitmentRemove(t *testing.T) {
	keeper, ctx := keepertest.ChannelKeeper(t)
	items := createNFwdcommitment(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveFwdcommitment(ctx,
			item.Index,
		)
		_, found := keeper.GetFwdcommitment(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestFwdcommitmentGetAll(t *testing.T) {
	keeper, ctx := keepertest.ChannelKeeper(t)
	items := createNFwdcommitment(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllFwdcommitment(ctx)),
	)
}
