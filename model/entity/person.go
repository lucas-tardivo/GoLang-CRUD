package entity

import(
	"time"
	"github.com/google/uuid"
)

type Person struct {
	ID     uuid.UUID  `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id,omitempty"`
	Name     string  `gorm:"not null" json:"name,omitempty"`
	DateBirth     time.Time  `gorm:"not null" json:"datebirth,omitempty"`
	Contact     string  `gorm:"nullable" json:"contact,omitempty"`
}