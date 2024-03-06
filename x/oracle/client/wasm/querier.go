package wasm

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	wasmbindingstypes "github.com/elys-network/elys/wasmbindings/types"
	"github.com/ojo-network/ojo/x/oracle/keeper"
)

// Querier handles queries for the Oracle module.
type Querier struct {
	keeper *keeper.Keeper
}

func NewQuerier(keeper *keeper.Keeper) *Querier {
	return &Querier{keeper: keeper}
}

func (oq *Querier) HandleQuery(ctx sdk.Context, query wasmbindingstypes.ElysQuery) ([]byte, error) {
	switch {
	case query.OracleParams != nil:
		return oq.queryParams(ctx, query.OracleParams)
	case query.OracleAssetInfo != nil:
		return oq.queryAssetInfo(ctx, query.OracleAssetInfo)
	case query.OracleAssetInfoAll != nil:
		return oq.queryAssetInfoAll(ctx, query.OracleAssetInfoAll)
	case query.OraclePrice != nil:
		return oq.queryPrice(ctx, query.OraclePrice)
	case query.OraclePriceAll != nil:
		return oq.queryPriceAll(ctx, query.OraclePriceAll)
	default:
		// This handler cannot handle the query
		return nil, wasmbindingstypes.ErrCannotHandleQuery
	}
}
