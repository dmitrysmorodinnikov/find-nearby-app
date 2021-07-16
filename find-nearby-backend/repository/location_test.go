package repository_test

import (
	"find-nearby-backend/config"
	"find-nearby-backend/database"
	"find-nearby-backend/logger"
	"find-nearby-backend/model"
	"find-nearby-backend/repository"

	"testing"

	"github.com/golang-migrate/migrate"
	_ "github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/golang-migrate/migrate/source/file"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/suite"
)

type RepositoryTestSuite struct {
	suite.Suite
	db          *sqlx.DB
	dbMigration *migrate.Migrate
	repository  repository.LocationRepository
	originLat   float64
	originLng   float64
}

func (s *RepositoryTestSuite) SetupSuite() {
	cfg := config.LoadConfig()
	log := logger.New(cfg.LogLevel(), cfg.LogFormat())
	s.db, _ = database.New(cfg, log)
	m, err := migrate.New("file://../database/migrations", cfg.DatabaseConnectionURL())
	s.Require().NoError(err)
	s.dbMigration = m
	s.repository = repository.NewPostgresLocationRepository(s.db)
	s.originLat = 1.305649
	s.originLng = 103.926768
}

func (s *RepositoryTestSuite) SetupTest() {
	s.Require().NoError(s.migrateDB(true))
}

func (s *RepositoryTestSuite) TearDownTest() {
	s.Require().NoError(s.migrateDB(false))
}

func (s *RepositoryTestSuite) migrateDB(up bool) error {
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

func (s *RepositoryTestSuite) TestFindVehicleLocations_WhenLimitIsTwo_ShouldReturnTwoClosestLocations() {
	err := s.insertLocations()
	s.Require().NoError(err)

	candidateLocations := getData()
	actualLocations, err := s.repository.FindVehicleLocations(s.originLat, s.originLng, 1000, 2)
	s.Assert().NoError(err)
	s.Assert().Equal(2, len(actualLocations))
	s.Assert().Equal(candidateLocations[0].VehicleID, actualLocations[0].VehicleID)
	s.Assert().Equal(candidateLocations[0].Latitude, actualLocations[0].Latitude)
	s.Assert().Equal(candidateLocations[0].Longitude, actualLocations[0].Longitude)

	s.Assert().Equal(candidateLocations[1].VehicleID, actualLocations[1].VehicleID)
	s.Assert().Equal(candidateLocations[1].Latitude, actualLocations[1].Latitude)
	s.Assert().Equal(candidateLocations[1].Longitude, actualLocations[1].Longitude)
}

func (s *RepositoryTestSuite) TestFindVehicleLocations_WhenRadiusIsTooSmall_ShouldReturnZeroLocations() {
	err := s.insertLocations()
	s.Require().NoError(err)

	actualLocations, err := s.repository.FindVehicleLocations(s.originLat, s.originLng, 1, 20)
	s.Assert().NoError(err)
	s.Assert().Equal(0, len(actualLocations))
}

func (s *RepositoryTestSuite) TestFindVehicleLocations_WhenRadiusAndLimitAreBigEnough_ShouldReturnAllLocations() {
	err := s.insertLocations()
	s.Require().NoError(err)

	candidateLocations := getData()
	actualLocations, err := s.repository.FindVehicleLocations(s.originLat, s.originLng, 3000, 100)
	s.Assert().NoError(err)
	s.Assert().Equal(5, len(actualLocations))
	s.Assert().Equal(candidateLocations[0].VehicleID, actualLocations[0].VehicleID)
	s.Assert().Equal(candidateLocations[0].Latitude, actualLocations[0].Latitude)
	s.Assert().Equal(candidateLocations[0].Longitude, actualLocations[0].Longitude)

	s.Assert().Equal(candidateLocations[1].VehicleID, actualLocations[1].VehicleID)
	s.Assert().Equal(candidateLocations[1].Latitude, actualLocations[1].Latitude)
	s.Assert().Equal(candidateLocations[1].Longitude, actualLocations[1].Longitude)

	s.Assert().Equal(candidateLocations[2].VehicleID, actualLocations[2].VehicleID)
	s.Assert().Equal(candidateLocations[2].Latitude, actualLocations[2].Latitude)
	s.Assert().Equal(candidateLocations[2].Longitude, actualLocations[2].Longitude)

	s.Assert().Equal(candidateLocations[3].VehicleID, actualLocations[3].VehicleID)
	s.Assert().Equal(candidateLocations[3].Latitude, actualLocations[3].Latitude)
	s.Assert().Equal(candidateLocations[3].Longitude, actualLocations[3].Longitude)

	s.Assert().Equal(candidateLocations[4].VehicleID, actualLocations[4].VehicleID)
	s.Assert().Equal(candidateLocations[4].Latitude, actualLocations[4].Latitude)
	s.Assert().Equal(candidateLocations[4].Longitude, actualLocations[4].Longitude)
}

func (s *RepositoryTestSuite) TestFindVehicleLocations_WhenDbReturnsError_ShouldReturnError() {
	err := s.insertLocations()
	s.Require().NoError(err)

	actualLocations, err := s.repository.FindVehicleLocations(s.originLat, s.originLng, 3000, -2)
	s.Assert().Error(err)
	s.Assert().Nil(actualLocations)
}

func (s *RepositoryTestSuite) insertLocations() error {
	locations := getData()
	query := `INSERT INTO locations (vehicle_id, location) VALUES ($1, st_setsrid(st_makepoint($2, $3), 4326))`
	for i := 0; i < len(locations); i++ {
		_, err := s.db.Exec(query, locations[i].VehicleID, locations[i].Longitude, locations[i].Latitude)
		if err != nil {
			return err
		}
	}
	return nil
}

func getData() []model.Location {
	//data is sorted by distance from the first point
	loc1 := model.Location{VehicleID: 2, Longitude: 103.927337, Latitude: 1.306002} //0.07km
	loc2 := model.Location{VehicleID: 3, Longitude: 103.927858, Latitude: 1.306254}
	loc3 := model.Location{VehicleID: 4, Longitude: 103.928515, Latitude: 1.306598}
	loc4 := model.Location{VehicleID: 5, Longitude: 103.928938, Latitude: 1.306799}
	loc5 := model.Location{VehicleID: 6, Longitude: 103.947878, Latitude: 1.311528} //2.44km

	return []model.Location{loc1, loc2, loc3, loc4, loc5}
}

func TestRepository(t *testing.T) {
	suite.Run(t, new(RepositoryTestSuite))
}
