package handlers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/krakenlab/plataform/internal/repositories"
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
	for _, locale := range repositories.NewLocales().All() {
		req, err := http.NewRequest(http.MethodGet, strings.ReplaceAll(SignUpPath, ":locale", locale.Symbol), nil)
		suite.NoError(err)

		rr := httptest.NewRecorder()

		suite.ServeHTTP(rr, req)
		suite.Equal(http.StatusOK, rr.Code)
	}
}

func TestSignUpTestSuite(t *testing.T) {
	suite.Run(t, new(SignUpTestSuite))
}
