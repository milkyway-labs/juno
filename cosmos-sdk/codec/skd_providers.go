package codec

type ProtoCodecProvider = func() Codec

var protoCodecProvider ProtoCodecProvider

func SetProtoCodecProvider(p ProtoCodecProvider) {
	protoCodecProvider = p
}

func NewProtoCodec() Codec {
	if protoCodecProvider == nil {
		panic("proto codec provider not set, call SetProtoCodecProvider to set it")
	}

	return protoCodecProvider()
}
