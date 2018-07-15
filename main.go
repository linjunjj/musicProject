package main

import (
	"github.com/gin-gonic/gin"
	"musicProject/config"
	"musicProject/db"
	"musicProject/gin_router"
	"net/http"
)

func init() {
	err := config.Init()
	if err != nil {
		panic(err)
	}
	db.Init_mysql()
}
func main() {
	gin_router.GetEngine().GET("/healthcheck", func(context *gin.Context) {
		context.String(http.StatusOK, "ok")
	})
	gin_router.GetEngine().Run(":8080")

}
