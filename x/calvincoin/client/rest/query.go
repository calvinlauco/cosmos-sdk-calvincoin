package rest

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/types/rest"
)

// func registerQueryRoutes(cliCtx context.CLIContext, r *mux.Router) {
// 	// TODO: Define your GET REST endpoints
// 	r.HandleFunc(
// 		"/calvincoin/parameters",
// 		queryParamsHandlerFn(cliCtx),
// 	).Methods("GET")
// }

// func queryParamsHandlerFn(cliCtx context.CLIContext) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		cliCtx, ok := rest.ParseQueryHeightOrReturnBadRequest(w, cliCtx, r)
// 		if !ok {
// 			return
// 		}

// 		route := fmt.Sprintf("custom/%s/parameters", types.QuerierRoute)

// 		res, height, err := cliCtx.QueryWithData(route, nil)
// 		if err != nil {
// 			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
// 			return
// 		}

// 		cliCtx = cliCtx.WithHeight(height)
// 		rest.PostProcessResponse(w, cliCtx, res)
// 	}
// }

func balancesHandler(cliCtx context.CLIContext, storeName string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		paramType := vars["address"]

		res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/balances/%s", storeName, paramType), nil)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusNotFound, err.Error())
			return
		}

		rest.PostProcessResponse(w, cliCtx, res)
	}
}

func totalSupplyHandler(cliCtx context.CLIContext, storeName string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/totalsupply", storeName), nil)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusNotFound, err.Error())
			return
		}

		rest.PostProcessResponse(w, cliCtx, res)
	}
}
