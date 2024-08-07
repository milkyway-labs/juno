package transactions

import (
	"fmt"

	cmdtypes "github.com/forbole/juno/v5/types/cmd"

	"github.com/spf13/cobra"

	"github.com/forbole/juno/v5/parser"
)

const (
	flagStart = "start"
	flagEnd   = "end"
)

// newTransactionsCmd returns a Cobra command that allows to fix missing or incomplete transactions in database
func newTransactionsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "all",
		Short: "Parse missing or incomplete transactions",
		Long: fmt.Sprintf(`Refetch missing or incomplete transactions and store them inside the database. 
You can specify a custom height range by using the %s and %s flags. 
`, flagStart, flagEnd),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmdtypes.GetContext(cmd)
			junoCfg, err := ctx.GetJunoConfig()
			if err != nil {
				return err
			}

			parseCtx, err := ctx.GetParseContext()
			if err != nil {
				return err
			}

			worker := parser.NewWorker(parseCtx, nil, 0)

			// Get the flag values
			start, _ := cmd.Flags().GetInt64(flagStart)
			end, _ := cmd.Flags().GetInt64(flagEnd)

			// Get the start height, default to the config's height; use flagStart if set
			startHeight := junoCfg.Parser.StartHeight
			if start > 0 {
				startHeight = start
			}

			// Get the end height, default to the node latest height; use flagEnd if set
			endHeight, err := parseCtx.Node.LatestHeight()
			if err != nil {
				return fmt.Errorf("error while getting chain latest block height: %s", err)
			}
			if end > 0 {
				endHeight = end
			}

			parseCtx.Logger.Info("getting transactions...", "start height", startHeight, "end height", endHeight)
			for k := startHeight; k <= endHeight; k++ {
				parseCtx.Logger.Info("processing transactions...", "height", k)
				err = worker.ProcessTransactions(k)
				if err != nil {
					return fmt.Errorf("error while re-fetching transactions of height %d: %s", k, err)
				}
			}

			return nil
		},
	}

	cmd.Flags().Int64(flagStart, 0, "Height from which to start fetching missing transactions. If 0, the start height inside the config file will be used instead")
	cmd.Flags().Int64(flagEnd, 0, "Height at which to finish fetching missing transactions. If 0, the latest height available inside the node will be used instead")

	return cmd
}
