package repository

import (
	"find-nearby-backend/model"

	"github.com/jmoiron/sqlx"
	geojson "github.com/paulmach/go.geojson"
)

// LocationRepository represents the repository layer for locations
type LocationRepository interface {
	FindVehicleLocations(latitude, longitude float64, radius, limit int) ([]model.Location, error)
}

type postgresLocationRepository struct {
	db *sqlx.DB
}

// NewPostgresLocationRepository is a constructor for postgresLocationRepository
func NewPostgresLocationRepository(db *sqlx.DB) LocationRepository {
	return postgresLocationRepository{db: db}
}

// FindVehicleLocations fetches the nearby locations from the underlying storage
func (p postgresLocationRepository) FindVehicleLocations(latitude, longitude float64, radius, limit int) ([]model.Location, error) {
	var locations []model.Location
	query := `SELECT
 				vehicle_id,
 				st_asgeojson(location) as loc,
				st_distance(geography(location), geography(st_setsrid(st_makepoint($1, $2), 4326))) as distance
 				FROM locations
				WHERE st_within(location, geometry(st_buffer(geography(st_setsrid(st_makepoint($3, $4), 4326)), $5)))
				ORDER BY distance ASC
				LIMIT $6
`
	rows, err := p.db.Queryx(query, longitude, latitude, longitude, latitude, radius, limit)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var vehicleID int64
		var distance float64
		var location geojson.Geometry
		err = rows.Scan(&vehicleID, &location, &distance)
		if err != nil {
			return nil, err
		}
		locations = append(locations, model.Location{
			VehicleID: vehicleID,
			Latitude:  location.Point[1],
			Longitude: location.Point[0],
			Distance:  distance,
		})
	}
	return locations, nil
}
