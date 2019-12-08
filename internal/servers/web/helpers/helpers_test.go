package helpers

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type HelpersTestSuite struct {
	suite.Suite
}

func (suite *HelpersTestSuite) TestHelpers() {
	suite.NotNil(Helpers())

	suite.Contains(Helpers(), "translate")
}

func TestHelpersTestSuite(t *testing.T) {
	suite.Run(t, new(HelpersTestSuite))
}
