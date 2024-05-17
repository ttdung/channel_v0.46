package keeper_test

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "github.com/ttdung/channel_v0.46/testutil/keeper"
	"github.com/ttdung/channel_v0.46/x/channel/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.ChannelKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)
	hash := sha256.Sum256([]byte("bcdf"))
	fmt.Println("hashcode:", base64.StdEncoding.EncodeToString(hash[:]))

	require.EqualValues(t, params, k.GetParams(ctx))
}
