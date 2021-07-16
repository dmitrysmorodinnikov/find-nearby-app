package server

import "find-nearby-backend/model"

// FindLocationsResponse is a response message
type FindLocationsResponse struct {
	Data    []model.Location `json:"data"`
	Success bool             `json:"success"`
	Error   ErrorResponse    `json:"error"`
}

// ErrorResponse is an error response message
type ErrorResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}
