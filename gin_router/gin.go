package gin_router

import (
	"github.com/gin-gonic/gin"
	"github.com/Sirupsen/logrus"
)

var r *gin.Engine

func init() {
	logrus.Info("gin engine init ")
	r = gin.New()
	// 不打印 /healthcheck的请求日志
	r.Use(gin.LoggerWithWriter(gin.DefaultWriter, "/healthcheck"))
	r.Use(gin.Recovery())
	// 增加跨域支持
	r.Use(CORSMiddleware())
}

func GetEngine() *gin.Engine {
	return r
}

//跨域中间件
func CORSMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT,DELETE,PATCH")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}
