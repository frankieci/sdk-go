package api

import (
	"github.com/gin-gonic/gin"
)

func Register(g *gin.Engine, mw ...gin.HandlerFunc) {
	routeRegister(g, mw...)
}

func routeRegister(g *gin.Engine, mw ...gin.HandlerFunc) {
	addRouter(g, mw...)
}
