package model

// Location represents the location of the vehicle. The Vehicle can be of any type.
type Location struct {
	VehicleID int64   `db:"vehicle_id" json:"vehicle_id"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Distance  float64 `json:"distance"`
}
