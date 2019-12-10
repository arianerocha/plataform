package helpers

import (
	"fmt"
	"html/template"
)

// TranslateError a key
func TranslateError() func(string, error) template.HTML {
	return func(locale string, err error) template.HTML {
		return Translate()(locale, fmt.Sprintf("errors.%s", err.Error()))
	}
}
