package middleware

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Cors() gin.HandlerFunc {
	return cors.New(
		cors.Config{
			AllowOrigins:     []string{"*"},
			AllowMethods:     []string{"GET", "POST", "DELETE", "PUT", "OPTIONS"},
			AllowHeaders:     []string{"*"},
			ExposeHeaders:    []string{"Content-Length", "Authorization", "Content-Type"},
			AllowCredentials: true,
			MaxAge:           12 * time.Hour,
		},
	)
	// return func(c *gin.Context) {
	// 	method := c.Request.Method
	// 	c.Header("Access-Control-Allow-Origin", "")
	// 	c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
	// 	c.Header("Access-Control-Allow-Methods", "*")
	// 	c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
	// 	// c.Header("Access-Control-Allow-Credentials", "true")
	// 	//放行所有OPTIONS方法
	// 	if method == "OPTIONS" {
	// 		c.AbortWithStatus(http.StatusNoContent)
	// 	}
	// 	// 处理请求
	// 	c.Next()
	// }
}
