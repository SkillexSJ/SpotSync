package repository

import (
	"github.com/SkillexSJ/SpotSync/models"
	"gorm.io/gorm"
)

type ZoneRepository struct {
	db *gorm.DB
}

// create repo
func NewZoneRepository(db *gorm.DB) *ZoneRepository {
	return &ZoneRepository{db: db}
}

// create zone
func (r *ZoneRepository) Create(zone *models.ParkingZone) error {
	return r.db.Create(zone).Error
}

// FindAll
func (r *ZoneRepository) FindAll() ([]models.ParkingZone, error) {
	var zones []models.ParkingZone
	err := r.db.Find(&zones).Error
	return zones, err
}

// FindByID
func (r *ZoneRepository) FindByID(id uint) (*models.ParkingZone, error) {
	var zone models.ParkingZone
	err := r.db.First(&zone, id).Error
	if err != nil {
		return nil, err
	}
	return &zone, nil
}

// Update zone
func (r *ZoneRepository) Update(zone *models.ParkingZone) error {
	return r.db.Save(zone).Error
}

// Delete zone
func (r *ZoneRepository) Delete(id uint) error {
	return r.db.Delete(&models.ParkingZone{}, id).Error
}

// CountActiveReservations
func (r *ZoneRepository) CountActiveReservations(zoneID uint) (int64, error) {
	var count int64
	err := r.db.Model(&models.Reservation{}).
		Where("zone_id = ? AND status = ?", zoneID, "active").
		Count(&count).Error
	return count, err
}
