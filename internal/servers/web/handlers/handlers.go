package handlers

import "github.com/gin-gonic/gin"

// Handler can setup a gin engine
type Handler interface {
	Setup(*gin.Engine)
}

// Handlers export all handlers of this application
func Handlers() []Handler {
	return []Handler{
		NewRoot(),
	}
}
