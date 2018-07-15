package router

import (
	"github.com/gin-gonic/gin"
	"musicProject/gin_router"
	"musicProject/handler"
)

func init() {
	//获取全部音乐
	gin_router.GetEngine().GET("/music", func(context *gin.Context) {
		handler.HandlerQueryInterface(context, &MusicWithHeader{}, "post")
	})


	//删除音乐
	gin_router.GetEngine().DELETE("/music/delete", func(context *gin.Context) {
		handler.HandlerAddOrUpdateInterface(context, &handler.Music{}, "delete")
	})
	//更新音乐
	gin_router.GetEngine().PUT("/music/update", func(context *gin.Context) {
		handler.HandlerAddOrUpdateInterface(context, &handler.Music{}, "update")
	})
	//添加音乐
	gin_router.GetEngine().PUT("/music/add", func(context *gin.Context) {
		handler.UploadImageInterface(context)
	})
	//搜索音乐
	gin_router.GetEngine().POST("/music/search", func(context *gin.Context) {
		handler.HandlerAddOrUpdateInterface(context, &handler.Music{}, "search")
	})
}

type MusicWithHeader struct {
	handler.Pagination
	handler.Music
}
