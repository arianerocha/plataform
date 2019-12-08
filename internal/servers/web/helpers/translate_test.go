package helpers

import (
	"html/template"
	"testing"

	"github.com/stretchr/testify/suite"
)

type TranslateTestSuite struct {
	suite.Suite
}

func (suite *TranslateTestSuite) TestTranslate() {
	suite.Equal(template.HTML("Inicio"), Translate()("pt-BR", "layouts.navbar.home"))
	suite.Equal(template.HTML("Home"), Translate()("en-US", "layouts.navbar.home"))
}

func TestTranslateTestSuite(t *testing.T) {
	suite.Run(t, new(TranslateTestSuite))
}
