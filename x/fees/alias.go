package fees

// autogenerated code using github.com/haasted/alias-generator.
// based on functionality in github.com/rigelrozanski/multitool

import (
	"github.com/desmos-labs/desmos/x/fees/client/cli"
	"github.com/desmos-labs/desmos/x/fees/client/rest"
	"github.com/desmos-labs/desmos/x/fees/keeper"
	"github.com/desmos-labs/desmos/x/fees/types"
)

const (
	DefaultParamspace = types.DefaultParamspace
	ModuleName        = types.ModuleName
	RouterKey         = types.RouterKey
	QuerierRoute      = types.QuerierRoute
	QueryParams       = types.QueryParams
)

var (
	// functions aliases
	GetQueryCmd              = cli.GetQueryCmd
	GetCmdQueryFeesParams    = cli.GetCmdQueryFeesParams
	RegisterRoutes           = rest.RegisterRoutes
	NewQuerier               = keeper.NewQuerier
	NewKeeper                = keeper.NewKeeper
	NewGenesisState          = types.NewGenesisState
	DefaultGenesisState      = types.DefaultGenesisState
	ValidateGenesis          = types.ValidateGenesis
	RegisterCodec            = types.RegisterCodec
	ParamKeyTable            = types.ParamKeyTable
	NewParams                = types.NewParams
	DefaultParams            = types.DefaultParams
	ValidateFeeDenomParam    = types.ValidateFeeDenomParam
	ValidateRequiredFeeParam = types.ValidateRequiredFeeParam

	// variable aliases
	ModuleCdc           = types.ModuleCdc
	DefaultFeeDenom     = types.DefaultFeeDenom
	DefaultRequiredFee  = types.DefaultRequiredFee
	FeeDenomStoreKey    = types.FeeDenomStoreKey
	RequiredFeeStoreKey = types.RequiredFeeStoreKey
)

type (
	Keeper       = keeper.Keeper
	GenesisState = types.GenesisState
	Params       = types.Params
)