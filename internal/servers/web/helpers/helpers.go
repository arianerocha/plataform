package helpers

import "html/template"

// Helpers that are injected in HTMLRender
func Helpers() template.FuncMap {
	return template.FuncMap{
		"translate":       Translate(),
		"translate_error": TranslateError(),
	}
}
