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

// Create an account
func (repo *Accounts) Create(account *models.Account) error {
	return repo.conn.Create(account).Error
}

// Size of this repository
func (repo *Accounts) Size() (int, error) {
	var c int
	err := repo.conn.Model(repo.Model()).Count(&c).Error

	return c, err
}

// DeleteAll records in this repository
func (repo *Accounts) DeleteAll() error {
	return repo.conn.Delete(repo.Model()).Error
}
