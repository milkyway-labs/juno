package utils

import (
	abci "github.com/cometbft/cometbft/abci/types"
	tmabci "github.com/tendermint/tendermint/abci/types"
)

// ConvertTendermintEventToCometBFTEvent converts a Tendermint event to a CometBFT event.
// This is needed because Cosmos SKD v0.46.16 (used by Dymension) replaces the Tendermint library with the
// CometBFT library, which have incompatible types.
func ConvertTendermintEventToCometBFTEvent(event tmabci.Event) abci.Event {
	attrs := make([]abci.EventAttribute, len(event.Attributes))
	for i, attr := range event.Attributes {
		attrs[i] = abci.EventAttribute{
			Key:   string(attr.Key),
			Value: string(attr.Value),
		}
	}

	return abci.Event{
		Type:       event.Type,
		Attributes: attrs,
	}
}

func ConvertCometBFTEventToTendermintEvent(event abci.Event) tmabci.Event {
	attrs := make([]tmabci.EventAttribute, len(event.Attributes))
	for i, attr := range event.Attributes {
		attrs[i] = tmabci.EventAttribute{
			Key:   []byte(attr.Key),
			Value: []byte(attr.Value),
		}
	}

	return tmabci.Event{
		Type:       event.Type,
		Attributes: attrs,
	}
}
