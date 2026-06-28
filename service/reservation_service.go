package service

import (
	"errors"

	"gorm.io/gorm"

	"github.com/SkillexSJ/SpotSync/dto"
	"github.com/SkillexSJ/SpotSync/models"
	"github.com/SkillexSJ/SpotSync/repository"
	"github.com/SkillexSJ/SpotSync/utils"
)

type ReservationService struct {
	reservationRepo *repository.ReservationRepository
	zoneRepo        *repository.ZoneRepository
}

// create reservation service
func NewReservationService(
	reservationRepo *repository.ReservationRepository,
	zoneRepo *repository.ZoneRepository,
) *ReservationService {
	return &ReservationService{
		reservationRepo: reservationRepo,
		zoneRepo:        zoneRepo,
	}
}

// create reservation
func (s *ReservationService) CreateReservation(userID uint, req dto.CreateReservationRequest) (*dto.ReservationResponse, error) {
	// Verify zone exists
	_, err := s.zoneRepo.FindByID(req.ZoneID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, utils.ErrNotFound
		}
		return nil, err
	}

	//  reservation model
	reservation := models.Reservation{
		UserID:       userID,
		ZoneID:       req.ZoneID,
		LicensePlate: req.LicensePlate,
		Status:       "active",
	}

	// concurrency method
	if err := s.reservationRepo.CreateWithLock(&reservation); err != nil {
		if errors.Is(err, repository.ErrZoneFull) {
			return nil, utils.ErrZoneFull
		}
		return nil, err
	}

	response := &dto.ReservationResponse{
		ID:           reservation.ID,
		UserID:       reservation.UserID,
		ZoneID:       reservation.ZoneID,
		LicensePlate: reservation.LicensePlate,
		Status:       reservation.Status,
		CreatedAt:    reservation.CreatedAt,
		UpdatedAt:    reservation.UpdatedAt,
	}

	return response, nil
}

// get my reservations
func (s *ReservationService) GetMyReservations(userID uint) ([]dto.ReservationResponse, error) {
	reservations, err := s.reservationRepo.FindByUserID(userID)
	if err != nil {
		return nil, err
	}

	var responses []dto.ReservationResponse
	for _, r := range reservations {
		responses = append(responses, dto.ReservationResponse{
			ID:           r.ID,
			LicensePlate: r.LicensePlate,
			Status:       r.Status,
			Zone: dto.ReservationZoneInfo{
				ID:   r.Zone.ID,
				Name: r.Zone.Name,
				Type: r.Zone.Type,
			},
			CreatedAt: r.CreatedAt,
		})
	}

	return responses, nil
}

// cancel reservation
func (s *ReservationService) CancelReservation(reservationID uint, userID uint) error {
	// Find reservation
	reservation, err := s.reservationRepo.FindByID(reservationID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return utils.ErrNotFound
		}
		return err
	}

	// Owner check
	if reservation.UserID != userID {
		return utils.ErrForbidden
	}

	return s.reservationRepo.UpdateStatus(reservationID, "cancelled")
}

// get all
func (s *ReservationService) GetAllReservations() ([]dto.AdminReservationResponse, error) {
	reservations, err := s.reservationRepo.FindAll()
	if err != nil {
		return nil, err
	}

	var responses []dto.AdminReservationResponse
	for _, r := range reservations {
		responses = append(responses, dto.AdminReservationResponse{
			ID:           r.ID,
			LicensePlate: r.LicensePlate,
			Status:       r.Status,
			User: dto.ReservationUserInfo{
				ID:    r.User.ID,
				Name:  r.User.Name,
				Email: r.User.Email,
			},
			Zone: dto.ReservationZoneInfo{
				ID:   r.Zone.ID,
				Name: r.Zone.Name,
				Type: r.Zone.Type,
			},
			CreatedAt: r.CreatedAt,
			UpdatedAt: r.UpdatedAt,
		})
	}

	return responses, nil
}
