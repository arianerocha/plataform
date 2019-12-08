package htmlrender

import (
	gintemplate "github.com/foolin/gin-template"
	"github.com/gin-gonic/gin/render"
	"github.com/krakenlab/plataform/internal/configs"
	"github.com/krakenlab/plataform/internal/servers/web/helpers"
)

// New render engine
func New() render.HTMLRender {
	return gintemplate.New(
		gintemplate.TemplateConfig{
			Root:      configs.Path().Views(),
			Extension: ".html",
			Master:    "layouts/master",
			Partials: []string{
				"layouts/partials",
				"layouts/navbar",
				"layouts/footbar",
			},
			Funcs:        helpers.Helpers(),
			DisableCache: true,
		},
	)
}
