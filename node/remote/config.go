package remote

import (
	"fmt"
)

// Details represents a node details for a remote node
type Details struct {
	RPC                          *RPCConfig  `yaml:"rpc"`
	GRPC                         *GRPCConfig `yaml:"grpc"`
	API                          *APIConfig  `yaml:"api"`
	IgnoreConnectVoteExtensionTx bool        `yaml:"ignore_connect_vote_extension_tx"` // ignore tx[0] for the chains that are using Skip Oracle
}

func NewDetails(rpc *RPCConfig, grpc *GRPCConfig, api *APIConfig, ignoreConnectVoteExtensionTx bool) *Details {
	return &Details{
		RPC:                          rpc,
		GRPC:                         grpc,
		API:                          api,
		IgnoreConnectVoteExtensionTx: ignoreConnectVoteExtensionTx,
	}
}

func DefaultDetails() *Details {
	return NewDetails(DefaultRPCConfig(), DefaultGrpcConfig(), DefaultAPIConfig(), false)
}

// Validate implements node.Details
func (d *Details) Validate() error {
	if d.RPC == nil {
		return fmt.Errorf("rpc config cannot be null")
	}

	if d.GRPC == nil {
		return fmt.Errorf("grpc config cannot be null")
	}

	return nil
}

// --------------------------------------------------------------------------------------------------------------------

// RPCConfig contains the configuration for the RPC endpoint
type RPCConfig struct {
	ClientName     string `yaml:"client_name"`
	Address        string `yaml:"address"`
	MaxConnections int    `yaml:"max_connections"`
}

// NewRPCConfig allows to build a new RPCConfig instance
func NewRPCConfig(clientName, address string, maxConnections int) *RPCConfig {
	return &RPCConfig{
		ClientName:     clientName,
		Address:        address,
		MaxConnections: maxConnections,
	}
}

// DefaultRPCConfig returns the default instance of RPCConfig
func DefaultRPCConfig() *RPCConfig {
	return NewRPCConfig("juno", "http://localhost:26657", 20)
}

// --------------------------------------------------------------------------------------------------------------------

// GRPCConfig contains the configuration for the RPC endpoint
type GRPCConfig struct {
	Address  string `yaml:"address"`
	Insecure bool   `yaml:"insecure"`
}

// NewGrpcConfig allows to build a new GrpcConfig instance
func NewGrpcConfig(address string, insecure bool) *GRPCConfig {
	return &GRPCConfig{
		Address:  address,
		Insecure: insecure,
	}
}

// DefaultGrpcConfig returns the default instance of a GrpcConfig
func DefaultGrpcConfig() *GRPCConfig {
	return NewGrpcConfig("localhost:9090", true)
}

// --------------------------------------------------------------------------------------------------------------------

// APIConfig contains the configuration for the API endpoint
type APIConfig struct {
	Address string `yaml:"address"`
}

// NewAPIConfig allows to build a new APIConfig instance
func NewAPIConfig(address string) *APIConfig {
	return &APIConfig{
		Address: address,
	}
}

// DefaultAPIConfig returns the default instance of APIConfig
func DefaultAPIConfig() *APIConfig {
	return NewAPIConfig("http://localhost:1317")
}
