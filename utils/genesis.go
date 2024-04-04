package utils

import (
	"encoding/json"
	"fmt"

	tmjson "github.com/cometbft/cometbft/libs/json"
	tmos "github.com/cometbft/cometbft/libs/os"
	tmtypes "github.com/cometbft/cometbft/types"
)

// ReadGenesisFileGenesisDoc reads the genesis file located at the given path
func ReadGenesisFileGenesisDoc(genesisPath string) (*tmtypes.GenesisDoc, error) {
	var genesisDoc *tmtypes.GenesisDoc
	bz, err := tmos.ReadFile(genesisPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read genesis file: %s", err)
	}

	err = tmjson.Unmarshal(bz, &genesisDoc)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal genesis doc: %s", err)
	}

	return genesisDoc, nil
}

// GetGenesisState returns the genesis state by getting it from the given genesis doc
func GetGenesisState(doc *tmtypes.GenesisDoc) (map[string]json.RawMessage, error) {
	var genesisState map[string]json.RawMessage
	err := json.Unmarshal(doc.AppState, &genesisState)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal genesis state: %s", err)
	}
	return genesisState, nil
}
