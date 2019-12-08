package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Root handler
type Root struct{}

// NewRoot handler
func NewRoot() Handler {
	return &Root{}
}

// Setup root handler
func (h *Root) Setup(engine *gin.Engine) {
	engine.GET(RootPath, h.GET)
}

// GET at RootPath
func (h *Root) GET(c *gin.Context) {
	c.Redirect(http.StatusTemporaryRedirect, "/404")
}
