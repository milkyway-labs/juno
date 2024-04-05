package types

import (
	"fmt"
	"reflect"

	"github.com/forbole/juno/v5/parser"

	nodebuilder "github.com/forbole/juno/v5/node/builder"
	"github.com/forbole/juno/v5/types/config"

	"github.com/forbole/juno/v5/database"

	sdk "github.com/cosmos/cosmos-sdk/types"

	modsregistrar "github.com/forbole/juno/v5/modules/registrar"
)

// GetParserContext setups all the things that can be used to later parse the chain state
func GetParserContext(cfg config.Config, parseConfig *Config) (*parser.Context, error) {
	// Setup the logging
	logger := parseConfig.GetLogger()
	err := logger.SetLogFormat(cfg.Logging.LogFormat)
	if err != nil {
		return nil, fmt.Errorf("error while setting logging format: %s", err)
	}

	err = logger.SetLogLevel(cfg.Logging.LogLevel)
	if err != nil {
		return nil, fmt.Errorf("error while setting logging level: %s", err)
	}

	// Get the account parser and filters
	accountAddressParser := parseConfig.GetAccountAddressParser()
	transactionFilter := parseConfig.GetTransactionFilter()
	messageFilter := parseConfig.GetMessageFilterBuilder()(cfg)

	// Build the codec
	encodingConfig := parseConfig.GetEncodingConfigBuilder()()

	// Setup the SDK configuration
	sdkConfig, sealed := getConfig()
	if !sealed {
		parseConfig.GetSetupConfig()(cfg, sdkConfig)
		sdkConfig.Seal()
	}

	// Create the database
	databaseCtx := database.NewContext(
		cfg.Database,
		encodingConfig,
		logger,
		accountAddressParser,
		transactionFilter,
		messageFilter,
	)
	db, err := parseConfig.GetDBBuilder()(databaseCtx)
	if err != nil {
		return nil, err
	}

	// Create the node
	nodeCtx := nodebuilder.NewContext(encodingConfig, accountAddressParser)
	node, err := nodebuilder.BuildNode(cfg.Node, nodeCtx)
	if err != nil {
		return nil, fmt.Errorf("failed to start client: %s", err)
	}

	// Build the modules
	context := modsregistrar.NewContext(
		cfg,
		sdkConfig,
		encodingConfig,
		db,
		node,
		logger,
		accountAddressParser,
	)
	mods := parseConfig.GetRegistrar().BuildModules(context)
	registeredModules := modsregistrar.GetModules(mods, cfg.Chain.Modules, logger)

	return parser.NewContext(encodingConfig, node, db, logger, registeredModules), nil
}

// getConfig returns the SDK Config instance as well as if it's sealed or not
func getConfig() (config *sdk.Config, sealed bool) {
	sdkConfig := sdk.GetConfig()
	fv := reflect.ValueOf(sdkConfig).Elem().FieldByName("sealed")
	return sdkConfig, fv.Bool()
}
