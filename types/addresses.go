package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// AccountAddressParser represents a function that takes as input an address string and returns
// whether the address represents what should be considered a valid account address.
// This is used to determine the accounts that are involved in transactions and messages.
type AccountAddressParser func(address string) (isValid bool)

// DefaultAddressParser returns a default implementation of the AccountAddressParser.
func DefaultAddressParser() AccountAddressParser {
	return func(address string) bool {
		_, err := sdk.AccAddressFromBech32(address)
		if err == nil {
			return true
		}

		_, err = sdk.ValAddressFromBech32(address)
		return err == nil
	}
}
