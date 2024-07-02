package types

import (
	junocodec "github.com/forbole/juno/v5/cosmos-sdk/codec"
	"google.golang.org/grpc/encoding"
)

// EncodingConfig specifies the concrete encoding types to use for a given app.
// This is provided for compatibility between protobuf and amino implementations.
// NOTE: This is copied from simapp in order to avoid importing that package as a dependency
type EncodingConfig struct {
	Codec    junocodec.Codec
	Amino    junocodec.LegacyAmino
	GRPCodec encoding.Codec
}
