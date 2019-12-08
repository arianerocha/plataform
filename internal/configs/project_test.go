package configs

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type ProjectTestSuite struct {
	suite.Suite
}

func (suite *ProjectTestSuite) TestProject() {
	suite.NotNil(Project())
}

func (suite *ProjectTestSuite) TestProjectName() {
	suite.Equal("plataform", Project().Name())
}

func (suite *ProjectTestSuite) TestProjectOrg() {
	suite.Equal("krakenlab", Project().Org())
}

func TestProjectTestSuite(t *testing.T) {
	suite.Run(t, new(ProjectTestSuite))
}
