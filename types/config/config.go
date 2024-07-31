package config

import (
	"strings"

	"gopkg.in/yaml.v3"

	databaseconfig "github.com/forbole/juno/v5/database/config"
	loggingconfig "github.com/forbole/juno/v5/logging/config"
	nodeconfig "github.com/forbole/juno/v5/node/config"
	parserconfig "github.com/forbole/juno/v5/parser/config"
)

// Config defines all necessary juno configuration parameters.
type Config struct {
	bytes []byte

	Chain      ChainConfig           `yaml:"chain"`
	Node       nodeconfig.Config     `yaml:"node"`
	Parser     parserconfig.Config   `yaml:"parsing"`
	Database   databaseconfig.Config `yaml:"database"`
	Logging    loggingconfig.Config  `yaml:"logging"`
	Monitoring MonitoringConfig      `yaml:"monitoring"`
}

// NewConfig builds a new Config instance
func NewConfig(
	nodeCfg nodeconfig.Config,
	chainCfg ChainConfig, dbConfig databaseconfig.Config,
	parserConfig parserconfig.Config, loggingConfig loggingconfig.Config,
	monitoringConfig MonitoringConfig,
) Config {
	return Config{
		Node:       nodeCfg,
		Chain:      chainCfg,
		Database:   dbConfig,
		Parser:     parserConfig,
		Logging:    loggingConfig,
		Monitoring: monitoringConfig,
	}
}

func DefaultConfig() Config {
	cfg := NewConfig(
		nodeconfig.DefaultConfig(),
		DefaultChainConfig(), databaseconfig.DefaultDatabaseConfig(),
		parserconfig.DefaultParsingConfig(), loggingconfig.DefaultLoggingConfig(),
		DefaultMonitoringConfig(),
	)

	bz, err := yaml.Marshal(cfg)
	if err != nil {
		panic(err)
	}

	cfg.bytes = bz
	return cfg
}

func (c Config) GetBytes() ([]byte, error) {
	return c.bytes, nil
}

// ---------------------------------------------------------------------------------------------------------------------

type ChainConfig struct {
	Bech32Prefix string   `yaml:"bech32_prefix"`
	Modules      []string `yaml:"modules"`
}

// NewChainConfig returns a new ChainConfig instance
func NewChainConfig(bech32Prefix string, modules []string) ChainConfig {
	return ChainConfig{
		Bech32Prefix: bech32Prefix,
		Modules:      modules,
	}
}

// DefaultChainConfig returns the default instance of ChainConfig
func DefaultChainConfig() ChainConfig {
	return NewChainConfig("cosmos", nil)
}

func (cfg ChainConfig) IsModuleEnabled(moduleName string) bool {
	for _, module := range cfg.Modules {
		if strings.EqualFold(module, moduleName) {
			return true
		}
	}

	return false
}

// ---------------------------------------------------------------------------------------------------------------------

type MonitoringConfig struct {
	Enabled bool  `yaml:"enabled"`
	Port    int16 `yaml:"port"`
}

// DefaultMonitoringConfig returns the default instance of MonitoringConfig
func DefaultMonitoringConfig() MonitoringConfig {
	return MonitoringConfig{
		Enabled: true,
		Port:    2112,
	}
}

// NewMonitoringConfig returns a new instance of MonitoringConfig
func NewMonitoringConfig(enabled bool, port int16) MonitoringConfig {
	return MonitoringConfig{
		Enabled: enabled,
		Port:    port,
	}
}
