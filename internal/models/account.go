package models

import "github.com/jinzhu/gorm"

// Account structures the basic authentication data and basic information for access to the platform and its games.
type Account struct {
	gorm.Model

	Email          string `json:"email"`
	SecurePassword string `json:"-"`
}
