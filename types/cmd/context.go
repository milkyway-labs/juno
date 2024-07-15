package cmd

import (
	"context"
	"fmt"
	"os"
	"path"

	"github.com/spf13/cobra"

	"github.com/forbole/juno/v5/database"
	"github.com/forbole/juno/v5/logging"
	"github.com/forbole/juno/v5/modules"
	modsregistrar "github.com/forbole/juno/v5/modules/registrar"
	"github.com/forbole/juno/v5/node"
	nodebuilder "github.com/forbole/juno/v5/node/builder"
	"github.com/forbole/juno/v5/parser"
	"github.com/forbole/juno/v5/types"
	configtypes "github.com/forbole/juno/v5/types/config"
)

type JunoContextKey string

const ContextKey = JunoContextKey("juno.context")

// Context represents the context that will be
// inject the cobra command use to run juno
type Context struct {
	cfg     *Config
	junoCfg *configtypes.Config

	// Fields initialized at runtime
	home               string
	logger             *logging.Logger
	database           *database.Database
	node               *node.Node
	modulesInitialized bool
	modules            []modules.Module
}

func NewContextFromConfig(cfg *Config) *Context {
	return &Context{
		cfg:     cfg,
		junoCfg: nil,
	}
}

func InjectContext(cmd *cobra.Command, ctx *Context) {
	cmdContext := cmd.Context()
	if cmdContext == nil {
		cmdContext = context.TODO()
	}
	cmd.SetContext(context.WithValue(cmdContext, ContextKey, ctx))
}

func GetContext(cmd *cobra.Command) *Context {
	var ctx *Context
	currCmd := cmd
	for {
		ctxValue, ok := currCmd.Context().Value(ContextKey).(*Context)
		if !ok {
			currCmd = currCmd.Parent()
			// No more parents
			if currCmd == nil {
				break
			}
		} else {
			ctx = ctxValue
			break
		}
	}
	if ctx == nil {
		panic("no juno context found, please inject it with the InjectCmdContext function")
	}

	// Set the context home path from the cmd flag
	homePath, err := cmd.Flags().GetString(FlagHome)
	if err != nil {
		panic(fmt.Sprintf("can't get context from cmd, cmd don't have the %s flag", FlagHome))
	}
	ctx.SetHome(homePath)

	return ctx
}

func (ctx *Context) Home() string {
	return ctx.home
}

func (ctx *Context) SetHome(home string) {
	ctx.home = home
}

func (ctx *Context) GetConfigFilePath() string {
	if ctx.home == "" {
		panic("Can't get config file path, home path is not set")
	}

	return path.Join(ctx.home, "config.yaml")
}

// GetConfig returns the cmd configuration
func (ctx *Context) GetConfig() *Config {
	return ctx.cfg
}

// GetJunoConfig returns the juno configuration parsed from the config file
func (ctx *Context) GetJunoConfig() (*configtypes.Config, error) {
	if ctx.junoCfg == nil {
		configFilePath := ctx.GetConfigFilePath()

		// Make sure the path exists
		if _, err := os.Stat(configFilePath); os.IsNotExist(err) {
			return nil, fmt.Errorf("config file does not exist (%s). Make sure you have run the init command", configFilePath)
		}

		// Read the config
		junoConfig, err := configtypes.Read(configFilePath, ctx.cfg.GetParseConfig().GetConfigParser())
		if err != nil {
			return nil, err
		}
		ctx.junoCfg = &junoConfig
	}

	return ctx.junoCfg, nil
}

// GetLogger returns the juno logger configured as specified in the juno config
func (ctx *Context) GetLogger() (logging.Logger, error) {
	if ctx.logger == nil {
		logger := ctx.cfg.GetParseConfig().GetLogger()
		junoConfig, err := ctx.GetJunoConfig()
		if err != nil {
			return nil, err
		}

		err = logger.SetLogFormat(junoConfig.Logging.LogFormat)
		if err != nil {
			return nil, fmt.Errorf("error while setting logging format: %s", err)
		}

		err = logger.SetLogLevel(junoConfig.Logging.LogLevel)
		if err != nil {
			return nil, fmt.Errorf("error while setting logging level: %s", err)
		}

		ctx.logger = &logger
	}

	return *ctx.logger, nil
}

// GetAccountAddressParser returns the account address parser to be used
func (ctx *Context) GetAccountAddressParser() types.AccountAddressParser {
	return ctx.cfg.GetParseConfig().GetAccountAddressParser()
}

