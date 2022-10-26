package labrador_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "labrador/testutil/keeper"
	"labrador/testutil/nullify"
	"labrador/x/labrador"
	"labrador/x/labrador/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.LabradorKeeper(t)
	labrador.InitGenesis(ctx, *k, genesisState)
	got := labrador.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
