package repository

import (
	"github.com/SkillexSJ/SpotSync/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ReservationRepository struct {
	db *gorm.DB
}

// create repo
func NewReservationRepository(db *gorm.DB) *ReservationRepository {
	return &ReservationRepository{db: db}
}

// create reservation
func (r *ReservationRepository) Create(reservation *models.Reservation) error {
	return r.db.Create(reservation).Error
}

// FindByID
func (r *ReservationRepository) FindByID(id uint) (*models.Reservation, error) {
	var reservation models.Reservation
	err := r.db.Preload("Zone").First(&reservation, id).Error
	if err != nil {
		return nil, err
	}
	return &reservation, nil
}

// FindByUserID with preload zone
func (r *ReservationRepository) FindByUserID(userID uint) ([]models.Reservation, error) {
	var reservations []models.Reservation
	err := r.db.Preload("Zone").
		Where("user_id = ?", userID).
		Order("created_at DESC").
		Find(&reservations).Error
	return reservations, err
}

// FindAll with preload user and zone
func (r *ReservationRepository) FindAll() ([]models.Reservation, error) {
	var reservations []models.Reservation
	err := r.db.Preload("User").Preload("Zone").
		Order("created_at DESC").
		Find(&reservations).Error
	return reservations, err
}

// UpdateStatus
func (r *ReservationRepository) UpdateStatus(id uint, status string) error {
	return r.db.Model(&models.Reservation{}).
		Where("id = ?", id).
		Update("status", status).Error
}

// CRITICAL RESERVATION
func (r *ReservationRepository) CreateWithLock(reservation *models.Reservation) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		// Lock  parking zone row
		var zone models.ParkingZone
		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).
			First(&zone, reservation.ZoneID).Error; err != nil {
			return err
		}

		// Count current active reservations for this zone
		var activeCount int64
		if err := tx.Model(&models.Reservation{}).
			Where("zone_id = ? AND status = ?", reservation.ZoneID, "active").
			Count(&activeCount).Error; err != nil {
			return err
		}

		//  Check capacity
		if activeCount >= int64(zone.TotalCapacity) {
			return ErrZoneFull
		}

		//  Create the reservation
		if err := tx.Create(reservation).Error; err != nil {
			return err
		}

		return nil
	})
}
