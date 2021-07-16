package server_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"find-nearby-backend/logger"
	"find-nearby-backend/model"
	"find-nearby-backend/server"
	usecaseMocks "find-nearby-backend/usecase/mocks"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestHandler_Ping(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(echo.GET, "/ping", bytes.NewReader(nil))
	req.Header.Set(echo.HeaderContentType, echo.MIMETextPlainCharsetUTF8)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := server.NewHandler(nil, nil).Ping(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, echo.MIMETextPlainCharsetUTF8, rec.Header().Get("Content-Type"))
	assert.Equal(t, "pong", rec.Body.String())
}

func TestHandler_FindLocations_Success(t *testing.T) {
	lat := 45.13
	lng := 23.23
	radius := 10
	limit := 20
	expectedLocations := []model.Location{
		{VehicleID: 1, Latitude: 12.12, Longitude: 12.12, Distance: 10},
		{VehicleID: 2, Latitude: 22.22, Longitude: 22.22, Distance: 20},
	}
	expectedResponse := server.FindLocationsResponse{
		Data:    expectedLocations,
		Success: true,
		Error:   server.ErrorResponse{},
	}

	e := echo.New()
	url := fmt.Sprintf("/locations/find?latitude=%f&longitude=%f&radius=%d&limit=%d", lat, lng, radius, limit)
	req := httptest.NewRequest(echo.GET, url, bytes.NewReader(nil))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	cfg := config.LoadConfig()
	log := logger.New(cfg.LogLevel(), cfg.LogFormat())

	locationsUsecaseMock := new(usecaseMocks.LocationUsecase)
	locationsUsecaseMock.On("FindVehicleLocations", lat, lng, radius, limit).Return(expectedLocations, nil)
	server.NewHandler(log, locationsUsecaseMock).FindLocations(c)
	assert.Equal(t, http.StatusOK, rec.Code)

	resp := server.FindLocationsResponse{}
	assert.NoError(t, json.Unmarshal(rec.Body.Bytes(), &resp))
	assert.Equal(t, expectedResponse, resp)
	locationsUsecaseMock.AssertExpectations(t)
}

func TestHandler_FindLocations_WhenNoLatitudeParam_ShouldReturn400(t *testing.T) {
	lng := 23.23
	radius := 10
	limit := 20
	expectedErr := errors.New("latitude is a required param")
	expectedResponse := server.FindLocationsResponse{
		Data:    nil,
		Success: false,
		Error: server.ErrorResponse{
			Code:    "400",
			Message: expectedErr.Error(),
		},
	}

	e := echo.New()
	url := fmt.Sprintf("/locations/find?longitude=%f&radius=%d&limit=%d", lng, radius, limit)
	req := httptest.NewRequest(echo.GET, url, bytes.NewReader(nil))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	cfg := config.LoadConfig()
	log := logger.New(cfg.LogLevel(), cfg.LogFormat())

	locationsUsecaseMock := new(usecaseMocks.LocationUsecase)
	server.NewHandler(log, locationsUsecaseMock).FindLocations(c)
	assert.Equal(t, http.StatusBadRequest, rec.Code)

	resp := server.FindLocationsResponse{}
	assert.NoError(t, json.Unmarshal(rec.Body.Bytes(), &resp))
	assert.Equal(t, expectedResponse, resp)
	locationsUsecaseMock.AssertNotCalled(t, "FindVehicleLocations", mock.Anything, mock.Anything, mock.Anything, mock.Anything)
}

func TestHandler_FindLocations_WhenNoLongitudeParam_ShouldReturn400(t *testing.T) {
	lat := 23.23
	radius := 10
	limit := 20
	expectedErr := errors.New("longitude is a required param")
	expectedResponse := server.FindLocationsResponse{
		Data:    nil,
		Success: false,
		Error: server.ErrorResponse{
			Code:    "400",
			Message: expectedErr.Error(),
		},
	}

	e := echo.New()
	url := fmt.Sprintf("/locations/find?latitude=%f&radius=%d&limit=%d", lat, radius, limit)
	req := httptest.NewRequest(echo.GET, url, bytes.NewReader(nil))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	cfg := config.LoadConfig()
	log := logger.New(cfg.LogLevel(), cfg.LogFormat())

	locationsUsecaseMock := new(usecaseMocks.LocationUsecase)
	server.NewHandler(log, locationsUsecaseMock).FindLocations(c)
	assert.Equal(t, http.StatusBadRequest, rec.Code)

	resp := server.FindLocationsResponse{}
	assert.NoError(t, json.Unmarshal(rec.Body.Bytes(), &resp))
	assert.Equal(t, expectedResponse, resp)
	locationsUsecaseMock.AssertNotCalled(t, "FindVehicleLocations", mock.Anything, mock.Anything, mock.Anything, mock.Anything)
}

