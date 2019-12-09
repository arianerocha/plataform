package db

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type LocalesTestSuite struct {
	suite.Suite
}

func (suite *LocalesTestSuite) TestConnection() {
	suite.NotNil(Connection())
}

func TestLocalesTestSuite(t *testing.T) {
	suite.Run(t, new(LocalesTestSuite))
}
