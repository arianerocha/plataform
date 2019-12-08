package htmlrender

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type HTMLRenderTestSuite struct {
	suite.Suite
}

func (suite *HTMLRenderTestSuite) TestNew() {
	suite.NotNil(New())
}

func TestHTMLRenderTestSuite(t *testing.T) {
	suite.Run(t, new(HTMLRenderTestSuite))
}
