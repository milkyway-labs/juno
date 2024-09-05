package types

import (
	"github.com/cometbft/cometbft/crypto/tmhash"
)

// TxHashCalculator represents a function that given a transaction returns its hash
type TxHashCalculator func(tx []byte) []byte

// DefaultTxHashCalculator returns the default transaction hash calculator.
// This function uses the tmhash.Sum function to calculate the hash of the given transaction
func DefaultTxHashCalculator(tx []byte) []byte {
	return tmhash.Sum(tx)
}
