package utils

import (
	"encoding/json"
	"fmt"
	"strings"

	tmtypes "github.com/cometbft/cometbft/types"

	"github.com/forbole/juno/v5/node"
	"github.com/forbole/juno/v5/utils"
)

// GetGenesisDocAndState reads the genesis from node or file and returns genesis doc and state
func GetGenesisDocAndState(genesisPath string, node node.Node) (*tmtypes.GenesisDoc, map[string]json.RawMessage, error) {
	var genesisDoc *tmtypes.GenesisDoc
	if strings.TrimSpace(genesisPath) != "" {
		genDoc, err := utils.ReadGenesisFileGenesisDoc(genesisPath)
		if err != nil {
			return nil, nil, fmt.Errorf("error while reading genesis file: %s", err)
		}
		genesisDoc = genDoc

	} else {
		response, err := node.Genesis()
		if err != nil {
			return nil, nil, fmt.Errorf("failed to get genesis: %s", err)
		}
		genesisDoc = response.Genesis
	}

	genesisState, err := utils.GetGenesisState(genesisDoc)
	if err != nil {
		return nil, nil, err
	}

	return genesisDoc, genesisState, nil
}
