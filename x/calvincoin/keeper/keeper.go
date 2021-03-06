package keeper

import (
	"fmt"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/calvinlauco/cosmos-sdk-calvincoin/x/calvincoin/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Keeper of the calvincoin store
type Keeper struct {
	CoinKeeper   types.BankKeeper
	SupplyKeeper types.SupplyKeeper

	storeKey sdk.StoreKey
	cdc      *codec.Codec
}

// NewKeeper creates a calvincoin keeper
func NewKeeper(cdc *codec.Codec, key sdk.StoreKey, coinKeeper types.BankKeeper, supplyKeeper types.SupplyKeeper) Keeper {
	keeper := Keeper{
		CoinKeeper:   coinKeeper,
		SupplyKeeper: supplyKeeper,

		storeKey: key,
		cdc:      cdc,
	}
	return keeper
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// Get returns the pubkey from the adddress-pubkey relation
// func (k Keeper) Get(ctx sdk.Context, key string) (/* TODO: Fill out this type */, error) {
// 	store := ctx.KVStore(k.storeKey)
// 	var item /* TODO: Fill out this type */
// 	byteKey := []byte(key)
// 	err := k.cdc.UnmarshalBinaryLengthPrefixed(store.Get(byteKey), &item)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return item, nil
// }

// func (k Keeper) set(ctx sdk.Context, key string, value /* TODO: fill out this type */ ) {
// 	store := ctx.KVStore(k.storeKey)
// 	bz := k.cdc.MustMarshalBinaryLengthPrefixed(value)
// 	store.Set([]byte(key), bz)
// }

// func (k Keeper) delete(ctx sdk.Context, key string) {
// 	store := ctx.KVStore(k.storeKey)
// 	store.Delete([]byte(key))
// }

// func (k Keeper) Transfer(ctx sdk.Context) {
// 	k.Transfer(ctx)
// }
