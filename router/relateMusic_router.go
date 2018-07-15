package router

import (
	"github.com/sirupsen/logrus"
	"musicProject/gin_router"
	"github.com/gin-gonic/gin"
	"musicProject/handler"
)

func init() {
	logrus.Info("router relateMusic init")
    gin_router.GetEngine().GET("/relateMusic", func(context *gin.Context) {
		handler.HandlerAddOrUpdateInterface(context,&RelateMusicWithHeader{},"get")
	})
	gin_router.GetEngine().PUT("/relateMusic/update", func(context *gin.Context) {
		handler.HandlerAddOrUpdateInterface(context,&RelateMusicWithHeader{},"update")
	})
	gin_router.GetEngine().PUT("/relateMusic/add", func(context *gin.Context) {
		handler.HandlerAddOrUpdateInterface(context,&RelateMusicWithHeader{},"add")
	})
	gin_router.GetEngine().DELETE("/relateMusic/delete", func(context *gin.Context) {
		handler.HandlerAddOrUpdateInterface(context,&RelateMusicWithHeader{},"delete")
	})
	gin_router.GetEngine().POST("/relateMusic/search", func(context *gin.Context) {
		handler.HandlerAddOrUpdateInterface(context,&RelateMusicWithHeader{},"search")
	})
}

type RelateMusicWithHeader struct {
	handler.Pagination
	handler.RelateMusic
}
