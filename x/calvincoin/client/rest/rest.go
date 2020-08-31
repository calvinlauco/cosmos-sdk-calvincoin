package rest

import (
	"fmt"

	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client/context"
)

// RegisterRoutes registers calvincoin-related REST handlers to a router
func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router, storeName string) {
	// registerQueryRoutes(cliCtx, r)
	// registerTxRoutes(cliCtx, r)
	r.HandleFunc(fmt.Sprintf("/balances"), balancesHandler(cliCtx, storeName)).Methods("GET")
	r.HandleFunc(fmt.Sprintf("/totalsupply"), totalSupplyHandler(cliCtx, storeName)).Methods("GET")
	r.HandleFunc(fmt.Sprintf("/transfer"), transferHandler(cliCtx)).Methods("POST")
}
