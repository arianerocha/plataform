package configs

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type PathTestSuite struct {
	suite.Suite
}

func (suite *PathTestSuite) TestPath() {
	suite.NotNil(Path())
}

func (suite *PathTestSuite) TestPathApp() {
	suite.Contains(Path().App(), Sys().HomeDir())
	suite.Contains(Path().App(), Project().Org())
	suite.Contains(Path().App(), Project().Name())
}

func (suite *PathTestSuite) TestPathViews() {
	suite.Contains(Path().Views(), Path().App())
	suite.Contains(Path().Views(), "/views")
}

func (suite *PathTestSuite) TestPathStatic() {
	suite.Contains(Path().Static(), Path().App())
	suite.Contains(Path().Static(), "/static")
}

func TestPathTestSuite(t *testing.T) {
	suite.Run(t, new(PathTestSuite))
}
