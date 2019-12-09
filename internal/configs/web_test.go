package configs

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type WebTestSuite struct {
	suite.Suite
}

func (suite *WebTestSuite) TestSys() {
	suite.NotNil(Sys())
}

func (suite *WebTestSuite) TestPort() {
	suite.Equal("3030", Web().Port())
}

func (suite *WebTestSuite) TestRunnableInterface() {
	suite.Equal("0.0.0.0:3030", Web().RunnableInterface())
}

func TestWebTestSuite(t *testing.T) {
	suite.Run(t, new(WebTestSuite))
}
