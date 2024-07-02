package types

import (
	"github.com/cosmos/gogoproto/proto"
	protov2 "google.golang.org/protobuf/proto"
)

// MsgTypeURL returns the TypeURL of a `sdk.Msg`.
func MsgTypeURL(msg Msg) string {
	return "/" + proto.MessageName(msg)
}

// TxDecoder unmarshals transaction bytes
type TxDecoder func(txBytes []byte) (Tx, error)

// TxEncoder marshals transaction to bytes
type TxEncoder func(tx Tx) ([]byte, error)

type (
	// Msg defines the interface a transaction message needed to fulfill.
	Msg = proto.Message

	// HasMsgs defines an interface a transaction must fulfill.
	HasMsgs interface {
		// GetMsgs gets the all the transaction's messages.
		GetMsgs() []Msg
	}

	// Tx defines an interface a transaction must fulfill.
	Tx interface {
		HasMsgs

		// GetMsgsV2 gets the transaction's messages as google.golang.org/protobuf/proto.Message's.
		GetMsgsV2() ([]protov2.Message, error)
	}
)
