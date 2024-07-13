package cmd

import (
	"fmt"
	"os"
	"path"

	initcmd "github.com/forbole/juno/v5/cmd/init"
	migratecmd "github.com/forbole/juno/v5/cmd/migrate"
	parsecmd "github.com/forbole/juno/v5/cmd/parse"
	startcmd "github.com/forbole/juno/v5/cmd/start"
	cmdtypes "github.com/forbole/juno/v5/types/cmd"

	"github.com/cometbft/cometbft/libs/cli"
	"github.com/spf13/cobra"
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
	rootCmd := RootCmd(config)

	rootCmd.AddCommand(
		VersionCmd(),
		initcmd.NewInitCmd(),
		parsecmd.NewParseCmd(),
		startcmd.NewStartCmd(),
		migratecmd.NewMigrateCmd(config.GetName()),
	)

	return PrepareRootCmd(rootCmd)
}

// RootCmd allows to build the default root command having the given name
func RootCmd(config *cmdtypes.Config) *cobra.Command {
	name := config.GetName()
	rootCmd := &cobra.Command{
		Use:   name,
		Short: fmt.Sprintf("%s is a Chain SDK-based chain data aggregator and exporter", name),
		Long: fmt.Sprintf(`A Chain chain data aggregator. It improves the chain's data accessibility
by providing an indexed database exposing aggregated resources and models such as blocks, validators, pre-commits, 
transactions, and various aspects of the governance module. 
%s is meant to run with a GraphQL layer on top so that it even further eases the ability for developers and
downstream clients to answer queries such as "What is the average gas cost of a block?" while also allowing
them to compose more aggregate and complex queries.`, name),
	}

	// Set the default home path
	home, _ := os.UserHomeDir()
	defaultConfigPath := path.Join(home, fmt.Sprintf(".%s", config.GetName()))
	rootCmd.PersistentFlags().String(cmdtypes.FlagHome, defaultConfigPath, "Set the home folder of the application, where all files will be stored")

	// Inject the juno context into the cmd context
	cmdtypes.InjectCmdContext(rootCmd, cmdtypes.NewCmdContextFromConfig(config))

	return rootCmd
}

// PrepareRootCmd is meant to prepare the given command binding all the viper flags
func PrepareRootCmd(cmd *cobra.Command) cli.Executor {
	return cli.Executor{Command: cmd, Exit: os.Exit}
}
