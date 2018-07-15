package router

import (
	"github.com/gin-gonic/gin"
	"musicProject/gin_router"
	"musicProject/handler"
)

func init() {
	gin_router.GetEngine().GET("/music", func(context *gin.Context) {
		handler.HandlerAddOrUpdateInterface(context, &MusicWithHeader{}, "post")
	})
	gin_router.GetEngine().DELETE("/music/delete", func(context *gin.Context) {
		handler.HandlerAddOrUpdateInterface(context, &handler.Music{}, "delete")
	})
	gin_router.GetEngine().PUT("/music/update", func(context *gin.Context) {
		handler.HandlerAddOrUpdateInterface(context, &handler.Music{}, "update")
	})
	gin_router.GetEngine().PUT("/music/add", func(context *gin.Context) {
		handler.HandlerAddOrUpdateInterface(context, &handler.Music{}, "add")
	})
	gin_router.GetEngine().POST("/music/search", func(context *gin.Context) {
		handler.HandlerAddOrUpdateInterface(context, &handler.Music{}, "search")
	})
}

type MusicWithHeader struct {
	handler.Pagination
	handler.Music
}
