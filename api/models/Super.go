package models

import (
	"errors"
	"time"

	"github.com/Lgdev07/superapi/api/utils"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

// Super is the struct for the table supers
type Super struct {
	ID              uuid.UUID `gorm:"type:uuid;primary_key;" json:"id"`
	Name            string    `gorm:"size:100;not null" json:"name"`
	FullName        string    `gorm:"not null" json:"full_name"`
	Intelligence    int       `gorm:"not null" json:"intelligence"`
	Power           int       `gorm:"not null" json:"power"`
	Occupation      string    `gorm:"not null" json:"occupation"`
	Image           string    `gorm:"not null" json:"image"`
	Alignment       string    `gorm:"not null" json:"-"`
	Groups          string    `json:"groups"`
	NumberOfParents int       `json:"number_of_parents"`
	CreatedAt       time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"-"`
	UpdatedAt       time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"-"`
}

// BeforeCreate will set a UUID rather than numeric ID.
func (s *Super) BeforeCreate(db *gorm.Scope) error {
	return db.SetColumn("ID", uuid.NewV4().String())
}

// Save creates a record in database with the struct super
func (s *Super) Save(db *gorm.DB) (*Super, error) {
	err := db.Save(&s).Error
	if err != nil {
		return &Super{}, err
	}
	return s, nil
}

// DeleteSuperByID deletes the super record in database by the given ID
func DeleteSuperByID(db *gorm.DB, id string) error {
	err := db.Model(&Super{}).Where("id = ?", id).Delete(&Super{}).Error
	if err != nil {
		return err
	}
	return nil
}

// GetSuperByID returns the super with the given ID
func GetSuperByID(db *gorm.DB, id string) (*Super, error) {
	super := Super{}

	err := db.Model(&Super{}).Where("id = ?", id).Take(&super).Error
	if err != nil {
		return &Super{}, err
	}

	return &super, nil
}

// FindSupers returns the supers from database with the given params
func FindSupers(db *gorm.DB, params map[string]string) (*[]Super, error) {
	supers := []Super{}

	chain := db.Model(&Super{})

	if params["name"] != "" {
		chain = chain.Where("name = ?", params["name"])
	}

	if params["alignment"] != "" {
		if utils.Contains([]string{"good", "bad"}, params["alignment"]) {
			chain = chain.Where("alignment = ?", params["alignment"])
		} else {
			return &[]Super{}, errors.New("alignment must be good or bad")
		}
	}

	if err := chain.Find(&supers).Error; err != nil {
		return &[]Super{}, err
	}
	return &supers, nil

}
