package config_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"

	nodeconfig "github.com/forbole/juno/v5/node/config"
	"github.com/forbole/juno/v5/node/remote"
)

func TestConfig_UnmarshalYAML(t *testing.T) {
	var remoteData = `
type: "remote"
config:
  rpc:
    client_name: "juno"
    max_connections: 1
    address: "http://localhost:26657"

  grpc:
    insecure: true
    address: "http://localhost:9090"
  
  ignore_connect_vote_extension_tx: false
`

	var config nodeconfig.Config
	err := yaml.Unmarshal([]byte(remoteData), &config)
	require.NoError(t, err)
	require.IsType(t, &remote.Details{}, config.Details)
}

func TestConfig_MarshalYAML(t *testing.T) {
	config := nodeconfig.Config{
		Type: nodeconfig.TypeRemote,
		Details: &remote.Details{
			RPC: &remote.RPCConfig{
				ClientName:     "juno",
				Address:        "http://localhost:26657",
				MaxConnections: 10,
			},
			GRPC: &remote.GRPCConfig{
				Address: "http://localhost:9090",
			},
		},
	}
	bz, err := yaml.Marshal(&config)
	require.NoError(t, err)

	expected := `
type: remote
config:
    rpc:
        client_name: juno
        address: http://localhost:26657
        max_connections: 10
    grpc:
        address: http://localhost:9090
    ignore_connect_vote_extension_tx: false
`
	require.Equal(t, strings.TrimLeft(expected, "\n"), string(bz))
}
