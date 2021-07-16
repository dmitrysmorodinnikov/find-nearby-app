package usecase_test

import (
	"find-nearby-backend/config"
	"find-nearby-backend/model"
	locationMock "find-nearby-backend/repository/mocks"
	"find-nearby-backend/usecase"

	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/suite"
)

type LocationTestSuite struct {
	suite.Suite
	cfg        config.Config
	usecase    usecase.LocationUsecase
	repository *locationMock.LocationRepository
}

func (suite *LocationTestSuite) SetupTest() {
	suite.cfg = config.LoadConfig()
	suite.repository = &locationMock.LocationRepository{}
	suite.usecase = usecase.NewLocationUsecase(suite.repository)
}

func (suite *LocationTestSuite) TestFindVehicleLocations_WhenRepoReturnsNoError_ShouldReturnNoError() {
	latitude := 45.4211
	longitude := -75.6903
	radius := 10
	limit := 10
	loc1 := model.Location{
		VehicleID: 1,
		Latitude:  45.4211,
		Longitude: -75.6903,
	}
	loc2 := model.Location{
		VehicleID: 2,
		Latitude:  46.4211,
		Longitude: -76.6903,
	}
	expectedLocs := []model.Location{loc1, loc2}
	suite.repository.On("FindVehicleLocations", latitude, longitude, radius, limit).Return(expectedLocs, nil)
	actualLocs, err := suite.usecase.FindVehicleLocations(latitude, longitude, radius, limit)
	suite.NoError(err)
	suite.Equal(expectedLocs, actualLocs)
	suite.repository.AssertExpectations(suite.T())
}

func (suite *LocationTestSuite) TestFindVehicleLocations_WhenRepoReturnsError_ShouldReturnError() {
	latitude := 45.4211
	longitude := -75.6903
	radius := 10
	limit := 10

	err := errors.New("some repo error")
	expectedErr := errors.Wrapf(err, "failed to find the locations within the range")

	suite.repository.On("FindVehicleLocations", latitude, longitude, radius, limit).Return(nil, err)
	actualLocs, actualErr := suite.usecase.FindVehicleLocations(latitude, longitude, radius, limit)
	suite.EqualError(actualErr, expectedErr.Error())
	suite.Nil(actualLocs)
	suite.repository.AssertExpectations(suite.T())
}

func TestUsecase(t *testing.T) {
	suite.Run(t, new(LocationTestSuite))
}
