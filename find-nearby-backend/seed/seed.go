package seed

import (
	"find-nearby-backend/config"
	"find-nearby-backend/database"
	"find-nearby-backend/logger"
	"find-nearby-backend/model"

	"encoding/csv"
	"os"
	"strconv"
	"strings"

	"github.com/golang-migrate/migrate"
	_ "github.com/golang-migrate/migrate/database/postgres" // required
	_ "github.com/golang-migrate/migrate/source/file"       // required
	"github.com/jmoiron/sqlx"
)

type Seed struct {
	db          *sqlx.DB
	dbMigration *migrate.Migrate
}

func NewSeed(cfg config.Config) *Seed {
	log := logger.New(cfg.LogLevel(), cfg.LogFormat())
	db, _ := database.New(cfg, log)
	m, err := migrate.New("file://./database/migrations", cfg.DatabaseConnectionURL())
	if err != nil {
		return nil
	}
	return &Seed{
		db:          db,
		dbMigration: m,
	}
}

func (s *Seed) Generate() error {
	err := s.migrateDB(true)
	if err != nil {
		return err
	}
	var generatedLocations []model.Location

	f, err := os.Open("./seed/locations.csv")
	if err != nil {
		return err
	}
	defer f.Close()
	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return err
	}
	for i, line := range lines {
		strs := strings.Split(line[0], " ")
		lat, _ := strconv.ParseFloat(strs[0], 64)
		lng, _ := strconv.ParseFloat(strs[1], 64)
		loc := model.Location{
			VehicleID: int64(i),
			Latitude:  lat,
			Longitude: lng,
		}
		generatedLocations = append(generatedLocations, loc)
	}
	return s.insertLocations(generatedLocations)
}

func (s *Seed) migrateDB(up bool) error {
	if up {
		if err := s.dbMigration.Up(); err != nil && err != migrate.ErrNoChange {
			return err
		}
	} else {
		if err := s.dbMigration.Down(); err != nil && err != migrate.ErrNoChange {
			return err
		}
	}
	return nil
}

func (s *Seed) insertLocations(locations []model.Location) error {
	query := `INSERT INTO locations (vehicle_id, location) VALUES ($1, st_setsrid(st_makepoint($2, $3), 4326))`
	for i := 0; i < len(locations); i++ {
		_, err := s.db.Exec(query, locations[i].VehicleID, locations[i].Longitude, locations[i].Latitude)
		if err != nil {
			return err
		}
	}
	return nil
}
