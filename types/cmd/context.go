package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/forbole/juno/v5/database"
	"github.com/forbole/juno/v5/modules"
	modsregistrar "github.com/forbole/juno/v5/modules/registrar"
	nodebuilder "github.com/forbole/juno/v5/node/builder"
	"github.com/forbole/juno/v5/parser"
	configtypes "github.com/forbole/juno/v5/types/config"
)

type JunoContextKey string

const ContextKey = JunoContextKey("juno.context")

// CmdContext represents the context that will be
// inject the cobra command use to run juno
type CmdContext struct {
	cfg               *Config
	junoCfg           *configtypes.Config
	moduleInitialized bool
	modules           map[string]modules.Module
}

func NewCmdContextFromConfig(cfg *Config) *CmdContext {
	return &CmdContext{
		cfg:               cfg,
		junoCfg:           nil,
		moduleInitialized: false,
		modules:           make(map[string]modules.Module),
	}
}

func InjectCmdContext(cmd *cobra.Command, ctx *CmdContext) {
	cmd.SetContext(context.WithValue(cmd.Context(), ContextKey, ctx))
}

func GetCmdContext(cmd *cobra.Command) *CmdContext {
	ctx := cmd.Context().Value(ContextKey).(*CmdContext)
	if ctx == nil {
		panic("no juno context found, please inject it with the InjectCmdContext function")
	}

	return ctx
}

// GetConfig returns the juno's config
func (ctx *CmdContext) GetConfig() *Config {
	return ctx.cfg
}

func (ctx *CmdContext) GetModule(name string) modules.Module {
	return ctx.modules[name]
}

func (ctx *CmdContext) GetJunoConfig() (*configtypes.Config, error) {
	if ctx.junoCfg == nil {
		file := configtypes.GetConfigFilePath()

		// Make sure the path exists
		if _, err := os.Stat(file); os.IsNotExist(err) {
			return nil, fmt.Errorf("config file does not exist (%s). Make sure you have run the init command", file)
		}

		// Read the config
		junoConfig, err := configtypes.Read(file, ctx.cfg.GetParseConfig().GetConfigParser())
		if err != nil {
			return nil, err
		}
		ctx.junoCfg = &junoConfig
	}

	return ctx.junoCfg, nil
}

// GetParserContext setups all the things that can be used to later parse the chain state
func (ctx *CmdContext) GetParseContext() (*parser.Context, error) {
	parseConfig := ctx.cfg.GetParseConfig()
	cfg, err := ctx.GetJunoConfig()
	if err != nil {
		return nil, err
	}

	// Setup the logging
	logger := parseConfig.GetLogger()
	err = logger.SetLogFormat(cfg.Logging.LogFormat)
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
	messageFilter := parseConfig.GetMessageFilterBuilder()(*cfg)

	// Build the codec
	encodingConfig := parseConfig.GetEncodingConfigBuilder()()
	// Create the database
	databaseCtx := database.NewContext(
		cfg.Database,
		encodingConfig,
		logger,
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
		*cfg,
		encodingConfig,
		db,
		node,
		logger,
		accountAddressParser,
	)
	mods := parseConfig.GetRegistrar().BuildModules(context)
	registeredModules := modsregistrar.GetModules(mods, cfg.Chain.Modules, logger)

	return parser.NewContext(*cfg, encodingConfig, node, db, logger, registeredModules), nil
}
