package cmd

import (
	"github.com/spf13/cobra"
)

func NewRootCmd() *cobra.Command {
	cli := &cobra.Command{
		Use: "Find Nearby Backend",
	}
	cli.AddCommand(newStartCmd())
	cli.AddCommand(newMigrateCmd())
	cli.AddCommand(newRollbackCmd())
	cli.AddCommand(newSeedCmd())

	return cli
}
