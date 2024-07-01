package tx

import transaction "github.com/forbole/juno/v5/cosmos-sdk/core/transaction"

// GetMsgs implements the GetMsgs method on sdk.Tx.
func (t *Tx) GetMsgs() []transaction.Msg {
	if t == nil || t.Body == nil {
		return nil
	}

	anys := t.Body.Messages
	res, err := GetMsgs(anys, "transaction")
	if err != nil {
		panic(err)
	}
	return res
}
