package init

import (
	"fmt"
	"os"

	cmdtypes "github.com/forbole/juno/v5/types/cmd"
	initcmdtypes "github.com/forbole/juno/v5/types/cmd/init"

	"github.com/spf13/cobra"
)

const (
	flagReplace = "replace"
)

// NewInitCmd returns the command that should be run in order to properly initialize Juno
func NewInitCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "init",
		Short: "Initializes the configuration files",
		RunE: func(cmd *cobra.Command, args []string) error {
			// Get the cmd context
			cmdContext := cmdtypes.GetCmdContext(cmd)

			// Create the config path if not present
			if _, err := os.Stat(cmdContext.Home()); os.IsNotExist(err) {
				err = os.MkdirAll(cmdContext.Home(), os.ModePerm)
				if err != nil {
					return err
				}
			}

			replace, err := cmd.Flags().GetBool(flagReplace)
			if err != nil {
				return err
			}

			// Get the config file
			configFilePath := cmdContext.GetConfigFilePath()
			file, _ := os.Stat(configFilePath)

			// Check if the file exists and replace is false
			if file != nil && !replace {
				return fmt.Errorf(
					"configuration file already present at %s. If you wish to overwrite it, use the --%s flag",
					configFilePath, flagReplace)
			}

			// Get the config from the flags
			yamlCfg := cmdContext.GetConfig().GetInitConfig().GetConfigCreator()(cmd)
			return writeConfig(yamlCfg, configFilePath)
		},
	}

	cmd.Flags().Bool(flagReplace, false, "overrides any existing configuration")

	return cmd
}

// writeConfig allows to write the given configuration into the file present at the given path
func writeConfig(cfg initcmdtypes.WritableConfig, path string) error {
	bz, err := cfg.GetBytes()
	if err != nil {
		return err
	}

	return os.WriteFile(path, bz, 0600)
}
