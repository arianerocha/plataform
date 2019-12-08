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

func TestPathTestSuite(t *testing.T) {
	suite.Run(t, new(PathTestSuite))
}
