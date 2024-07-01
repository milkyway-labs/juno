package types

import "fmt"

// String returns a string representation of Height
func (h Height) String() string {
	return fmt.Sprintf("%d-%d", h.RevisionNumber, h.RevisionHeight)
}
