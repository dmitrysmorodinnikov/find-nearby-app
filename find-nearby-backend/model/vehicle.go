package model

// Vehicle represents a vehicle of a certain type (e.g scooter)
// This entity is not used in this implementation.
// The purpose of it being here is to showcase how the system can be extended to display Vehicle details together with its location.
type Vehicle struct {
	ID     int64
	Type   string
	City   string
	Status string
}
