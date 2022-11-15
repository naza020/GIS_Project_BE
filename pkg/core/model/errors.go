package model

import (
	"github.com/pkg/errors"
)

var (
	ErrValidationFailed          = errors.New("Validation Failed")
	ErrNotFound                  = errors.New("Record Not Found")
	ErrUnauthorized              = errors.New("Unauthorized")
	ErrCallExternalServiceFailed = errors.New("Validation Failed")

	ErrUpdateFailed = errors.New("DB Update Error")
	ErrInsertFailed = errors.New("DB Insert Error")
	ErrDeleteFailed = errors.New("DB Delete Error")
	ErrQueryFailed  = errors.New("DB Query Error")
)
