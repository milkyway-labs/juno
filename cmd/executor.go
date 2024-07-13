package cmd

import (
	initcmd "github.com/forbole/juno/v5/cmd/init"
	migratecmd "github.com/forbole/juno/v5/cmd/migrate"
	parsecmd "github.com/forbole/juno/v5/cmd/parse"
	startcmd "github.com/forbole/juno/v5/cmd/start"
	cmdtypes "github.com/forbole/juno/v5/types/cmd"

	"github.com/cometbft/cometbft/libs/cli"
)

// BuildDefaultExecutor allows to build an Executor containing a root command that
// has the provided name and description and the default version and parse sub-commands implementations.
//
// registrar will be used to register custom modules. Be sure to provide an implementation that returns all
// the modules that you want to use. If you don't want any custom module, use modules.EmptyRegistrar.
//
// setupCfg method will be used to customize the SDK configuration. If you don't want any customization
// you can use the config.DefaultConfigSetup variable.
//
// encodingConfigBuilder is used to provide a codec that will later be used to deserialize the
// transaction messages. Make sure you register all the types you need properly.
//
// dbBuilder is used to provide the database that will be used to save the data. If you don't have any
// particular need, you can use the Create variable to build a default database instance.
func BuildDefaultExecutor(config *cmdtypes.Config) cli.Executor {
	rootCmd := cmdtypes.RootCmd(config)

	rootCmd.AddCommand(
		VersionCmd(),
		initcmd.NewInitCmd(),
		parsecmd.NewParseCmd(),
		startcmd.NewStartCmd(),
		migratecmd.NewMigrateCmd(config.GetName()),
	)

	return cmdtypes.PrepareRootCmd(rootCmd)
}
