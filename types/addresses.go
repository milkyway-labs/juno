package types

// AccountAddressParser represents a function that takes as input an address string and returns
// whether the address represents what should be considered a valid account address.
// This is used to determine the accounts that are involved in transactions and messages.
type AccountAddressParser func(address string) (isValid bool)
