package types

// A representation of the Cosmos-SDK config struct.
// This will be used when we need to access the Config from the SDK.
type Config interface {
	// GetBech32AccountAddrPrefix returns the Bech32 prefix for account address
	GetBech32AccountAddrPrefix() string

	// GetBech32ValidatorAddrPrefix returns the Bech32 prefix for validator address
	GetBech32ValidatorAddrPrefix() string

	// GetBech32ConsensusAddrPrefix returns the Bech32 prefix for consensus node address
	GetBech32ConsensusAddrPrefix() string

	// GetBech32AccountPubPrefix returns the Bech32 prefix for account public key
	GetBech32AccountPubPrefix() string

	// GetBech32ValidatorPubPrefix returns the Bech32 prefix for validator public key
	GetBech32ValidatorPubPrefix() string

	// GetBech32ConsensusPubPrefix returns the Bech32 prefix for consensus node public key
	GetBech32ConsensusPubPrefix() string
}

type ConfigGetter = func() Config

var configGetter ConfigGetter

func SetSdkConfigGetter(getter ConfigGetter) {
	configGetter = getter
}

func GetSdkConfig() Config {
	if configGetter == nil {
		panic("set the cosmos sdk confg using the SetSdkConfig function")
	}
	return configGetter()
}
