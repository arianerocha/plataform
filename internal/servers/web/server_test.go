package web

import (
	"os"
	"testing"

	"github.com/krakenlab/plataform/internal/configs"
	"github.com/stretchr/testify/suite"
)

type ServerTestSuite struct {
	suite.Suite
}

func (suite *ServerTestSuite) TestServer() {
	suite.NotNil(NewServer())
	suite.NotNil(NewServer().engine)
}

func (suite *ServerTestSuite) TestSetup() {
	suite.NotPanics(NewServer().Setup)
}

func (suite *ServerTestSuite) TestSetupStatic() {
	server := NewServer()
	suite.NotPanics(server.SetupStatic)

	suite.Equal(2, len(server.engine.Routes()))
}

func (suite *ServerTestSuite) TestSetupHTMLRender() {
	server := NewServer()
	suite.NotPanics(server.SetupHTMLRender)
}

func (suite *ServerTestSuite) TestRun() {
	defer os.Setenv("PORT", configs.Web().Port())

	err := os.Setenv("PORT", "54321")
	suite.NoError(err)

	server := NewServer()

	suite.NotPanics(func() { go server.Run() })
}

func TestServerTestSuite(t *testing.T) {
	suite.Run(t, new(ServerTestSuite))
}
