package models

import "time"

type Reservation struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	UserID       uint      `gorm:"not null" json:"user_id"`
	ZoneID       uint      `gorm:"not null" json:"zone_id"`
	LicensePlate string    `gorm:"type:varchar(15);not null" json:"license_plate"`
	Status       string    `gorm:"type:varchar(15);default:'active';not null" json:"status"`
	CreatedAt    time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime" json:"updated_at"`

	User User        `gorm:"foreignKey:UserID" json:"user,omitempty"`
	Zone ParkingZone `gorm:"foreignKey:ZoneID" json:"zone,omitempty"`
}
