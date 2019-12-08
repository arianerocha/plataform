package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/suite"
)

type RootTestSuite struct {
	handlerSuite
}

func (suite *RootTestSuite) TestNewRoot() {
	suite.NotNil(NewRoot())
}

func (suite *RootTestSuite) TestGET() {
	NewRoot().Setup(suite.engine)

	req, err := http.NewRequest(http.MethodGet, RootPath, nil)
	suite.NoError(err)

	rr := httptest.NewRecorder()

	suite.ServeHTTP(rr, req)
	suite.Equal(http.StatusTemporaryRedirect, rr.Code)
}

func TestRootTestSuite(t *testing.T) {
	suite.Run(t, new(RootTestSuite))
}
