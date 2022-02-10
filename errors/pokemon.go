package errors

import (
	errs "errors"
)

type ServiceError error

var (
	ErrNotFound             ServiceError = errs.New("pokemon not found")
	ErrEmptyData            ServiceError = errs.New("data is empty")
	ErrDataNotInitialized   ServiceError = errs.New("data not initialized")
	ErrPokemonAlreadyExists ServiceError = errs.New("pokemon already exists")
)
