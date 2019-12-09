package handlers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/krakenlab/plataform/internal/repositories"
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
	c.Redirect(http.StatusPermanentRedirect, strings.ReplaceAll(LocalizedSignUpPath, ":locale", repositories.NewLocales().Default().Symbol))
}
