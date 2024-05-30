package types

import (
	"fmt"
	"strconv"
	"strings"
)

type HeightRange struct {
	Start int64
	End   int64
}

func (hr *HeightRange) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var heightRangeStr string
	if err := unmarshal(&heightRangeStr); err != nil {
		return err
	}

	heights := strings.Split(heightRangeStr, "-")
	if len(heights) != 2 {
		return fmt.Errorf("invalid height range format")
	}

	start, err := strconv.ParseInt(strings.TrimSpace(heights[0]), 10, 64)
	if err != nil {
		return fmt.Errorf("invalid start height: %w", err)
	}

	end, err := strconv.ParseInt(strings.TrimSpace(heights[1]), 10, 64)
	if err != nil {
		return fmt.Errorf("invalid end height: %w", err)
	}

	hr.Start = start
	hr.End = end

	return nil
}

// Includes returns true if the given height is within the range
func (hr *HeightRange) Includes(height int64) bool {
	return height >= hr.Start && height <= hr.End
}
