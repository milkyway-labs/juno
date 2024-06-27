package types

type ConsAddressBuilder = func([]byte) ConsAddress

type ConsAddress interface {
	String() string
}

var buildConsAddress ConsAddressBuilder

func SetConsAddressBuilder(b ConsAddressBuilder) {
	buildConsAddress = b
}

func NewConsAddress(bz []byte) ConsAddress {
	if buildConsAddress == nil {
		panic("cons address builder not set, please set it using the SetConsAddressBuilder function")
	}
	return buildConsAddress(bz)
}
