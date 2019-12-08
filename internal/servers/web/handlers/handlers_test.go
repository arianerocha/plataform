package handlers

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type HandlersTestSuite struct {
	handlerSuite
}

func (suite *HandlersTestSuite) TestHandlers() {
	suite.NotNil(Handlers())
}

func TestHandlersTestSuite(t *testing.T) {
	suite.Run(t, new(HandlersTestSuite))
}
