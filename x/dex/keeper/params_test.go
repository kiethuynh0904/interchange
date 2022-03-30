package keeper_test

import (
	"testing"

	testkeeper "github.com/kiethuynh0904/interchange/testutil/keeper"
	"github.com/kiethuynh0904/interchange/x/dex/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.DexKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
