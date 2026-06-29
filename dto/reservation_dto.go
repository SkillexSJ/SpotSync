package dto

import "time"

// Request DTo

// CreateReservationRequest
type CreateReservationRequest struct {
	ZoneID       uint   `json:"zone_id"       validate:"required"`
	LicensePlate string `json:"license_plate" validate:"required,max=15"`
}

// Response DTOs

// ReservationZoneInfo
type ReservationZoneInfo struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

// ReservationUserInfo
type ReservationUserInfo struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// ReservationResponse
type ReservationResponse struct {
	ID           uint                `json:"id"`
	UserID       uint                `json:"user_id,omitempty"`
	ZoneID       uint                `json:"zone_id,omitempty"`
	LicensePlate string              `json:"license_plate"`
	Status       string              `json:"status"`
	Zone         ReservationZoneInfo `json:"zone,omitempty"`
	CreatedAt    time.Time           `json:"created_at"`
	UpdatedAt    time.Time           `json:"updated_at,omitempty"`
}

// AdminReservationResponse
type AdminReservationResponse struct {
	ID           uint                `json:"id"`
	LicensePlate string              `json:"license_plate"`
	Status       string              `json:"status"`
	User         ReservationUserInfo `json:"user"`
	Zone         ReservationZoneInfo `json:"zone"`
	CreatedAt    time.Time           `json:"created_at"`
	UpdatedAt    time.Time           `json:"updated_at"`
}
