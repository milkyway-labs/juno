package types

import (
	fmt "fmt"
)

// ABCIMessageLogs represents a slice of ABCIMessageLog.
type ABCIMessageLogs []ABCIMessageLog

// StringAttributes defines a slice of StringEvents objects.
type StringEvents []StringEvent

func (r TxResponse) String() string {
	//bz, _ := codec.MarshalYAML(codec.NewProtoCodec(nil), &r)
	//return string(bz)
	return ""
}

func (a Attribute) String() string {
	return fmt.Sprintf("%s: %s", a.Key, a.Value)
}

func (gi GasInfo) String() string {
	//bz, _ := codec.MarshalYAML(codec.NewProtoCodec(nil), &gi)
	//return string(bz)
	return ""
}

func (r Result) String() string {
	// bz, _ := codec.MarshalYAML(codec.NewProtoCodec(nil), &r)
	// return string(bz)
	return ""
}
