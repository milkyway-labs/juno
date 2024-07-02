package bech32

import (
	"fmt"

	"github.com/cosmos/btcutil/bech32"
)

// ConvertAndEncode converts a byte array to base32 encoded byte string and then to bech32.
func ConvertAndEncode(hrp string, data []byte) (string, error) {
	// Converts the byte array to base32.
	// Base32 represents number using 5 bits, so we convert 8 bits encoded numbers to 5 bits encoded numbers.
	converted, err := bech32.ConvertBits(data, 8, 5, true)
	if err != nil {
		return "", fmt.Errorf("encoding bech32 failed: %w", err)
	}

	return bech32.Encode(hrp, converted)
}

// DecodeAndConvert decodes a bech32 encoded string and converts to bytes.
func DecodeAndConvert(bech string) (string, []byte, error) {
	// Decode the bech32 encode data, we limit the total resulting bytes to 1023.
	hrp, data, err := bech32.Decode(bech, 1023)
	if err != nil {
		return "", nil, fmt.Errorf("decoding bech32 failed: %w", err)
	}

	// Converts the bas32 encoded data back to byte array.
	// Base32 represents number using 5 bits, so we convert the 5 bits encoded number to 8 bits encoded numbers.
	converted, err := bech32.ConvertBits(data, 5, 8, false)
	if err != nil {
		return "", nil, fmt.Errorf("decoding bech32 failed: %w", err)
	}

	return hrp, converted, nil
}