// GetTransactionFilter returns the transaction filter configured in the
// parse configuration
func (ctx *Context) GetTransactionFilter() types.TransactionFilter {
	return ctx.cfg.GetParseConfig().GetTransactionFilter()
}

// GetMessageFilter returns the message filter configured in the
// parse configuration
func (ctx *Context) GetMessageFilter() (types.MessageFilter, error) {
	junoConfig, err := ctx.GetJunoConfig()
	if err != nil {
		return nil, err
	}

	msgFilter := ctx.cfg.GetParseConfig().GetMessageFilterBuilder()(*junoConfig)
	return msgFilter, nil
}

// GetEncodingConfig returns the encoding configuration
func (ctx *Context) GetEncodingConfig() types.EncodingConfig {
	return ctx.cfg.GetParseConfig().GetEncodingConfigBuilder()()
}

// GetDatabase returns the database used by juno to store the data
func (ctx *Context) GetDatabase() (database.Database, error) {
	if ctx.database == nil {
		junoCfg, err := ctx.GetJunoConfig()
		if err != nil {
			return nil, err
		}

		logger, err := ctx.GetLogger()
		if err != nil {
			return nil, err
		}

		msgFilter, err := ctx.GetMessageFilter()
		if err != nil {
			return nil, err
		}

		// Create the database
		databaseCtx := database.NewContext(
			junoCfg.Database,
			ctx.GetEncodingConfig(),
			logger,
			ctx.GetTransactionFilter(),
			msgFilter,
		)
		db, err := ctx.cfg.GetParseConfig().GetDBBuilder()(databaseCtx)
		if err != nil {
			return nil, err
		}
		ctx.database = &db
	}

	return *ctx.database, nil
}

// GetNode returns the node from which we parse the data
func (ctx *Context) GetNode() (node.Node, error) {
	if ctx.node == nil {
		junoCfg, err := ctx.GetJunoConfig()
		if err != nil {
			return nil, err
		}

		// Create the node
		nodeCtx := nodebuilder.NewContext(ctx.GetEncodingConfig(), ctx.GetAccountAddressParser())
		node, err := nodebuilder.BuildNode(junoCfg.Node, nodeCtx)
		if err != nil {
			return nil, fmt.Errorf("failed to start client: %s", err)
		}

		ctx.node = &node
	}
	return *ctx.node, nil
}

// GetAllModules returns all the modules that have been registred
func (ctx *Context) GetAllModules() ([]modules.Module, error) {
	if !ctx.modulesInitialized {
		junoConfig, err := ctx.GetJunoConfig()
		if err != nil {
			return nil, err
		}

		logger, err := ctx.GetLogger()
		if err != nil {
			return nil, err
		}

		db, err := ctx.GetDatabase()
		if err != nil {
			return nil, err
		}

		node, err := ctx.GetNode()
		if err != nil {
			return nil, err
		}

		context := modsregistrar.NewContext(
			ctx.Home(),
			*junoConfig,
			ctx.GetEncodingConfig(),
			db,
			node,
			logger,
			ctx.GetAccountAddressParser(),
		)
		ctx.modules = ctx.cfg.GetParseConfig().GetRegistrar().BuildModules(context)
		ctx.modulesInitialized = true
	}

	return ctx.modules, nil
}

// GetModule returns the module with the given name
func (ctx *Context) GetModule(module string) (modules.Module, error) {
	allModules, err := ctx.GetAllModules()
	if err != nil {
		return nil, err
	}

	for _, mod := range allModules {
		if mod.Name() == module {
			return mod, nil
		}
	}

	return nil, fmt.Errorf("module %s not found", module)
}

// GetParseContext setups all the things that can be used to later parse the chain state
func (ctx *Context) GetParseContext() (*parser.Context, error) {
	cfg, err := ctx.GetJunoConfig()
	if err != nil {
		return nil, err
	}

	// Setup the logging
	logger, err := ctx.GetLogger()
	if err != nil {
		return nil, err
	}

	// Get the encoding config
	encodingConfig := ctx.GetEncodingConfig()

	// Get database
	db, err := ctx.GetDatabase()
	if err != nil {
		return nil, err
	}

	// Get the node
	node, err := ctx.GetNode()
	if err != nil {
		return nil, err
	}

	mods, err := ctx.GetAllModules()
	if err != nil {
		return nil, err
	}

	registeredModules := modsregistrar.GetModules(mods, cfg.Chain.Modules, logger)
	return parser.NewContext(*cfg, encodingConfig, node, db, logger, registeredModules), nil
}
