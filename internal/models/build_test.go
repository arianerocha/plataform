package models

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type BuildTestSuite struct {
	suite.Suite
}

func (suite *BuildTestSuite) TestBuild() {
	suite.T().Skip()
}

func TestBuildTestSuite(t *testing.T) {}
