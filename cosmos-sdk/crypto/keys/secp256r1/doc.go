package secp256r1

import "crypto/elliptic"

const (
	// fieldSize is the curve domain size.
	fieldSize  = 32
	pubKeySize = fieldSize + 1

	name = "secp256r1"
)

var secp256r1 elliptic.Curve
