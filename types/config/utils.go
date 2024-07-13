package config

import (
	"time"
)

// GetAvgBlockTime returns the average_block_time in the configuration file or
// returns 3 seconds if it is not configured
func (cfg *Config) GetAvgBlockTime() time.Duration {
	if cfg.Parser.AvgBlockTime == nil {
		return 3 * time.Second
	}
	return *cfg.Parser.AvgBlockTime
}
