package main

import (
	"os"

	"github.com/forbole/juno/v5/cmd"
	"github.com/forbole/juno/v5/modules/registrar"
	"github.com/forbole/juno/v5/types"
	cmdtypes "github.com/forbole/juno/v5/types/cmd"
	parsecmdtypes "github.com/forbole/juno/v5/types/cmd/parse"
)

func main() {
	// JunoConfig the runner
	config := cmdtypes.NewConfig("juno").
		WithParseConfig(parsecmdtypes.NewConfig().
			WithEncodingConfigBuilder(func() types.EncodingConfig {
				return types.EncodingConfig{}
			}).
			WithRegistrar(registrar.NewDefaultRegistrar()),
		)

	// Run the commands and panic on any error
	exec := cmd.BuildDefaultExecutor(config)
	err := exec.Execute()
	if err != nil {
		os.Exit(1)
	}
}
