package configs

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type DbTestSuite struct {
	suite.Suite
}

func (suite *DbTestSuite) TestDb() {
	suite.NotNil(Db())
}

func (suite *DbTestSuite) TestAdapter() {
	suite.NotEqual("", Db().Adapter())
}

func (suite *DbTestSuite) TestURL() {
	suite.NotEqual("", Db().URL())
}

func TestDbTestSuite(t *testing.T) {
	suite.Run(t, new(DbTestSuite))
}
