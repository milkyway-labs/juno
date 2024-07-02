package types

import (
	"encoding/json"

	"github.com/forbole/juno/v5/cosmos-sdk/codec"
)

func (gi GasInfo) String() string {
	bz, _ := codec.MarshalYAML(codec.NewProtoCodec(), &gi)
	return string(bz)
}

func (r Result) String() string {
	bz, _ := codec.MarshalYAML(codec.NewProtoCodec(), &r)
	return string(bz)
}

func (r TxResponse) String() string {
	bz, _ := codec.MarshalYAML(codec.NewProtoCodec(), &r)
	return string(bz)
}

// ABCIMessageLogs represents a slice of ABCIMessageLog.
type ABCIMessageLogs []ABCIMessageLog

// ParseABCILogs attempts to parse a stringified ABCI tx log into a slice of
// ABCIMessageLog types. It returns an error upon JSON decoding failure.
func ParseABCILogs(logs string) (res ABCIMessageLogs, err error) {
	err = json.Unmarshal([]byte(logs), &res)
	return res, err
}
