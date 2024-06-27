package types

import (
	"github.com/forbole/juno/v5/database"
	"github.com/forbole/juno/v5/database/builder"
	"github.com/forbole/juno/v5/logging"
	"github.com/forbole/juno/v5/modules/registrar"
	"github.com/forbole/juno/v5/types"
	"github.com/forbole/juno/v5/types/config"
)

// Config contains all the configuration for the "parse" command
type Config struct {
	registrar             registrar.Registrar
	configParser          config.Parser
	encodingConfigBuilder EncodingConfigBuilder
	buildDb               database.Builder
	logger                logging.Logger
	accountAddressParser  types.AccountAddressParser

	transactionFilter    types.TransactionFilter
	messageFilterBuilder MessageFilterBuilder
}

// NewConfig allows to build a new Config instance
func NewConfig() *Config {
	return &Config{}
}

// WithRegistrar sets the modules registrar to be used
func (cfg *Config) WithRegistrar(r registrar.Registrar) *Config {
	cfg.registrar = r
	return cfg
}

// GetRegistrar returns the modules registrar to be used
func (cfg *Config) GetRegistrar() registrar.Registrar {
	if cfg.registrar == nil {
		return &registrar.EmptyRegistrar{}
	}
	return cfg.registrar
}

// WithConfigParser sets the configuration parser to be used
func (cfg *Config) WithConfigParser(p config.Parser) *Config {
	cfg.configParser = p
	return cfg
}

// GetConfigParser returns the configuration parser to be used
func (cfg *Config) GetConfigParser() config.Parser {
	if cfg.configParser == nil {
		return config.DefaultConfigParser
	}
	return cfg.configParser
}

// WithEncodingConfigBuilder sets the configurations builder to be used
func (cfg *Config) WithEncodingConfigBuilder(b EncodingConfigBuilder) *Config {
	cfg.encodingConfigBuilder = b
	return cfg
}

// GetEncodingConfigBuilder returns the encoding config builder to be used
func (cfg *Config) GetEncodingConfigBuilder() EncodingConfigBuilder {
	if cfg.encodingConfigBuilder == nil {
		return func() types.EncodingConfig {
			encodingConfig := types.MakeTestEncodingConfig()
			// std.RegisterLegacyAminoCodec(encodingConfig.Amino)
			// std.RegisterInterfaces(encodingConfig.InterfaceRegistry)
			// ModuleBasics.RegisterLegacyAminoCodec(encodingConfig.Amino)
			// ModuleBasics.RegisterInterfaces(encodingConfig.InterfaceRegistry)
			return encodingConfig
		}
	}
	return cfg.encodingConfigBuilder
}

// WithDBBuilder sets the database builder to be used
func (cfg *Config) WithDBBuilder(b database.Builder) *Config {
	cfg.buildDb = b
	return cfg
}

// GetDBBuilder returns the database builder to be used
func (cfg *Config) GetDBBuilder() database.Builder {
	if cfg.buildDb == nil {
		return builder.Builder
	}
	return cfg.buildDb
}

// WithLogger sets the logger to be used while parsing the data
func (cfg *Config) WithLogger(logger logging.Logger) *Config {
	cfg.logger = logger
	return cfg
}

// GetLogger returns the logger to be used when parsing the data
func (cfg *Config) GetLogger() logging.Logger {
	if cfg.logger == nil {
		return logging.DefaultLogger()
	}
	return cfg.logger
}

// WithAccountAddressParser sets the account address parser to be used
func (cfg *Config) WithAccountAddressParser(parser types.AccountAddressParser) *Config {
	cfg.accountAddressParser = parser
	return cfg
}

// GetAccountAddressParser returns the account address parser to be used
func (cfg *Config) GetAccountAddressParser() types.AccountAddressParser {
	return cfg.accountAddressParser
}

// MessageFilterBuilder represents a function that takes as input the configuration and returns a message filter
type MessageFilterBuilder func(cfg config.Config) types.MessageFilter

// WithMessageFilterBuilder sets the message filter to be used
func (cfg *Config) WithMessageFilterBuilder(builder MessageFilterBuilder) *Config {
	cfg.messageFilterBuilder = builder
	return cfg
}

// GetMessageFilterBuilder returns the message filter builder to be used
func (cfg *Config) GetMessageFilterBuilder() MessageFilterBuilder {
	if cfg.messageFilterBuilder == nil {
		return func(cfg config.Config) types.MessageFilter {
			return types.DefaultMessageFilter()
		}
	}
	return cfg.messageFilterBuilder
}

// WithTransactionFilter sets the transaction filter to be used
func (cfg *Config) WithTransactionFilter(filter types.TransactionFilter) *Config {
	cfg.transactionFilter = filter
	return cfg
}

// GetTransactionFilter returns the transaction filter to be used
func (cfg *Config) GetTransactionFilter() types.TransactionFilter {
	if cfg.transactionFilter == nil {
		return types.DefaultTransactionFilter()
	}
	return cfg.transactionFilter
}
