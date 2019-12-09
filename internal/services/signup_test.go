package services

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type SignUpTestSuite struct {
	suite.Suite
}

func (suite *SignUpTestSuite) TestNewSignUp() {
	service := NewSignUp("email@mail.to", "pwd", "pwd", true, true, true)
	suite.NotNil(service)

	suite.NotNil(service.repo)
	suite.NotNil(service.account)

	suite.Equal("email@mail.to", service.form.email)
	suite.Equal("pwd", service.form.password)
	suite.Equal("pwd", service.form.repeatPassword)
	suite.True(service.form.agree)
	suite.True(service.form.privacy)
	suite.True(service.form.newsletter)
}

func TestSignUpTestSuite(t *testing.T) {
	suite.Run(t, new(SignUpTestSuite))
}
