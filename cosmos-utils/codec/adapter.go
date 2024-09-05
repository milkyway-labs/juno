package codec

import (
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/gogoproto/proto"
	"google.golang.org/grpc/encoding"

	junocodec "github.com/forbole/juno/v5/cosmos-sdk/codec"
	junocodectypes "github.com/forbole/juno/v5/cosmos-sdk/codec/types"
)

type JunoCodecAdapater struct {
	codec *codec.ProtoCodec
}

func NewJunoCodecAdapter(codec *codec.ProtoCodec) *JunoCodecAdapater {
	return &JunoCodecAdapater{
		codec: codec,
	}
}

func (j *JunoCodecAdapater) MarshalJSON(o proto.Message) ([]byte, error) {
	return j.codec.MarshalJSON(o)
}

func (j *JunoCodecAdapater) GetSdkCodec() *codec.ProtoCodec {
	return j.codec
}

func (j *JunoCodecAdapater) GRPCCodec() encoding.Codec {
	return j.codec.GRPCCodec()
}

func (j *JunoCodecAdapater) UnpackAny(any *junocodectypes.Any, iface interface{}) error {
	sdkAny := codectypes.Any{
		TypeUrl: any.TypeUrl,
		Value:   any.Value,
	}
	return j.codec.UnpackAny(&sdkAny, iface)
}

func CastCodec(c junocodec.Codec) *JunoCodecAdapater {
	protoCodec, ok := c.(*JunoCodecAdapater)
	if !ok {
		panic("invalid codec")
	}
	return protoCodec
}
