package codec

import (
	"github.com/cosmos/gogoproto/proto"
	"github.com/forbole/juno/v5/cosmos-sdk/codec/types"
	"google.golang.org/grpc/encoding"
)

type (
	// Codec defines a functionality for serializing other objects.
	// Users can defin a custom Protobuf-based serialization.
	// Note, Amino can still be used without any dependency on Protobuf.
	// SDK provides to Codec implementations:
	//
	// 1. AminoCodec: Provides full Amino serialization compatibility.
	// 2. ProtoCodec: Provides full Protobuf serialization compatibility.
	Codec interface {
		types.AnyUnpacker
		JSONCodec
	}

	GRPCodec interface {
		Codec
		GRPCCodecProvider
	}

	JSONCodec interface {
		// MarshalJSON returns JSON encoding of v.
		MarshalJSON(o proto.Message) ([]byte, error)
	}

	// GRPCCodecProvider is implemented by the Codec
	// implementations which return a gRPC encoding.Codec.
	// And it is used to decode requests and encode responses
	// passed through gRPC.
	GRPCCodecProvider interface {
		GRPCCodec() encoding.Codec
	}

	LegacyAmino interface {
		Marshal(v interface{}) ([]byte, error)
		Unmarshal(data []byte, v interface{}) error
		MarshalJSON(o interface{}) ([]byte, error)
	}
)
