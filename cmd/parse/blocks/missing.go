package blocks

import (
	"fmt"
	"strconv"

	cmdtypes "github.com/forbole/juno/v5/types/cmd"

	"github.com/spf13/cobra"

	"github.com/forbole/juno/v5/parser"
)

// newMissingCmd returns a Cobra command that allows to fix missing blocks in database
func newMissingCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "missing [start height]",
		Short: "Refetch all the missing heights in the database starting from the given start height",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cmdContext := cmdtypes.GetContext(cmd)
			startHeight, err := strconv.ParseInt(args[0], 10, 64)
			if err != nil {
				return fmt.Errorf("make sure the given start height is a positive integer")
			}

			parseCtx, err := cmdContext.GetParseContext()
			if err != nil {
				return err
			}

			worker := parser.NewWorker(parseCtx, nil, 0)

			dbLastHeight, err := parseCtx.Database.GetLastBlockHeight()
			if err != nil {
				return fmt.Errorf("error while getting DB last block height: %s", err)
			}

			for _, k := range parseCtx.Database.GetMissingHeights(startHeight, dbLastHeight) {
				err = worker.Process(k)
				if err != nil {
					return fmt.Errorf("error while re-fetching block %d: %s", k, err)
				}
			}

			return nil
		},
	}

	return cmd
}
