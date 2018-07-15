package router

import (
	"musicProject/handler"
	"musicProject/gin_router"
	"github.com/gin-gonic/gin"
)

func init() {
	gin_router.GetEngine().GET("/music", func(context *gin.Context) {
		handler.HandlerAddOrUpdateInterface(context,&MusicWithHeader{},"post")
	})
	gin_router.GetEngine().DELETE("/music/delete", func(context *gin.Context) {
		handler.HandlerAddOrUpdateInterface(context,&MusicWithHeader{},"delete")
	})
	gin_router.GetEngine().PUT("/music/update", func(context *gin.Context) {
		handler.HandlerAddOrUpdateInterface(context,&MusicWithHeader{},"update")
	})
	gin_router.GetEngine().PUT("/music/add", func(context *gin.Context) {
		handler.HandlerAddOrUpdateInterface(context,&MusicWithHeader{},"add")
	})
	gin_router.GetEngine().POST("/music/search", func(context *gin.Context) {
		handler.HandlerAddOrUpdateInterface(context,&MusicWithHeader{},"search")
	})
}

type MusicWithHeader struct {
	handler.Pagination
	handler.Music
}