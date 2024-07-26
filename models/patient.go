package models

import (
	"time"

	"gorm.io/gorm"
)

type Patient struct {
	ID        string         `json:"id" gorm:"primary_key"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	Name      string         `json:"name"`
	ContactNo string         `json:"contact_no"`
	Address   string         `json:"address"`
	DoctorID  string         `json:"doctor_id"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
