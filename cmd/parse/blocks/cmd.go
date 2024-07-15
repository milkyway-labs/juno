package blocks

import (
	"github.com/spf13/cobra"
)

// NewBlocksCmd returns the Cobra command that allows to fix all the things related to blocks
func NewBlocksCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "blocks",
		Short: "Fix things related to blocks and transactions",
	}

	cmd.AddCommand(
		newAllCmd(),
		newMissingCmd(),
	)

	return cmd
}
