package cmd

import (
	"find-nearby-backend/config"
	"find-nearby-backend/seed"

	"github.com/spf13/cobra"
)

func newSeedCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "seed",
		Short: "Seed Singapore Locations",
		Run: func(_ *cobra.Command, _ []string) {
			cfg := config.LoadConfig()
			if err := seed.NewSeed(cfg).Generate(); err != nil {
				panic(err.Error())
			}
		},
	}
}
