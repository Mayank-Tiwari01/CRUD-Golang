package models

import (
	"time"

	"gorm.io/gorm"
)

type Doctor struct {
	ID        string         `json:"id" gorm:"primary_key"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	Name      string         `json:"name"`
	ContactNo string         `json:"contact_no"`
	Address   string         `json:"address"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
