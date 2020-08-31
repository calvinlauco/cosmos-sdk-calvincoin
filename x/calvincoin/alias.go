package calvincoin

import (
	"github.com/calvinlauco/cosmos-sdk-calvincoin/x/calvincoin/keeper"
	"github.com/calvinlauco/cosmos-sdk-calvincoin/x/calvincoin/types"
)

const (
	ModuleName   = types.ModuleName
	RouterKey    = types.RouterKey
	StoreKey     = types.StoreKey
	QuerierRoute = types.QuerierRoute
)

var (
	NewKeeper      = keeper.NewKeeper
	NewQuerier     = keeper.NewQuerier
	NewMsgTransfer = types.NewMsgTransfer
	ModuleCdc      = types.ModuleCdc
	RegisterCodec  = types.RegisterCodec
)

type (
	Keeper      = keeper.Keeper
	MsgTransfer = types.MsgTransfer
)
