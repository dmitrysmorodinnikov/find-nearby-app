package cmd

import (
	"find-nearby-backend/config"
	"find-nearby-backend/server"

	"github.com/spf13/cobra"
)

func newStartCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "start",
		Short: "Start HTTP Server",
		Run: func(_ *cobra.Command, _ []string) {
			cfg := config.LoadConfig()
			server.Start(cfg)
		},
	}
}
