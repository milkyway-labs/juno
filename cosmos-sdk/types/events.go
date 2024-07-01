package types

import fmt "fmt"

// StringAttributes defines a slice of StringEvents objects.
type StringEvents []StringEvent

func (a Attribute) String() string {
	return fmt.Sprintf("%s: %s", a.Key, a.Value)
}
