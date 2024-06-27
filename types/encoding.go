package types

import (
	"github.com/forbole/juno/v5/cosmos-sdk/codec"
)

// EncodingConfig specifies the concrete encoding types to use for a given app.
// This is provided for compatibility between protobuf and amino implementations.
// NOTE: This is copied from simapp in order to avoid importing that package as a dependency
type EncodingConfig struct {
	Codec codec.GRPCodec
	Amino codec.LegacyAmino
}
