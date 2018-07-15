package router

import (
	"github.com/sirupsen/logrus"
	"musicProject/gin_router"
	"github.com/gin-gonic/gin"
	"musicProject/handler"
)

func init() {
	logrus.Info("router account_master init")
     gin_router.GetEngine().GET("/user", func(context *gin.Context) {
		 handler.HandlerAddOrUpdateInterface(context,&UserWithHeader{},"get")
	 })
	gin_router.GetEngine().DELETE("/user/delete", func(context *gin.Context) {
		handler.HandlerAddOrUpdateInterface(context,&UserWithHeader{},"delete")
	})
	gin_router.GetEngine().PUT("/user/update", func(context *gin.Context) {
		handler.HandlerAddOrUpdateInterface(context,&UserWithHeader{}, "update")
	})
	gin_router.GetEngine().PUT("/user/add", func(context *gin.Context) {
		handler.HandlerAddOrUpdateInterface(context,&UserWithHeader{},"add")
	})
	gin_router.GetEngine().POST("/user/search", func(context *gin.Context) {
		handler.HandlerAddOrUpdateInterface(context,&UserWithHeader{},"search")
	})
}

type  UserWithHeader struct {
	handler.Pagination
	handler.User
}