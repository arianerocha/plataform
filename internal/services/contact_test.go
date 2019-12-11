package services

import (
	"testing"

	"github.com/krakenlab/plataform/internal/repositories"
	"github.com/stretchr/testify/suite"
)

type ContactTestSuite struct {
	suite.Suite
}

func (suite *ContactTestSuite) TestNewContact() {
	err := repositories.NewAccounts().DeleteAll()
	suite.NoError(err)

	service := NewContact("email@mail.to", "pt-BR")
	suite.NotNil(service)

	suite.NotNil(service.repo)
	suite.NotNil(service.contact)

	suite.Equal("email@mail.to", service.contact.Email)
	suite.Equal("pt-BR", service.contact.Locale)
}

func (suite *ContactTestSuite) TestCreate() {
	err := repositories.NewAccounts().DeleteAll()
	suite.NoError(err)

	service := NewContact("email@mail.to", "pt-BR")
	suite.NotNil(service)

	suite.NotPanics(service.Create)
}

func TestContactTestSuite(t *testing.T) {
	suite.Run(t, new(ContactTestSuite))
}