func TestHandler_FindLocations_WhenNoRadiusParam_ShouldReturn400(t *testing.T) {
	lat := 23.22
	lng := 23.22
	limit := 20
	expectedErr := errors.New("radius is a required param")
	expectedResponse := server.FindLocationsResponse{
		Data:    nil,
		Success: false,
		Error: server.ErrorResponse{
			Code:    "400",
			Message: expectedErr.Error(),
		},
	}

	e := echo.New()
	url := fmt.Sprintf("/locations/find?latitude=%f&longitude=%f&limit=%d", lat, lng, limit)
	req := httptest.NewRequest(echo.GET, url, bytes.NewReader(nil))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	cfg := config.LoadConfig()
	log := logger.New(cfg.LogLevel(), cfg.LogFormat())

	locationsUsecaseMock := new(usecaseMocks.LocationUsecase)
	server.NewHandler(log, locationsUsecaseMock).FindLocations(c)
	assert.Equal(t, http.StatusBadRequest, rec.Code)

	resp := server.FindLocationsResponse{}
	assert.NoError(t, json.Unmarshal(rec.Body.Bytes(), &resp))
	assert.Equal(t, expectedResponse, resp)
	locationsUsecaseMock.AssertNotCalled(t, "FindVehicleLocations", mock.Anything, mock.Anything, mock.Anything, mock.Anything)
}

func TestHandler_FindLocations_WhenNoLimitParam_ShouldReturn400(t *testing.T) {
	lat := 23.22
	lng := 23.22
	radius := 20
	expectedErr := errors.New("limit is a required param")
	expectedResponse := server.FindLocationsResponse{
		Data:    nil,
		Success: false,
		Error: server.ErrorResponse{
			Code:    "400",
			Message: expectedErr.Error(),
		},
	}

	e := echo.New()
	url := fmt.Sprintf("/locations/find?latitude=%f&longitude=%f&radius=%d", lat, lng, radius)
	req := httptest.NewRequest(echo.GET, url, bytes.NewReader(nil))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	cfg := config.LoadConfig()
	log := logger.New(cfg.LogLevel(), cfg.LogFormat())

	locationsUsecaseMock := new(usecaseMocks.LocationUsecase)
	server.NewHandler(log, locationsUsecaseMock).FindLocations(c)
	assert.Equal(t, http.StatusBadRequest, rec.Code)

	resp := server.FindLocationsResponse{}
	assert.NoError(t, json.Unmarshal(rec.Body.Bytes(), &resp))
	assert.Equal(t, expectedResponse, resp)
	locationsUsecaseMock.AssertNotCalled(t, "FindVehicleLocations", mock.Anything, mock.Anything, mock.Anything, mock.Anything)
}

func TestHandler_FindLocations_WhenInvalidLatitude_ShouldReturn400(t *testing.T) {
	lat := 93.23
	lng := 23.23
	radius := 10
	limit := 20
	expectedErr := fmt.Errorf("invalid latitude: %f; latitude must be between -/+ 90", lat)
	expectedResponse := server.FindLocationsResponse{
		Data:    nil,
		Success: false,
		Error: server.ErrorResponse{
			Code:    "400",
			Message: expectedErr.Error(),
		},
	}

	e := echo.New()
	url := fmt.Sprintf("/locations/find?latitude=%f&longitude=%f&radius=%d&limit=%d", lat, lng, radius, limit)
	req := httptest.NewRequest(echo.GET, url, bytes.NewReader(nil))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	cfg := config.LoadConfig()
	log := logger.New(cfg.LogLevel(), cfg.LogFormat())

	locationsUsecaseMock := new(usecaseMocks.LocationUsecase)
	server.NewHandler(log, locationsUsecaseMock).FindLocations(c)
	assert.Equal(t, http.StatusBadRequest, rec.Code)

	resp := server.FindLocationsResponse{}
	assert.NoError(t, json.Unmarshal(rec.Body.Bytes(), &resp))
	assert.Equal(t, expectedResponse, resp)
	locationsUsecaseMock.AssertNotCalled(t, "FindVehicleLocations", mock.Anything, mock.Anything, mock.Anything, mock.Anything)
}

func TestHandler_FindLocations_WhenInvalidLongitude_ShouldReturn400(t *testing.T) {
	lat := 23.23
	lng := -181.1
	radius := 10
	limit := 20
	expectedErr := fmt.Errorf("invalid longitude: %f; longitude must be between -/+ 180", lng)
	expectedResponse := server.FindLocationsResponse{
		Data:    nil,
		Success: false,
		Error: server.ErrorResponse{
			Code:    "400",
			Message: expectedErr.Error(),
		},
	}

	e := echo.New()
	url := fmt.Sprintf("/locations/find?latitude=%f&longitude=%f&radius=%d&limit=%d", lat, lng, radius, limit)
	req := httptest.NewRequest(echo.GET, url, bytes.NewReader(nil))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	cfg := config.LoadConfig()
	log := logger.New(cfg.LogLevel(), cfg.LogFormat())

	locationsUsecaseMock := new(usecaseMocks.LocationUsecase)
	server.NewHandler(log, locationsUsecaseMock).FindLocations(c)
	assert.Equal(t, http.StatusBadRequest, rec.Code)

	resp := server.FindLocationsResponse{}
	assert.NoError(t, json.Unmarshal(rec.Body.Bytes(), &resp))
	assert.Equal(t, expectedResponse, resp)
	locationsUsecaseMock.AssertNotCalled(t, "FindVehicleLocations", mock.Anything, mock.Anything, mock.Anything, mock.Anything)
}

