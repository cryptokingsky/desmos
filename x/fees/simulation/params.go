package simulation

// DONTCOVER

import (
	"fmt"
	"math/rand"

	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/desmos-labs/desmos/x/fees/types"
)

func ParamChanges(r *rand.Rand) []simulation.ParamChange {
	params := types.DefaultParams()
	return []simulation.ParamChange{
		simulation.NewSimParamChange(types.ModuleName, string(types.FeeDenomStoreKey),
			func(r *rand.Rand) string {
				return fmt.Sprintf(`{"fee_denom":"%s"`,
					params.FeeDenom)
			},
		),
		simulation.NewSimParamChange(types.ModuleName, string(types.RequiredFeeStoreKey),
			func(r *rand.Rand) string {
				return fmt.Sprintf(`{"required_fee":"%s"`,
					params.RequiredFee)
			},
		),
	}
}