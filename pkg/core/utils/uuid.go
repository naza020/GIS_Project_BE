package utils

import (
	guuid "github.com/google/uuid"
)

// NewID Gen
func NewSID() string {
	return guuid.New().String()
}
