package repositories

import (
	"github.com/jinzhu/gorm"
	"github.com/krakenlab/plataform/internal/models"
	"github.com/krakenlab/plataform/internal/repositories/db"
)

// Contacts repository manage data from an Contact in the database
type Contacts struct {
	conn *gorm.DB
}

// NewContacts start a new repository instance
func NewContacts() *Contacts {
	return &Contacts{conn: db.Connection()}
}

// Model instance
func (repo *Contacts) Model() *models.Contact {
	return &models.Contact{}
}

// AutoMigrate this repository
func (repo *Contacts) AutoMigrate() error {
	return repo.conn.AutoMigrate(repo.Model()).Error
}

// Create an Contact
func (repo *Contacts) Create(contact *models.Contact) error {
	return repo.conn.FirstOrCreate(contact, "email = ?", contact.Email).Error
}
