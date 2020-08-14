package models

import (
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type Super struct {
	ID              uuid.UUID `gorm:"type:uuid;primary_key;" json:"id"`
	Name            string    `gorm:"size:100;not null" json:"name"`
	FullName        string    `gorm:"not null" json:"full_name"`
	Intelligence    int       `gorm:"not null" json:"intelligence"`
	Power           int       `gorm:"not null" json:"power"`
	Occupation      string    `gorm:"not null" json:"occupation"`
	Image           string    `gorm:"not null" json:"image"`
	Alignment       string    `gorm:"not null" json:"alignment"`
	Groups          string    `json:"groups"`
	NumberOfParents int       `json:"number_of_parents"`
	CreatedAt       time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt       time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

// BeforeCreate will set a UUID rather than numeric ID.
func (s *Super) BeforeCreate(db *gorm.Scope) error {
	return db.SetColumn("ID", uuid.NewV4().String())
}

func (s *Super) CreateSuper(db *gorm.DB) (*Super, error) {
	err := db.Debug().Save(&s).Error
	if err != nil {
		return &Super{}, err
	}
	return s, nil
}
