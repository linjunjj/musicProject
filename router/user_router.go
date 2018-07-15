package router

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"musicProject/gin_router"
	"musicProject/handler"
)

func init() {
	logrus.Info("router account_master init")
	gin_router.GetEngine().GET("/user", func(context *gin.Context) {
		handler.HandlerAddOrUpdateInterface(context, &UserWithHeader{}, "get")
	})
	gin_router.GetEngine().DELETE("/user/delete", func(context *gin.Context) {
		handler.HandlerAddOrUpdateInterface(context, &handler.User{}, "delete")
	})
	gin_router.GetEngine().PUT("/user/update", func(context *gin.Context) {
		handler.HandlerAddOrUpdateInterface(context, &handler.User{}, "update")
	})
	gin_router.GetEngine().PUT("/user/add", func(context *gin.Context) {
		handler.HandlerAddOrUpdateInterface(context, &handler.User{}, "add")
	})
	gin_router.GetEngine().POST("/user/search", func(context *gin.Context) {
		handler.HandlerAddOrUpdateInterface(context, &handler.User{}, "search")
	})
}

type UserWithHeader struct {
	handler.Pagination
	handler.User
}
