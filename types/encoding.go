package types

import (
	"github.com/forbole/juno/v5/cosmos-sdk/codec"
)

// EncodingConfig specifies the concrete encoding types to use for a given app.
// This is provided for compatibility between protobuf and amino implementations.
// NOTE: This is copied from simapp in order to avoid importing that package as a dependency
type EncodingConfig struct {
	Codec codec.ProtoCodec
	Amino *codec.LegacyAmino
}

// MakeTestEncodingConfig creates an EncodingConfig for a non-amino based test configuration.
// This function should be used only internally (in the SDK).
// App user shouldn't create new codecs - use the app.AppCodec instead.
// [DEPRECATED]
// NOTE: This is copied from simapp in order to avoid importing that package as a dependency
func MakeTestEncodingConfig() EncodingConfig {
	// amino := codec.NewLegacyAmino()
	// interfaceRegistry := types.NewInterfaceRegistry()
	// cdc := codec.NewProtoCodec(interfaceRegistry)
	//
	// return EncodingConfig{
	// 	InterfaceRegistry: interfaceRegistry,
	// 	Codec:             cdc,
	// 	TxConfig:          tx.NewTxConfig(cdc, tx.DefaultSignModes),
	// 	Amino:             amino,
	// }
	// TODO: Fix me.
	panic("Fix MakeTestEncodingConfig")
}
