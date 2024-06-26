package types

import (
	"encoding/hex"
	"encoding/json"
	strings "strings"

	coretypes "github.com/cometbft/cometbft/rpc/core/types"
	codectypes "github.com/forbole/juno/v5/cosmos-sdk/codec/types"
)

// NewResponseResultTx returns a TxResponse given a ResultTx from tendermint
func NewResponseResultTx(res *coretypes.ResultTx, anyTx *codectypes.Any, timestamp string) *TxResponse {
	if res == nil {
		return nil
	}

	parsedLogs, _ := ParseABCILogs(res.TxResult.Log)

	return &TxResponse{
		TxHash:    res.Hash.String(),
		Height:    res.Height,
		Codespace: res.TxResult.Codespace,
		Code:      res.TxResult.Code,
		Data:      strings.ToUpper(hex.EncodeToString(res.TxResult.Data)),
		RawLog:    res.TxResult.Log,
		Logs:      parsedLogs,
		Info:      res.TxResult.Info,
		GasWanted: res.TxResult.GasWanted,
		GasUsed:   res.TxResult.GasUsed,
		Tx:        anyTx,
		Timestamp: timestamp,
		Events:    res.TxResult.Events,
	}
}

// ParseABCILogs attempts to parse a stringified ABCI tx log into a slice of
// ABCIMessageLog types. It returns an error upon JSON decoding failure.
func ParseABCILogs(logs string) (res ABCIMessageLogs, err error) {
	err = json.Unmarshal([]byte(logs), &res)
	return res, err
}
