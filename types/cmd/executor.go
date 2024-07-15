package cmd

import (
	"fmt"
	"os"
	"path"

	"github.com/cometbft/cometbft/libs/cli"
	"github.com/spf13/cobra"
)

// RootCmd allows to build the default root command having the given name
func RootCmd(config *Config) *cobra.Command {
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
	rootCmd.PersistentFlags().String(FlagHome, defaultConfigPath, "Set the home folder of the application, where all files will be stored")

	// Inject the juno context into the cmd context
	InjectContext(rootCmd, NewContextFromConfig(config))

	return rootCmd
}

// PrepareRootCmd is meant to prepare the given command binding all the viper flags
func PrepareRootCmd(cmd *cobra.Command) cli.Executor {
	return cli.Executor{Command: cmd, Exit: os.Exit}
}
