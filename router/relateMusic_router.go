package router

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"musicProject/gin_router"
	"musicProject/handler"
	"net/http"
)

func init() {
	logrus.Info("router relateMusic init")
	gin_router.GetEngine().POST("/relateMusic", func(context *gin.Context) {
		handler.HandlerQueryInterface(context, &RelateMusicWithHeader{}, "get")
	})
	gin_router.GetEngine().PUT("/relateMusic/update", func(context *gin.Context) {
		handler.HandlerAddOrUpdateInterface(context, &handler.RelateMusic{}, "update")
	})
	gin_router.GetEngine().PUT("/relateMusic/add", func(context *gin.Context) {
		handler.HandlerAddOrUpdateInterface(context, &handler.RelateMusic{}, "add")
	})
	gin_router.GetEngine().DELETE("/relateMusic/delete", func(context *gin.Context) {
		handler.HandlerAddOrUpdateInterface(context, &handler.RelateMusic{}, "delete")
	})
	gin_router.GetEngine().POST("/relateMusic/search", func(context *gin.Context) {
		handler.HandlerAddOrUpdateInterface(context, &handler.RelateMusic{}, "search")
	})
	gin_router.GetEngine().GET("/healthcheck", func(context *gin.Context) {
		context.String(http.StatusOK, "ok")
	})
}

type RelateMusicWithHeader struct {
	handler.Pagination
	handler.RelateMusic
}
