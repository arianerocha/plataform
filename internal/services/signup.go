package services

import (
	"github.com/krakenlab/plataform/internal/models"
	"github.com/krakenlab/plataform/internal/repositories"

	_ "github.com/badoux/checkmail"
)

// SignUp service
type SignUp struct {
	account *models.Account
	repo    *repositories.Accounts

	form struct {
		email          string
		password       string
		repeatPassword string
		agree          bool
		privacy        bool
		newsletter     bool
	}
}

// NewSignUp constructor
func NewSignUp(email, password, repeatPassword string, agree, privacy, newsletter bool) *SignUp {
	repo := repositories.NewAccounts()
	service := &SignUp{
		repo:    repo,
		account: repo.Model(),
	}

	service.form.email = email
	service.form.password = password
	service.form.repeatPassword = repeatPassword
	service.form.agree = agree
	service.form.privacy = privacy
	service.form.newsletter = newsletter

	return service
}

// Valid the form
func (service *SignUp) Valid() {}
