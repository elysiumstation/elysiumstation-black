package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/elysiumstation/blackfury/testutil/keeper"
	"github.com/elysiumstation/blackfury/x/voter/keeper"
	"github.com/elysiumstation/blackfury/x/voter/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.VoterKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
