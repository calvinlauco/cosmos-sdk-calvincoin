package keeper

import (
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/calvinlauco/cosmos-sdk-calvincoin/x/calvincoin/types"
)

const (
	QueryBalances    = "balances"
	QueryTotalSupply = "totalsupply"
)

// NewQuerier creates a new querier for calvincoin clients.
func NewQuerier(k Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) ([]byte, error) {
		switch path[0] {
		// case types.QueryParams:
		// 	return queryParams(ctx, k)
		// TODO: Put the modules query routes
		case QueryBalances:
			return queryBalances(ctx, path[1:], req, k)
		case QueryTotalSupply:
			return queryTotalSupply(ctx, path[1:], req, k)
		default:
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "unknown calvincoin query endpoint")
		}
	}
}

// func queryParams(ctx sdk.Context, k Keeper) ([]byte, error) {
// 	params := k.GetParams(ctx)

// 	res, err := codec.MarshalJSONIndent(types.ModuleCdc, params)
// 	if err != nil {
// 		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
// 	}

// 	return res, nil
// }

// TODO: Add the modules query functions
// They will be similar to the above one: queryParams()

func queryBalances(ctx sdk.Context, path []string, req abci.RequestQuery, keeper Keeper) ([]byte, error) {
	var err error

	var addr sdk.AccAddress
	addr, err = sdk.AccAddressFromBech32(path[0])
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, err.Error())
	}

	balances := keeper.CoinKeeper.GetCoins(ctx, addr)
	res, err := codec.MarshalJSONIndent(keeper.cdc, types.QueryResResolve{Value: balances.String()})
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

func queryTotalSupply(ctx sdk.Context, path []string, req abci.RequestQuery, keeper Keeper) ([]byte, error) {
	var err error

	totalSupply := keeper.SupplyKeeper.GetSupply(ctx).GetTotal()

	res, err := codec.MarshalJSONIndent(keeper.cdc, types.QueryResResolve{Value: totalSupply.String()})
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}
