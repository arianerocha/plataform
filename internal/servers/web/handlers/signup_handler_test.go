package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/suite"
)

type SignUpTestSuite struct {
	handlerSuite
}

func (suite *SignUpTestSuite) TestNewNewSignUp() {
	suite.NotNil(NewSignUp())
}

func (suite *SignUpTestSuite) TestGET() {
	NewSignUp().Setup(suite.engine)

	req, err := http.NewRequest(http.MethodGet, SignUpPath, nil)
	suite.NoError(err)

	rr := httptest.NewRecorder()

	suite.ServeHTTP(rr, req)
	suite.Equal(http.StatusPermanentRedirect, rr.Code)
	suite.Equal("<a href=\"/signup/en-US\">Permanent Redirect</a>.\n\n", rr.Body.String())
}

func TestSignUpTestSuite(t *testing.T) {
	suite.Run(t, new(SignUpTestSuite))
}
