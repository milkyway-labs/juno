package blocks

import (
	"fmt"

	cmdtypes "github.com/forbole/juno/v5/types/cmd"
	"github.com/forbole/juno/v5/utils"

	"github.com/spf13/cobra"

	"github.com/forbole/juno/v5/parser"
)

const (
	flagForce = "force"
	flagStart = "start"
	flagEnd   = "end"
)

// newAllCmd returns a Cobra command that allows to fix missing blocks in database
func newAllCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "all",
		Short: "Reparse blocks and transactions ranged from the given start height to the given end height",
		Long: fmt.Sprintf(`Refetch all the blocks in the specified range and stores them inside the database. 
You can specify a custom blocks range by using the %s and %s flags. 
By default, all the blocks fetched from the node will not be stored inside the database if they are already present. 
You can override this behaviour using the %s flag. If this is set, even the blocks already present inside the database 
will be replaced with the data downloaded from the node.
`, flagStart, flagEnd, flagForce),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmdtypes.GetCmdContext(cmd)
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
			force, _ := cmd.Flags().GetBool(flagForce)

			lastDbBlockHeight, err := parseCtx.Database.GetLastBlockHeight()
			if err != nil {
				return err
			}

			// Compare start height from config file and last block height in database
			// and set higher block as start height
			startHeight := utils.MaxInt64(junoCfg.Parser.StartHeight, lastDbBlockHeight)

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

			parseCtx.Logger.Info("getting blocks and transactions", "start height", startHeight, "end height", endHeight)
			for k := startHeight; k <= endHeight; k++ {
				if force {
					err = worker.Process(k)
				} else {
					err = worker.ProcessIfNotExists(k)
				}

				if err != nil {
					return fmt.Errorf("error while re-fetching block %d: %s", k, err)
				}
			}

			return nil
		},
	}

	cmd.Flags().Bool(flagForce, false, "Whether or not to overwrite any existing ones in database (default false)")
	cmd.Flags().Int64(flagStart, 0, "Height from which to start getting missing blocks. If 0, the start height inside the config will be used instead")
	cmd.Flags().Int64(flagEnd, 0, "Height at which to finish getting missing. If 0, the latest height available inside the node will be used instead")

	return cmd
}
