package services

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type ContactTestSuite struct {
	suite.Suite
}

func (suite *ContactTestSuite) TestNewContact() {
	service := NewContact("email@mail.to", "pt-BR")
	suite.NotNil(service)

	suite.NotNil(service.repo)
	suite.NotNil(service.contact)

	suite.Equal("email@mail.to", service.contact.Email)
	suite.Equal("pt-BR", service.contact.Locale)
}

func (suite *ContactTestSuite) TestCreate() {
	service := NewContact("email@mail.to", "pt-BR")
	suite.NotNil(service)

	suite.NotPanics(service.Create)
}

func TestContactTestSuite(t *testing.T) {
	suite.Run(t, new(ContactTestSuite))
}
