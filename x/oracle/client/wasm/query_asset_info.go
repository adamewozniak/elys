package wasm

import (
	"encoding/json"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	oracletypes "github.com/elys-network/elys/x/oracle/types"
)

func (oq *Querier) queryAssetInfo(ctx sdk.Context, req *oracletypes.QueryGetAssetInfoRequest) ([]byte, error) {
	res := oq.keeper.GetParams(ctx)

	responseBytes, err := json.Marshal(res)
	if err != nil {
		return nil, errorsmod.Wrap(err, "failed to serialize asset info response")
	}
	return responseBytes, nil
}
