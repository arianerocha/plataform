package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/krakenlab/plataform/internal/repositories"
	"net/http"
)

// LocalizedSignUp handler
type LocalizedSignUp struct{}

// NewLocalizedSignUp handler
func NewLocalizedSignUp() Handler {
	return &LocalizedSignUp{}
}

// Setup LocalizedSignUp handler
func (h *LocalizedSignUp) Setup(engine *gin.Engine) {
	engine.GET(LocalizedSignUpPath, h.GET)
}

// GET at LocalizedSignUpPath
func (h *LocalizedSignUp) GET(c *gin.Context) {
	locale := repositories.NewLocales().LocaleOrDefault(c.Param("locale"))
	c.HTML(http.StatusOK, "auth/signup", gin.H{"locale": locale.Symbol})
}
