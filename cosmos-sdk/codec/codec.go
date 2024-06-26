package codec

import (
	"github.com/cosmos/gogoproto/proto"
	gogoproto "github.com/cosmos/gogoproto/types/any"
	"google.golang.org/grpc/encoding"
)

type Codec interface {
	encoding.Codec
	MarshalJSON(o interface{}) ([]byte, error)
	UnpackAny(any *gogoproto.Any, iface interface{}) error
}

type LegacyAmino interface {
	MarshalJSON(o interface{}) ([]byte, error)
}

type ProtoCodec interface {
	Codec
	GRPCCodec() Codec
}

type JSONCodec interface {
	// MarshalJSON returns JSON encoding of v.
	MarshalJSON(o proto.Message) ([]byte, error)
	// MustMarshalJSON calls MarshalJSON and panics if error is returned.
	MustMarshalJSON(o proto.Message) []byte
	// MarshalInterfaceJSON is a helper method which will wrap `i` into `Any` for correct
	// JSON interface (de)serialization.
	MarshalInterfaceJSON(i proto.Message) ([]byte, error)
	// UnmarshalInterfaceJSON is a helper method which will parse JSON enoded data
	// into `Any` and unpack any into the `ptr`. It fails if the target interface type
	// is not registered in codec, or is not compatible with the serialized data
	UnmarshalInterfaceJSON(bz []byte, ptr interface{}) error

	// UnmarshalJSON parses the data encoded with MarshalJSON method and stores the result
	// in the value pointed to by v.
	UnmarshalJSON(bz []byte, ptr proto.Message) error
	// MustUnmarshalJSON calls Unmarshal and panics if error is returned.
	MustUnmarshalJSON(bz []byte, ptr proto.Message)
}
