package services

import (
	"github.com/krakenlab/plataform/internal/models"
	"github.com/krakenlab/plataform/internal/repositories"
)

// Contact service
type Contact struct {
	contact *models.Contact
	repo    *repositories.Contacts

	errors []error
}

// NewContact constructor
func NewContact(email, locale string) *Contact {
	repo := repositories.NewContacts()
	service := &Contact{
		repo:    repo,
		contact: repo.Model(),
		errors:  []error{},
	}

	service.contact.Email = email
	service.contact.Locale = repositories.NewLocales().LocaleOrDefault(locale).Symbol

	return service
}

// Create contact
func (service *Contact) Create() {
	_ = service.repo.CreateIfNotExists(service.contact)
}
