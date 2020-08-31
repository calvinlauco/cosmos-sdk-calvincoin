package calvincoin

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/calvinlauco/cosmos-sdk-calvincoin/x/calvincoin/types"
)

// NewHandler creates an sdk.Handler for all the calvincoin type messages
func NewHandler(k Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())
		switch msg := msg.(type) {
		// TODO: Define your msg cases
		//
		//Example:
		// case Msg<Action>:
		// 	return handleMsg<Action>(ctx, k, msg)
		case types.MsgTransfer:
			return handleMsgTransfer(ctx, k, msg)
		default:
			errMsg := fmt.Sprintf("unrecognized %s message type: %T", ModuleName, msg)
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
		}
	}
}

// handle<Action> does x
// func handleMsg<Action>(ctx sdk.Context, k Keeper, msg Msg<Action>) (*sdk.Result, error) {
// 	err := k.<Action>(ctx, msg.ValidatorAddr)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// TODO: Define your msg events
// 	ctx.EventManager().EmitEvent(
// 		sdk.NewEvent(
// 			sdk.EventTypeMessage,
// 			sdk.NewAttribute(sdk.AttributeKeyModule, AttributeValueCategory),
// 			sdk.NewAttribute(sdk.AttributeKeySender, msg.ValidatorAddr.String()),
// 		),
// 	)

// 	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
// }

func handleMsgTransfer(ctx sdk.Context, keeper Keeper, msg types.MsgTransfer) (*sdk.Result, error) {
	var err error
	_, err = keeper.CoinKeeper.SubtractCoins(ctx, msg.From, msg.Amount)
	if err != nil {
		return nil, err
	}

	err = keeper.CoinKeeper.SendCoins(ctx, msg.From, msg.To, msg.Amount)
	if err != nil {
		return nil, err
	}

	return &sdk.Result{}, nil
}
