package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/krakenlab/plataform/internal/models"
	"github.com/krakenlab/plataform/internal/repositories"
	"github.com/krakenlab/plataform/internal/services"
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
	engine.POST(LocalizedSignUpPath, h.POST)
}

// GET at LocalizedSignUpPath
func (h *LocalizedSignUp) GET(c *gin.Context) {
	locale := repositories.NewLocales().LocaleOrDefault(c.Param("locale"))
	h.renderForm(c, locale, []error{})
}

// POST at LocalizedSignUpPath
func (h *LocalizedSignUp) POST(c *gin.Context) {
	locale := repositories.NewLocales().LocaleOrDefault(c.Param("locale"))

	service := services.NewSignUp(
		c.PostForm("email"),
		c.PostForm("password"),
		c.PostForm("repeat_password"),
		locale.Symbol,
		c.PostForm("terms") == "on",
		c.PostForm("privacy") == "on",
		c.PostForm("newslatter") == "on",
	)

	if service.Valid() {
		if service.Create() == nil {
			h.renderSuccess(c, locale)
			return
		}
	}

	h.renderForm(c, locale, service.Errors())
}

func (h *LocalizedSignUp) renderForm(c *gin.Context, locale *models.Locale, errors []error) {
	c.HTML(http.StatusOK, "signup/get", gin.H{
		"locale": locale.Symbol,
		"errors": errors,
	})
}

func (h *LocalizedSignUp) renderSuccess(c *gin.Context, locale *models.Locale) {
	c.HTML(http.StatusCreated, "signup/post", gin.H{
		"locale": locale.Symbol,
	})
}
