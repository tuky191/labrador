package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "labrador/testutil/keeper"
	"labrador/x/labrador/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.LabradorKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
