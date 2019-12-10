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

func (suite *LocalizedSignUpTestSuite) TestPOST() {
	NewLocalizedSignUp().Setup(suite.engine)
	// Success Case
	for _, locale := range repositories.NewLocales().All() {
		err := repositories.NewAccounts().DeleteAll()
		suite.NoError(err)

		count, err := repositories.NewAccounts().Size()
		suite.NoError(err)
		suite.Zero(count)

		form := "email=testing@testing.onion&password=123123123&repeat_password=123123123&terms=on&privacy=on&newslatter=on"

		reader := strings.NewReader(form)
		req, err := http.NewRequest(http.MethodPost, strings.ReplaceAll(LocalizedSignUpPath, ":locale", locale.Symbol), reader)
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		suite.NoError(err)

		rr := httptest.NewRecorder()

		suite.ServeHTTP(rr, req)

		count, err = repositories.NewAccounts().Size()
		suite.NoError(err)
		suite.Equal(1, count)

		suite.Equal(http.StatusCreated, rr.Code)
	}

	// Fail Case
	for _, locale := range repositories.NewLocales().All() {
		err := repositories.NewAccounts().DeleteAll()
		suite.NoError(err)

		count, err := repositories.NewAccounts().Size()
		suite.NoError(err)
		suite.Zero(count)

		form := "email=testing@testing.onion&password=123123123&repeat_password=123123123&terms=on&newslatter=on"

		reader := strings.NewReader(form)
		req, err := http.NewRequest(http.MethodPost, strings.ReplaceAll(LocalizedSignUpPath, ":locale", locale.Symbol), reader)
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		suite.NoError(err)

		rr := httptest.NewRecorder()

		suite.ServeHTTP(rr, req)

		count, err = repositories.NewAccounts().Size()
		suite.NoError(err)
		suite.Zero(count)

		suite.Equal(http.StatusOK, rr.Code)
	}
}

func TestLocalizedSignUpTestSuite(t *testing.T) {
	suite.Run(t, new(LocalizedSignUpTestSuite))
}
