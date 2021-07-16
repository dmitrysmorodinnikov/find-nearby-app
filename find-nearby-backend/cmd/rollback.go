package cmd

import (
	"find-nearby-backend/config"
	"fmt"
	"log"

	"github.com/golang-migrate/migrate"
	"github.com/spf13/cobra"
)

func newRollbackCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "rollback",
		Short: "Rollback all databases migration",
		Run: func(_ *cobra.Command, args []string) {
			cfg := config.LoadConfig()
			migrateDown(args, cfg)
		},
	}
}

func migrateDown(args []string, cfg config.Config) {
	migrationPath := "file://database/migrations"

	if len(args) > 0 {
		migrationPath = fmt.Sprintf("file://%s", args[0])
	}

	m, err := migrate.New(migrationPath, cfg.DatabaseConnectionURL())
	if err != nil {
		log.Fatal(err)
	}

	if err := m.Down(); err != nil {
		if err == migrate.ErrNoChange {
			fmt.Printf("[FIND-NEARBY-BACKEND][MIGRATION] %s", err.Error())
			return
		}

		log.Fatal(err)
	}
	fmt.Println("[FIND-NEARBY-BACKEND][MIGRATION] Rollback success!")
}
