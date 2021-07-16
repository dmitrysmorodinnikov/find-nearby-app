package server

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"find-nearby-backend/logger"
	"find-nearby-backend/usecase"

	"github.com/labstack/echo"
)

// Handler parses and validates the incoming requests, asks Usecase layer to perform business logic and constructs the responses
type Handler struct {
	logger           logger.Logger
	locationsUsecase usecase.LocationUsecase
}

// NewHandler is a constructor for Handler
func NewHandler(logger logger.Logger, locationsUsecase usecase.LocationUsecase) *Handler {
	return &Handler{
		logger:           logger,
		locationsUsecase: locationsUsecase,
	}
}

// Ping ensures that the server is healthy and is able to respond
func (h *Handler) Ping(c echo.Context) error {
	return c.String(http.StatusOK, "pong")
}

// FindLocations returns nearby vehicle locations
func (h *Handler) FindLocations(c echo.Context) error {
	c.Response().Header().Set("Access-Control-Allow-Origin", "*")
	lat, lng, radius, limit, err := h.getRequestParams(c)
	if err != nil {
		h.logger.Errorf("failed to validate the request, err: %s", err.Error())
		return c.JSON(http.StatusBadRequest, FindLocationsResponse{
			Data:    nil,
			Success: false,
			Error: ErrorResponse{
				Code:    "400",
				Message: err.Error(),
			},
		})
	}
	locations, err := h.locationsUsecase.FindVehicleLocations(lat, lng, radius, limit)
	if err != nil {
		h.logger.ErrorWithTag(err, logger.Fields{
			"msg":    "failed to find vehicle locations",
			"lat":    lat,
			"lng":    lng,
			"radius": radius,
			"limit":  limit,
		})
		return c.JSON(http.StatusInternalServerError, FindLocationsResponse{
			Data:    nil,
			Success: false,
			Error: ErrorResponse{
				Code:    "500",
				Message: err.Error(),
			},
		})
	}
	return c.JSON(http.StatusOK, FindLocationsResponse{
		Data:    locations,
		Success: true,
		Error:   ErrorResponse{},
	})
}

func (h *Handler) getRequestParams(c echo.Context) (float64, float64, int, int, error) {
	lat, err := h.validateLatitude(c.QueryParam("latitude"))
	if err != nil {
		return 0, 0, 0, 0, err
	}
	lng, err := h.validateLongitude(c.QueryParam("longitude"))
	if err != nil {
		return 0, 0, 0, 0, err
	}
	radius, err := h.validateRadius(c.QueryParam("radius"))
	if err != nil {
		return 0, 0, 0, 0, err
	}
	limit, err := h.validateLimit(c.QueryParam("limit"))
	if err != nil {
		return 0, 0, 0, 0, err
	}
	return lat, lng, radius, limit, nil
}

func (h *Handler) validateLatitude(latitude string) (float64, error) {
	if latitude == "" {
		return 0, errors.New("latitude is a required param")
	}
	lat, err := strconv.ParseFloat(latitude, 64)
	if err != nil {
		return 0, fmt.Errorf("failed to parse the latitude value: %v", lat)
	}
	if lat < -90 || lat > 90 {
		return 0, fmt.Errorf("invalid latitude: %f; latitude must be between -/+ 90", lat)
	}
	return lat, nil
}

func (h *Handler) validateLongitude(longitude string) (float64, error) {
	if longitude == "" {
		return 0, errors.New("longitude is a required param")
	}
	lng, err := strconv.ParseFloat(longitude, 64)
	if err != nil {
		return 0, fmt.Errorf("failed to parse the longitude value: %v", lng)
	}
	if lng < -180 || lng > 180 {
		return 0, fmt.Errorf("invalid longitude: %f; longitude must be between -/+ 180", lng)
	}
	return lng, nil
}

func (h *Handler) validateRadius(radius string) (int, error) {
	if radius == "" {
		return 0, errors.New("radius is a required param")
	}
	rad, err := strconv.ParseInt(radius, 10, 32)
	if err != nil {
		return 0, fmt.Errorf("failed to parse the latitude value: %s", radius)
	}
	if rad < 0 {
		return 0, fmt.Errorf("invalid radius: %d; radius must be a positive int32", rad)
	}
	return int(rad), nil
}

func (h *Handler) validateLimit(limit string) (int, error) {
	if limit == "" {
		return 0, errors.New("limit is a required param")
	}
	lim, err := strconv.ParseInt(limit, 10, 32)
	if err != nil {
		return 0, fmt.Errorf("failed to parse the limit value: %s", limit)
	}
	if lim < 0 {
		return 0, fmt.Errorf("invalid limit: %d; limit must be a positive int32", lim)
	}
	return int(lim), nil
}
