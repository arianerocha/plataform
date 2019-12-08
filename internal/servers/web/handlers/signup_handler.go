package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/krakenlab/plataform/internal/repositories"
	"net/http"
)

// SignUp handler
type SignUp struct{}

// NewSignUp handler
func NewSignUp() Handler {
	return &SignUp{}
}

// Setup SignUp handler
func (h *SignUp) Setup(engine *gin.Engine) {
	engine.GET(SignUpPath, h.GET)
}

// GET at SignUpPath
func (h *SignUp) GET(c *gin.Context) {
	locale := repositories.NewLocales().LocaleOrDefault(c.Param("locale"))
	c.HTML(http.StatusOK, "auth/signup", gin.H{"locale": locale.Symbol})
}
