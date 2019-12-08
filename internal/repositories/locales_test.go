package repositories

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type LocalesTestSuite struct {
	suite.Suite
}

func (suite *LocalesTestSuite) TestNewLocales() {
	suite.NotNil(NewLocales())
}

func (suite *LocalesTestSuite) TestAll() {
	suite.Contains(NewLocales().All(), NewLocales().PtBR())
	suite.Contains(NewLocales().All(), NewLocales().EnUS())
	suite.Contains(NewLocales().All(), NewLocales().EnGB())

	suite.Contains(NewLocales().All(), NewLocales().Default())
}

func (suite *LocalesTestSuite) TestLocaleOrDefault() {
	suite.Equal(NewLocales().PtBR(), NewLocales().LocaleOrDefault("pt-BR"))
	suite.Equal(NewLocales().Default(), NewLocales().LocaleOrDefault("zz-ZZ"))
}

func (suite *LocalesTestSuite) TestPtBR() {
	suite.Equal("pt-BR", NewLocales().PtBR().Symbol)
}

func (suite *LocalesTestSuite) TestEnUS() {
	suite.Equal("en-US", NewLocales().EnUS().Symbol)
}

func (suite *LocalesTestSuite) TestEnGB() {
	suite.Equal("en-GB", NewLocales().EnGB().Symbol)
}

func (suite *LocalesTestSuite) TestDefault() {
	suite.Equal("en-US", NewLocales().Default().Symbol)
}

func TestLocalesTestSuite(t *testing.T) {
	suite.Run(t, new(LocalesTestSuite))
}
