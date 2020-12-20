package nameservice

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/Rajakavitha1/nameservice/x/nameservice/types"
	"github.com/Rajakavitha1/nameservice/x/nameservice/keeper"
)

func handleMsgCreateWhois(ctx sdk.Context, k keeper.Keeper, msg types.MsgCreateWhois) (*sdk.Result, error) {
	var whois = types.Whois{
		Creator: msg.Creator,
		ID:      msg.ID,
    	Value: msg.Value,
    	Price: msg.Price,
	}
	k.CreateWhois(ctx, whois)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
