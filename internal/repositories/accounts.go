package repositories

import (
	"github.com/jinzhu/gorm"
	"github.com/krakenlab/plataform/internal/models"
	"github.com/krakenlab/plataform/internal/repositories/db"
)

// Accounts repository manage data from an account in the database
type Accounts struct {
	conn *gorm.DB
}

// NewAccounts start a new repository instance
func NewAccounts() *Accounts {
	return &Accounts{conn: db.Connection()}
}

// Model instance
func (repo *Accounts) Model() *models.Account {
	return &models.Account{}
}

// AutoMigrate this repository
func (repo *Accounts) AutoMigrate() error {
	return repo.conn.AutoMigrate(repo.Model()).Error
}
