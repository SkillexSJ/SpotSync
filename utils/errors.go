package utils

import "errors"

// Centralized application level errors
var (
	ErrNotFound           = errors.New("resource not found")
	ErrDuplicateEmail     = errors.New("email already registered")
	ErrInvalidCredentials = errors.New("invalid email or password")
	ErrUnauthorized       = errors.New("unauthorized access")
	ErrForbidden          = errors.New("you do not have permission to perform this action")
	ErrZoneFull           = errors.New("parking zone is at full capacity")
)
