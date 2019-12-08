package repositories

import "github.com/krakenlab/plataform/internal/models"

// Locales repository manage data from an account in memory
type Locales struct{}

// NewLocales start a new repository instance
func NewLocales() *Locales {
	return &Locales{}
}

// All locales used in this application
func (rep *Locales) All() []*models.Locale {
	return []*models.Locale{
		rep.PtBR(),
		rep.EnUS(),
		rep.EnGB(),
	}
}

// LocaleOrDefault try to fetch a locale. If not, returns the default value.
func (rep *Locales) LocaleOrDefault(symbol string) *models.Locale {
	for _, locale := range rep.All() {
		if locale.Symbol == symbol {
			return locale
		}
	}

	return rep.Default()
}

// Default locale is en-US
func (rep *Locales) Default() *models.Locale {
	return &models.Locale{Symbol: "en-US", Name: "American English"}
}

// PtBR locale
func (rep *Locales) PtBR() *models.Locale {
	return &models.Locale{Symbol: "pt-BR", Name: "Português Brasileiro"}
}

// EnUS locale
func (rep *Locales) EnUS() *models.Locale {
	return &models.Locale{Symbol: "en-US", Name: "American English"}
}

// EnGB locale
func (rep *Locales) EnGB() *models.Locale {
	return &models.Locale{Symbol: "en-GB", Name: "British English"}
}
