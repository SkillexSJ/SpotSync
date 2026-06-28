package service

import (
	"errors"

	"gorm.io/gorm"

	"github.com/SkillexSJ/SpotSync/dto"
	"github.com/SkillexSJ/SpotSync/models"
	"github.com/SkillexSJ/SpotSync/repository"
	"github.com/SkillexSJ/SpotSync/utils"
)

type ZoneService struct {
	zoneRepo *repository.ZoneRepository
}

// create service
func NewZoneService(zoneRepo *repository.ZoneRepository) *ZoneService {
	return &ZoneService{zoneRepo: zoneRepo}
}

// create zone (admin only)
func (s *ZoneService) CreateZone(req dto.CreateZoneRequest) (*dto.ZoneResponse, error) {
	zone := models.ParkingZone{
		Name:          req.Name,
		Type:          req.Type,
		TotalCapacity: req.TotalCapacity,
		PricePerHour:  req.PricePerHour,
	}

	if err := s.zoneRepo.Create(&zone); err != nil {
		return nil, err
	}

	response := &dto.ZoneResponse{
		ID:             zone.ID,
		Name:           zone.Name,
		Type:           zone.Type,
		TotalCapacity:  zone.TotalCapacity,
		AvailableSpots: zone.TotalCapacity,
		PricePerHour:   zone.PricePerHour,
		CreatedAt:      zone.CreatedAt,
		UpdatedAt:      zone.UpdatedAt,
	}

	return response, nil
}

// get all zones
func (s *ZoneService) GetAllZones() ([]dto.ZoneResponse, error) {
	zones, err := s.zoneRepo.FindAll()
	if err != nil {
		return nil, err
	}

	var responses []dto.ZoneResponse
	for _, zone := range zones {
		// active reservations count
		activeCount, err := s.zoneRepo.CountActiveReservations(zone.ID)
		if err != nil {
			return nil, err
		}

		responses = append(responses, dto.ZoneResponse{
			ID:             zone.ID,
			Name:           zone.Name,
			Type:           zone.Type,
			TotalCapacity:  zone.TotalCapacity,
			AvailableSpots: zone.TotalCapacity - int(activeCount),
			PricePerHour:   zone.PricePerHour,
			CreatedAt:      zone.CreatedAt,
		})
	}

	return responses, nil
}

// get zone by id
func (s *ZoneService) GetZoneByID(id uint) (*dto.ZoneResponse, error) {
	zone, err := s.zoneRepo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, utils.ErrNotFound
		}
		return nil, err
	}

	activeCount, err := s.zoneRepo.CountActiveReservations(zone.ID)
	if err != nil {
		return nil, err
	}

	response := &dto.ZoneResponse{
		ID:             zone.ID,
		Name:           zone.Name,
		Type:           zone.Type,
		TotalCapacity:  zone.TotalCapacity,
		AvailableSpots: zone.TotalCapacity - int(activeCount),
		PricePerHour:   zone.PricePerHour,
		CreatedAt:      zone.CreatedAt,
		UpdatedAt:      zone.UpdatedAt,
	}

	return response, nil
}

// update zone
func (s *ZoneService) UpdateZone(id uint, req dto.UpdateZoneRequest) (*dto.ZoneResponse, error) {
	zone, err := s.zoneRepo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, utils.ErrNotFound
		}
		return nil, err
	}

	// partial updates
	if req.Name != nil {
		zone.Name = *req.Name
	}
	if req.Type != nil {
		zone.Type = *req.Type
	}
	if req.TotalCapacity != nil {
		zone.TotalCapacity = *req.TotalCapacity
	}
	if req.PricePerHour != nil {
		zone.PricePerHour = *req.PricePerHour
	}

	if err := s.zoneRepo.Update(zone); err != nil {
		return nil, err
	}

	activeCount, err := s.zoneRepo.CountActiveReservations(zone.ID)
	if err != nil {
		return nil, err
	}

	response := &dto.ZoneResponse{
		ID:             zone.ID,
		Name:           zone.Name,
		Type:           zone.Type,
		TotalCapacity:  zone.TotalCapacity,
		AvailableSpots: zone.TotalCapacity - int(activeCount),
		PricePerHour:   zone.PricePerHour,
		CreatedAt:      zone.CreatedAt,
		UpdatedAt:      zone.UpdatedAt,
	}

	return response, nil
}

// delete zone
func (s *ZoneService) DeleteZone(id uint) error {
	_, err := s.zoneRepo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return utils.ErrNotFound
		}
		return err
	}

	return s.zoneRepo.Delete(id)
}
