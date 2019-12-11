package services

import (
	"errors"

	"github.com/krakenlab/plataform/internal/models"
	"github.com/krakenlab/plataform/internal/repositories"

	"github.com/badoux/checkmail"
)

// SignUpErrors
const (
	PrivacySignUpError         = "signup.privacy_error"
	TermsSignUpError           = "signup.terms_error"
	ValidPasswordSignUpError   = "signup.valid_password_error"
	RepeatPasswordSignUpError  = "signup.repeat_password_error"
	EmailFormatSignUpError     = "signup.email_format_error"
	EmailMustBeUniqSignUpError = "signup.uniq_email_error"
)

// SignUp service
type SignUp struct {
	account *models.Account
	repo    *repositories.Accounts

	locale string

	form struct {
		email          string
		password       string
		repeatPassword string
		terms          bool
		privacy        bool
		newsletter     bool
	}

	errors []error
}

// NewSignUp constructor
func NewSignUp(email, password, repeatPassword, locale string, terms, privacy, newsletter bool) *SignUp {
	repo := repositories.NewAccounts()
	service := &SignUp{
		repo:    repo,
		account: repo.Model(),
		errors:  []error{},
		locale:  repositories.NewLocales().LocaleOrDefault(locale).Symbol,
	}

	service.form.email = email
	service.form.password = password
	service.form.repeatPassword = repeatPassword
	service.form.terms = terms
	service.form.privacy = privacy
	service.form.newsletter = newsletter

	return service
}

// Errors getter
func (service *SignUp) Errors() []error {
	return service.errors
}

// Valid the form
func (service *SignUp) Valid() bool {
	service.mustAcceptPrivacy()
	service.mustAcceptTerms()
	service.passwordsMustBeValid()
	service.passwordsMustMatch()
	service.emailMustBeValid()

	return len(service.errors) == 0
}

// Create the account
func (service *SignUp) Create() error {
	service.account.Locale = service.locale
	service.account.Email = service.form.email
	service.account.SecurePassword = NewPassword().SecurePassword(service.form.password)

	if service.form.newsletter {
		go NewContact(service.account.Email, service.account.Locale).Create()
	}

	return service.repo.Create(service.account)
}

func (service *SignUp) mustAcceptPrivacy() {
	if service.form.privacy {
		return
	}

	service.errors = append(service.errors, errors.New(PrivacySignUpError))
}

func (service *SignUp) mustAcceptTerms() {
	if service.form.terms {
		return
	}

	service.errors = append(service.errors, errors.New(TermsSignUpError))
}

func (service *SignUp) passwordsMustMatch() {
	if service.form.password == service.form.repeatPassword {
		return
	}

	service.errors = append(service.errors, errors.New(RepeatPasswordSignUpError))
}

func (service *SignUp) passwordsMustBeValid() {
	if len(service.form.password) >= 8 && len(service.form.password) <= 256 {
		return
	}

	service.errors = append(service.errors, errors.New(ValidPasswordSignUpError))
}

func (service *SignUp) emailMustBeValid() {
	if checkmail.ValidateFormat(service.form.email) != nil {
		service.errors = append(service.errors, errors.New(EmailFormatSignUpError))
	}

	registred, err := service.repo.EmailAlreadyRegistered(service.form.email)

	if err != nil || registred {
		service.errors = append(service.errors, errors.New(EmailMustBeUniqSignUpError))
	}
}
