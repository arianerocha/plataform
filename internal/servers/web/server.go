package web

import (
	"github.com/krakenlab/plataform/internal/configs"

	gintemplate "github.com/foolin/gin-template"
	"github.com/gin-gonic/gin"
)

// Server engine
type Server struct {
	engine *gin.Engine
}

// NewServer generate a default server without handlers.
func NewServer() *Server {
	engine := gin.Default()
	return &Server{engine: engine}
}

// Setup routes, sub-engines, session, etc
func (server *Server) Setup() {
	server.SetupHTMLRender()
}

// SetupHTMLRender engine
func (server *Server) SetupHTMLRender() {
	server.engine.HTMLRender = gintemplate.New(
		gintemplate.TemplateConfig{
			Root:      configs.Path().Views(),
			Extension: ".html",
			Master:    "layouts/master",
			Partials: []string{
				"layouts/partials",
				"layouts/navbar",
				"layouts/footbar",
			},
			// Funcs:        funcs.Funcs(),
			DisableCache: true,
		},
	)
}
