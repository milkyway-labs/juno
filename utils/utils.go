package utils

import (
	"fmt"
	"unicode/utf8"

	abci "github.com/cometbft/cometbft/abci/types"
	cometbfttypes "github.com/cometbft/cometbft/types"
)

// ConvertValidatorAddressToBech32String converts the given validator address to its Bech32 string representation
func ConvertValidatorAddressToBech32String(address cometbfttypes.Address) string {
	return ""
}

func FindEventByType(events []abci.Event, eventType string) (abci.Event, error) {
	for _, event := range events {
		if event.Type == eventType {
			return event, nil
		}
	}

	return abci.Event{}, fmt.Errorf("no event with type %s found", eventType)
}

func FindEventsByType(events []abci.Event, eventType string) []abci.Event {
	var found []abci.Event
	for _, event := range events {
		if event.Type == eventType {
			found = append(found, event)
		}
	}

	return found
}

func FindAttributeByKey(event abci.Event, attrKey string) (abci.EventAttribute, error) {
	for _, attr := range event.Attributes {
		if attr.Key == attrKey {
			return attr, nil
		}
	}

	return abci.EventAttribute{}, fmt.Errorf("no attribute with key %s found inside event with type %s", attrKey, event.Type)
}

func MaxInt64(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func TrimLastChar(s string) string {
	r, size := utf8.DecodeLastRuneInString(s)
	if r == utf8.RuneError && (size == 0 || size == 1) {
		size = 0
	}
	return s[:len(s)-size]
}
