package types

import "github.com/forbole/juno/v5/cosmos-sdk/types/bech32"

type ConsAddress []byte

func NewConsAddress(bz []byte) ConsAddress {
	return ConsAddress(bz)
}

func (bz ConsAddress) String() string {
	addr, err := bech32.ConvertAndEncode(GetSdkConfig().GetBech32ConsensusAddrPrefix(), bz)
	if err != nil {
		// Panic like the cosmos-sdk
		panic(err)
	}
	return addr
}
