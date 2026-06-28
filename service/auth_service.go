package service

import (
	"errors"

	"gorm.io/gorm"

	"github.com/SkillexSJ/SpotSync/dto"
	"github.com/SkillexSJ/SpotSync/models"
	"github.com/SkillexSJ/SpotSync/repository"
	"github.com/SkillexSJ/SpotSync/utils"
)

type AuthService struct {
	userRepo  *repository.UserRepository
	jwtSecret string
}

// create service
func NewAuthService(userRepo *repository.UserRepository, jwtSecret string) *AuthService {
	return &AuthService{
		userRepo:  userRepo,
		jwtSecret: jwtSecret,
	}
}

// register
func (s *AuthService) Register(req dto.RegisterRequest) (*dto.UserResponse, error) {
	// Check email
	existingUser, err := s.userRepo.FindByEmail(req.Email)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	if existingUser != nil {
		return nil, utils.ErrDuplicateEmail
	}

	// Hash the password
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	// Default role
	role := req.Role
	if role == "" {
		role = "driver"
	}

	// Create user model
	user := models.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: hashedPassword,
		Role:     role,
	}

	if err := s.userRepo.Create(&user); err != nil {
		return nil, err
	}

	// response dto
	response := &dto.UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Role:      user.Role,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	return response, nil
}

// login
func (s *AuthService) Login(req dto.LoginRequest) (*dto.LoginResponse, error) {
	// Find user
	user, err := s.userRepo.FindByEmail(req.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, utils.ErrInvalidCredentials
		}
		return nil, err
	}

	// Compare password
	if err := utils.CheckPassword(user.Password, req.Password); err != nil {
		return nil, utils.ErrInvalidCredentials
	}

	// generate jwt token
	token, err := utils.GenerateToken(user.ID, user.Role, s.jwtSecret)
	if err != nil {
		return nil, err
	}

	// response
	response := &dto.LoginResponse{
		Token: token,
		User: dto.UserResponse{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
			Role:  user.Role,
		},
	}

	return response, nil
}
