package request

import (
	"time"
)

type UpdatePersonRequest struct {
	Name     string  `json:"name"  binding:"required"`
	DateBirth     time.Time  `json:"datebirth,omitempty"  binding:"required"`
	Contact     string  `json:"contact"  binding:"required"`
}