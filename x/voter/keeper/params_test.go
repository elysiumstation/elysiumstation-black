package keeper_test

import (
	"testing"

	testkeeper "github.com/furya-official/blackfury/testutil/keeper"
	"github.com/furya-official/blackfury/x/voter/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.VoterKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
