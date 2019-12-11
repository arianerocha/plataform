package models

import "github.com/jinzhu/gorm"

// Contact store an email with your contact location
type Contact struct {
	gorm.Model

	Email  string `json:"email" gorm:"index"`
	Locale string `json:"locale" gorm:"size:5"`
}
