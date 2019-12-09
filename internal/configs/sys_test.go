package configs

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type SysTestSuite struct {
	suite.Suite
}

func (suite *SysTestSuite) TestSys() {
	suite.NotNil(Sys())
}

func (suite *SysTestSuite) TestEnv() {
	suite.Equal("testing", Sys().Env("NO ENV", "testing"))
	suite.NotEqual("testing", Sys().Env("GOPATH", "testing"))
}

func (suite *SysTestSuite) TestSysHomeDir() {
	suite.NotEqual(Sys().HomeDir(), "/")
}

func TestSysTestSuite(t *testing.T) {
	suite.Run(t, new(SysTestSuite))
}
