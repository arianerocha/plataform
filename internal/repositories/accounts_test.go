package repositories

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type AccountsTestSuite struct {
	suite.Suite
}

func (suite *AccountsTestSuite) TestNewAccounts() {
	suite.NotNil(NewAccounts())
}

func (suite *AccountsTestSuite) TestAutoMigrate() {
	suite.NoError(NewAccounts().AutoMigrate())
}

func (suite *AccountsTestSuite) TestCreate() {
	suite.NoError(NewAccounts().DeleteAll())

	size, err := NewAccounts().Size()
	suite.NoError(err)
	suite.Equal(0, size)

	account := NewAccounts().Model()
	account.Email = "testing@testing.io"
	account.SecurePassword = "FFFFFF"
	account.Locale = "pt-BR"

	suite.NoError(NewAccounts().Create(account))

	size, err = NewAccounts().Size()
	suite.NoError(err)
	suite.Equal(1, size)
}

func TestAccountsTestSuite(t *testing.T) {
	suite.Run(t, new(AccountsTestSuite))
}