func TestHandler_FindLocations_WhenInvalidRadius_ShouldReturn400(t *testing.T) {
	lat := 23.23
	lng := -23.1
	radius := -10
	limit := 20
	expectedErr := fmt.Errorf("invalid radius: %d; radius must be a positive int32", radius)
	expectedResponse := server.FindLocationsResponse{
		Data:    nil,
		Success: false,
		Error: server.ErrorResponse{
			Code:    "400",
			Message: expectedErr.Error(),
		},
	}

	e := echo.New()
	url := fmt.Sprintf("/locations/find?latitude=%f&longitude=%f&radius=%d&limit=%d", lat, lng, radius, limit)
	req := httptest.NewRequest(echo.GET, url, bytes.NewReader(nil))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	cfg := config.LoadConfig()
	log := logger.New(cfg.LogLevel(), cfg.LogFormat())

	locationsUsecaseMock := new(usecaseMocks.LocationUsecase)
	server.NewHandler(log, locationsUsecaseMock).FindLocations(c)
	assert.Equal(t, http.StatusBadRequest, rec.Code)

	resp := server.FindLocationsResponse{}
	assert.NoError(t, json.Unmarshal(rec.Body.Bytes(), &resp))
	assert.Equal(t, expectedResponse, resp)
	locationsUsecaseMock.AssertNotCalled(t, "FindVehicleLocations", mock.Anything, mock.Anything, mock.Anything, mock.Anything)
}

func TestHandler_FindLocations_WhenInvalidLimit_ShouldReturn400(t *testing.T) {
	lat := 23.23
	lng := -23.1
	radius := 10
	limit := -20
	expectedErr := fmt.Errorf("invalid limit: %d; limit must be a positive int32", limit)
	expectedResponse := server.FindLocationsResponse{
		Data:    nil,
		Success: false,
		Error: server.ErrorResponse{
			Code:    "400",
			Message: expectedErr.Error(),
		},
	}

	e := echo.New()
	url := fmt.Sprintf("/locations/find?latitude=%f&longitude=%f&radius=%d&limit=%d", lat, lng, radius, limit)
	req := httptest.NewRequest(echo.GET, url, bytes.NewReader(nil))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	cfg := config.LoadConfig()
	log := logger.New(cfg.LogLevel(), cfg.LogFormat())

	locationsUsecaseMock := new(usecaseMocks.LocationUsecase)
	server.NewHandler(log, locationsUsecaseMock).FindLocations(c)
	assert.Equal(t, http.StatusBadRequest, rec.Code)

	resp := server.FindLocationsResponse{}
	assert.NoError(t, json.Unmarshal(rec.Body.Bytes(), &resp))
	assert.Equal(t, expectedResponse, resp)
	locationsUsecaseMock.AssertNotCalled(t, "FindVehicleLocations", mock.Anything, mock.Anything, mock.Anything, mock.Anything)
}

func TestHandler_FindLocations_WhenUsecaseReturnsError_ShouldReturn500(t *testing.T) {
	lat := 45.13
	lng := 23.23
	radius := 10
	limit := 20
	expectedErr := errors.New("usecase error")
	expectedResponse := server.FindLocationsResponse{
		Data:    nil,
		Success: false,
		Error: server.ErrorResponse{
			Code:    "500",
			Message: expectedErr.Error(),
		},
	}

	e := echo.New()
	url := fmt.Sprintf("/locations/find?latitude=%f&longitude=%f&radius=%d&limit=%d", lat, lng, radius, limit)
	req := httptest.NewRequest(echo.GET, url, bytes.NewReader(nil))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	cfg := config.LoadConfig()
	log := logger.New(cfg.LogLevel(), cfg.LogFormat())

	locationsUsecaseMock := new(usecaseMocks.LocationUsecase)
	locationsUsecaseMock.On("FindVehicleLocations", lat, lng, radius, limit).Return(nil, expectedErr)
	server.NewHandler(log, locationsUsecaseMock).FindLocations(c)
	assert.Equal(t, http.StatusInternalServerError, rec.Code)

	resp := server.FindLocationsResponse{}
	assert.NoError(t, json.Unmarshal(rec.Body.Bytes(), &resp))
	assert.Equal(t, expectedResponse, resp)
	locationsUsecaseMock.AssertExpectations(t)
}
