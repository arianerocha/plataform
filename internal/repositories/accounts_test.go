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

func TestAccountsTestSuite(t *testing.T) {
	suite.Run(t, new(AccountsTestSuite))
}
