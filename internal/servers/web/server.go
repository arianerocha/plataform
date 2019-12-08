package web

import (
	"github.com/krakenlab/plataform/internal/configs"
	"github.com/krakenlab/plataform/internal/servers/web/handlers"
	"github.com/krakenlab/plataform/internal/servers/web/helpers"

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
	server.SetupStatic()
	server.SetupHandlers()
	server.SetupHTMLRender()
}

// SetupStatic engine
func (server *Server) SetupStatic() {
	server.engine.Static("static", configs.Path().Static())
}

// SetupHandlers engine
func (server *Server) SetupHandlers() {
	for _, handler := range handlers.Handlers() {
		handler.Setup(server.engine)
	}
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
			Funcs:        helpers.Helpers(),
			DisableCache: true,
		},
	)
}

// Run the engine
func (server *Server) Run() {
	panic(server.engine.Run("0.0.0.0:3000"))
}
