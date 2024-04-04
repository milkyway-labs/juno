package types

// TransactionFilter represents a function that takes as input a transaction and returns
// whether it should be stored or not.
type TransactionFilter func(tx *Tx) (shouldStore bool)

// DefaultTransactionFilter returns a default implementation of the TransactionFilter
func DefaultTransactionFilter() TransactionFilter {
	return func(tx *Tx) (shouldStore bool) {
		return true
	}
}

// MessageFilter represents a function that takes as input a message and returns
// whether the message should be stored or not.
type MessageFilter func(msg *Message) (shouldStore bool)

// DefaultMessageFilter returns a default implementation of the MessageFilter
func DefaultMessageFilter() MessageFilter {
	return func(msg *Message) bool {
		return true
	}
}
