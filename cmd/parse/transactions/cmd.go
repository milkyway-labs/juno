package transactions

import (
	"github.com/spf13/cobra"
)

// NewTransactionsCmd returns the Cobra command that allows to fix missing or incomplete transactions
func NewTransactionsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "transactions",
		Short: "Parse things related to transactions",
	}

	cmd.AddCommand(
		newTransactionsCmd(),
	)

	return cmd
}
