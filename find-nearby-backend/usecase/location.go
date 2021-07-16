package usecase

import (
	"find-nearby-backend/model"
	"find-nearby-backend/repository"

	"github.com/pkg/errors"
)

// LocationUsecase is responsible for any location-related business logic
type LocationUsecase interface {
	FindVehicleLocations(latitude, longitude float64, radius, limit int) ([]model.Location, error)
}

type locationUsecase struct {
	locationRepository repository.LocationRepository
}

// NewLocationUsecase is a constructor for locationUsecase
func NewLocationUsecase(locationRepository repository.LocationRepository) LocationUsecase {
	return &locationUsecase{locationRepository: locationRepository}
}

// FindVehicleLocations finds nearby locations
func (l locationUsecase) FindVehicleLocations(latitude, longitude float64, radius, limit int) ([]model.Location, error) {
	locations, err := l.locationRepository.FindVehicleLocations(latitude, longitude, radius, limit)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to find the locations within the range")
	}
	return locations, nil
}
