package services

import (
	"errors"
	"testing"

	"github.com/krakenlab/plataform/internal/repositories"
	"github.com/stretchr/testify/suite"
)

type SignUpTestSuite struct {
	suite.Suite
}

func (suite *SignUpTestSuite) TestNewSignUp() {
	err := repositories.NewAccounts().DeleteAll()
	suite.NoError(err)

	service := NewSignUp("email@mail.to", "pwd", "pwd", "pt-BR", true, true, true)
	suite.NotNil(service)

	suite.NotNil(service.repo)
	suite.NotNil(service.account)

	suite.Equal("email@mail.to", service.form.email)
	suite.Equal("pwd", service.form.password)
	suite.Equal("pwd", service.form.repeatPassword)
	suite.Equal("pt-BR", service.locale)
	suite.True(service.form.terms)
	suite.True(service.form.privacy)
	suite.True(service.form.newsletter)
}

func (suite *SignUpTestSuite) TestValid() {
	err := repositories.NewAccounts().DeleteAll()
	suite.NoError(err)

	service := NewSignUp("@fail@email@-unk-mail.to", "pwd", "pwd err", "pt-BR", false, false, false)
	suite.False(service.Valid())

	suite.Contains(service.Errors(), errors.New(PrivacySignUpError))
	suite.Contains(service.Errors(), errors.New(TermsSignUpError))
	suite.Contains(service.Errors(), errors.New(ValidPasswordSignUpError))
	suite.Contains(service.Errors(), errors.New(EmailFormatSignUpError))

	service = NewSignUp("marlon.schweigert@krakenlab.io", "nopasswd123123", "nopasswd123123", "pt-BR", true, true, true)
	suite.True(service.Valid())

	suite.Zero(len(service.Errors()))
}

func (suite *SignUpTestSuite) TestCreate() {
	err := repositories.NewAccounts().DeleteAll()
	suite.NoError(err)

	service := NewSignUp("marlon.schweigert@krakenlab.io", "nopasswd123123", "nopasswd123123", "pt-BR", true, true, true)
	suite.True(service.Valid())

	suite.Zero(len(service.Errors()))
	suite.NoError(service.Create())

	service = NewSignUp("marlon.schweigert@krakenlab.io", "nopasswd123123", "nopasswd123123", "pt-BR", true, true, true)
	suite.False(service.Valid())

	suite.Equal(1, len(service.Errors()))
	suite.Contains(service.Errors(), errors.New(EmailMustBeUniqSignUpError))
}

func TestSignUpTestSuite(t *testing.T) {
	suite.Run(t, new(SignUpTestSuite))
}
