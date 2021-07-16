package cmd

import (
	"find-nearby-backend/config"
	"fmt"
	"log"

	"github.com/golang-migrate/migrate"
	_ "github.com/golang-migrate/migrate/database/postgres" // Need this to talk to postgres
	_ "github.com/golang-migrate/migrate/source/file"
	"github.com/spf13/cobra"
)

func newMigrateCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "migrate",
		Short: "Migrate Database",
		Run: func(_ *cobra.Command, args []string) {
			cfg := config.LoadConfig()
			migrateUp(args, cfg)
		},
	}
}

func migrateUp(args []string, cfg config.Config) {
	migrationPath := "file://database/migrations"

	if len(args) > 0 {
		migrationPath = fmt.Sprintf("file://%s", args[0])
	}

	connURL := cfg.DatabaseConnectionURL()
	m, err := migrate.New(migrationPath, connURL)
	if err != nil {
		log.Fatal(err)
	}

	err = m.Up()
	if err != nil {
		if err == migrate.ErrNoChange {
			fmt.Printf("[FIND-NEARBY-BACKEND][MIGRATION] %s \n", err.Error())
			return
		}
		log.Fatal(err)
	}
	fmt.Println("[FIND-NEARBY-BACKEND][MIGRATION] Success!")
}
