package types

import (
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
)

// nolint
const (
	// DisplayDenom defines the denomination displayed to users in client applications.
	DisplayDenom = "fury"
	// BaseDenom defines to the default denomination used in Blackfury (staking, EVM, governance, etc.)
	BaseDenom = AttoFuryDenom

	AttoFuryDenom = "afury" // 1e-18
	MicroFUSDDenom = "ufusd"  // 1e-6
)

var (
	// MicroFUSDTarget defines the target exchange rate of ufusd denominated in uUSD.
	MicroFUSDTarget = sdk.OneDec()
)

func SetDenomMetaDataForStableCoins(ctx sdk.Context, k bankkeeper.Keeper) {
	for _, base := range []string{MicroFUSDDenom} {
		if _, ok := k.GetDenomMetaData(ctx, base); ok {
			continue
		}

		display := base[1:] // e.g., fusd
		// Register meta data to bank module
		k.SetDenomMetaData(ctx, banktypes.Metadata{
			Description: "The native stable token of the Blackfury.",
			DenomUnits: []*banktypes.DenomUnit{
				{Denom: "u" + display, Exponent: uint32(0), Aliases: []string{"micro" + display}}, // e.g., ufusd
				{Denom: "m" + display, Exponent: uint32(3), Aliases: []string{"milli" + display}}, // e.g., musm
				{Denom: display, Exponent: uint32(6), Aliases: []string{}},                        // e.g., fusd
			},
			Base:    base,
			Display: display,
			Name:    strings.ToUpper(display), // e.g., FUSD
			Symbol:  strings.ToUpper(display), // e.g., FUSD
		})
	}
}
