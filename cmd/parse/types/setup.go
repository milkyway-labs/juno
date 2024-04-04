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
	// Build the codec
	encodingConfig := parseConfig.GetEncodingConfigBuilder()()

	// Setup the SDK configuration
	sdkConfig, sealed := getConfig()
	if !sealed {
		parseConfig.GetSetupConfig()(cfg, sdkConfig)
		sdkConfig.Seal()
	}

	// Get the db
	databaseCtx := database.NewContext(
		cfg.Database,
		encodingConfig,
		parseConfig.GetLogger(),
		parseConfig.GetAccountAddressParser(),
	)
	db, err := parseConfig.GetDBBuilder()(databaseCtx)
	if err != nil {
		return nil, err
	}

	// Init the client
	nodeCtx := nodebuilder.NewContext(encodingConfig, parseConfig.GetAccountAddressParser())
	node, err := nodebuilder.BuildNode(cfg.Node, nodeCtx)
	if err != nil {
		return nil, fmt.Errorf("failed to start client: %s", err)
	}

	// Setup the logging
	err = parseConfig.GetLogger().SetLogFormat(cfg.Logging.LogFormat)
	if err != nil {
		return nil, fmt.Errorf("error while setting logging format: %s", err)
	}

	err = parseConfig.GetLogger().SetLogLevel(cfg.Logging.LogLevel)
	if err != nil {
		return nil, fmt.Errorf("error while setting logging level: %s", err)
	}

	// Get the modules
	context := modsregistrar.NewContext(
		cfg,
		sdkConfig,
		encodingConfig,
		db,
		node,
		parseConfig.GetLogger(),
		parseConfig.GetAccountAddressParser(),
	)
	mods := parseConfig.GetRegistrar().BuildModules(context)
	registeredModules := modsregistrar.GetModules(mods, cfg.Chain.Modules, parseConfig.GetLogger())

	return parser.NewContext(encodingConfig, node, db, parseConfig.GetLogger(), registeredModules), nil
}

// getConfig returns the SDK Config instance as well as if it's sealed or not
func getConfig() (config *sdk.Config, sealed bool) {
	sdkConfig := sdk.GetConfig()
	fv := reflect.ValueOf(sdkConfig).Elem().FieldByName("sealed")
	return sdkConfig, fv.Bool()
}
