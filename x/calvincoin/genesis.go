package calvincoin

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type GenesisState struct {
}

func NewGenesisState() GenesisState {
	return GenesisState{}
}

func ValidateGenesis(_data GenesisState) error {
	return nil
}

func DefaultGenesisState() GenesisState {
	return GenesisState{}
}

// InitGenesis initialize default parameters
// and the keeper's address to pubkey map
func InitGenesis(ctx sdk.Context, k Keeper, data GenesisState) {
	// TODO: Define logic for when you would like to initialize a new genesis
}

// ExportGenesis writes the current store values
// to a genesis file, which can be imported again
// with InitGenesis
func ExportGenesis(ctx sdk.Context, k Keeper) (data GenesisState) {
	// TODO: Define logic for exporting state
	return GenesisState{}
}
