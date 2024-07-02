package codec

import (
	"github.com/cosmos/gogoproto/proto"
	"google.golang.org/grpc/encoding"
)

type (
	// Codec defines a functionality for serializing other objects.
	Codec interface {
		JSONCodec
	}

	JSONCodec interface {
		// MarshalJSON returns JSON encoding of v.
		MarshalJSON(o proto.Message) ([]byte, error)
	}

	// ProtoMarshaler defines an interface a type must implement to serialize itself
	// as a protocol buffer defined message.
	ProtoMarshaler interface {
		proto.Message // for JSON serialization

		Marshal() ([]byte, error)
		MarshalTo(data []byte) (n int, err error)
		MarshalToSizedBuffer(dAtA []byte) (int, error)
		Size() int
		Unmarshal(data []byte) error
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
