package parse

import (
	"github.com/spf13/cobra"

	parseblocks "github.com/forbole/juno/v5/cmd/parse/blocks"
	parsegenesis "github.com/forbole/juno/v5/cmd/parse/genesis"
	parsetransactions "github.com/forbole/juno/v5/cmd/parse/transactions"
)

// NewParseCmd returns the Cobra command allowing to parse some chain data without having to re-sync the whole database
func NewParseCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "parse",
		Short: "Parse some data without the need to re-syncing the whole database from scratch",
	}

	cmd.AddCommand(
		parseblocks.NewBlocksCmd(),
		parsegenesis.NewGenesisCmd(),
		parsetransactions.NewTransactionsCmd(),
	)

	return cmd
}
