package services

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type PasswordTestSuite struct {
	suite.Suite
}

func (suite *PasswordTestSuite) TestNewPassword() {
	service := NewPassword()
	suite.NotNil(service)
}

func (suite *PasswordTestSuite) TestSecurePassword() {
	service := NewPassword()
	suite.NotNil(service)

	suite.Equal(
		60,
		len(service.SecurePassword("testing")),
	)
}

func (suite *PasswordTestSuite) TestCheckPasswordHash() {
	service := NewPassword()
	suite.NotNil(service)

	password := "testing"
	securePassword := service.SecurePassword(password)
	suite.True(service.CheckPasswordHash(password, securePassword))
}

func TestPasswordTestSuite(t *testing.T) {
	suite.Run(t, new(PasswordTestSuite))
}
