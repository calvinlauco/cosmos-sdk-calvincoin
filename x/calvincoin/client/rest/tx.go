package rest

// The packages below are commented out at first to prevent an error if this file isn't initially saved.
import (
	// "bytes"

	"net/http"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/calvinlauco/cosmos-sdk-calvincoin/x/calvincoin/types"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
)

type transferReq struct {
	BaseReq rest.BaseReq `json:"base_req"`
	From    string       `json:"from"`
	To      string       `json:"to"`
	Amount  string       `json:"amount"`
}

// func registerTxRoutes(cliCtx context.CLIContext, r *mux.Router) {
// 	// r.HandleFunc(
// 	// TODO: Define the Rest route ,
// 	// Call the function which should be executed for this route),
// 	// ).Methods("POST")
// }

/*
// Action TX body
type <Action>Req struct {
	BaseReq rest.BaseReq `json:"base_req" yaml:"base_req"`
	// TODO: Define more types if needed
}

func <Action>RequestHandlerFn(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req <Action>Req
		vars := mux.Vars(r)

		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}

		// TODO: Define the module tx logic for this action

		utils.WriteGenerateStdTxResponse(w, cliCtx, BaseReq, []sdk.Msg{msg})
	}
}
*/

func transferHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req transferReq

		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
			return
		}

		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}

		fromAddr, err := sdk.AccAddressFromBech32(req.From)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		toAddr, err := sdk.AccAddressFromBech32(req.From)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		amount, err := sdk.ParseCoins(req.Amount)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		// create the message
		msg := types.NewMsgTransfer(fromAddr, toAddr, amount)
		err = msg.ValidateBasic()
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		utils.WriteGenerateStdTxResponse(w, cliCtx, baseReq, []sdk.Msg{msg})
	}
}
