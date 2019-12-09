package handlers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/krakenlab/plataform/internal/repositories"
	"github.com/stretchr/testify/suite"
)

type LocalizedSignUpTestSuite struct {
	handlerSuite
}

func (suite *LocalizedSignUpTestSuite) TestNewNewSignUp() {
	suite.NotNil(NewLocalizedSignUp())
}

func (suite *LocalizedSignUpTestSuite) TestGET() {
	NewLocalizedSignUp().Setup(suite.engine)
	for _, locale := range repositories.NewLocales().All() {
		req, err := http.NewRequest(http.MethodGet, strings.ReplaceAll(LocalizedSignUpPath, ":locale", locale.Symbol), nil)
		suite.NoError(err)

		rr := httptest.NewRecorder()

		suite.ServeHTTP(rr, req)
		suite.Equal(http.StatusOK, rr.Code)
	}
}

func TestLocalizedSignUpTestSuite(t *testing.T) {
	suite.Run(t, new(LocalizedSignUpTestSuite))
}
