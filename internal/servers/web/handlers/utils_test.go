package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/krakenlab/plataform/internal/servers/web/htmlrender"
	"github.com/stretchr/testify/suite"
)

type handlerSuite struct {
	suite.Suite
	engine *gin.Engine
}

func (suite *handlerSuite) SetupTest() {
	suite.engine = gin.Default()
	suite.engine.HTMLRender = htmlrender.New()
}

func (suite *handlerSuite) ServeHTTP(rr http.ResponseWriter, req *http.Request) {
	suite.engine.ServeHTTP(rr, req)
}
