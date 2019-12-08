package helpers

import (
	"html/template"

	"github.com/krakenlab/plataform/internal/configs"
	"github.com/qor/i18n"
	"github.com/qor/i18n/backends/yaml"
)

// Translate a key
func Translate() func(string, string, ...interface{}) template.HTML {
	return i18n.New(yaml.New(configs.Path().Locales())).Fallbacks("en-US").T
}
