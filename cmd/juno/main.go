package main

import (
	"os"

	"github.com/forbole/juno/v5/cmd"
	"github.com/forbole/juno/v5/cmd/parse/types"
	"github.com/forbole/juno/v5/modules/registrar"
)

func main() {
	// JunoConfig the runner
	config := cmd.NewConfig("juno").
		WithParseConfig(types.NewConfig().
			WithRegistrar(registrar.NewDefaultRegistrar()),
		)

	// Run the commands and panic on any error
	exec := cmd.BuildDefaultExecutor(config)
	err := exec.Execute()
	if err != nil {
		os.Exit(1)
	}
}
