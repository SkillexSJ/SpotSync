package repository

import "errors"

// Custom repository-level errors
var (
	ErrZoneFull = errors.New("parking zone is at full capacity")
)
