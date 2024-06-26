package types

import (
	"github.com/cosmos/gogoproto/proto"
	"github.com/forbole/juno/v5/cosmos-sdk/core/transaction"
)

type (
	Msg = transaction.Msg
)

// MsgTypeURL returns the TypeURL of a `sdk.Msg`.
func MsgTypeURL(msg Msg) string {
	return "/" + proto.MessageName(msg)
}
