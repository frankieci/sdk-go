package api

import (
	"net/http"

	"gitlab.com/sdk-go/wire"

	"gitlab.com/sdk-go/library/builder"
	"gitlab.com/sdk-go/service"

	"github.com/gin-gonic/gin"
)

type infoCollectorApi struct {
	svc *service.InfoManageService
}

func newInfoCollectorApi() *infoCollectorApi {
	return &infoCollectorApi{svc: service.NewInfoManageService()}
}

func addRouter(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	g.Use(mw...)
	collectorApi := newInfoCollectorApi()
	g1 := g.Group("/sys")
	{
		g1.GET("/info", collectorApi.getInfoByType)
	}
	return g
}

func (api *infoCollectorApi) getInfoByType(c *gin.Context) {
	var condParam wire.InfoType
	if err := c.ShouldBindQuery(&condParam); err != nil {
		builder.BuildBindError(c, err)
		return
	}
	data, err := api.svc.GetInfo(&condParam)
	if err != nil {
		builder.BuildError(c, http.StatusInternalServerError, err.Error())
		return
	}
	builder.BuildSuccessWithData(c, http.StatusOK, data)
}
