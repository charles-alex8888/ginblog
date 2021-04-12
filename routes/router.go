package routes

import (
	v1 "ginblog/api/v1"
	"ginblog/utils"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.Default()

	router := r.Group("api/v1")
	{
		// 用户 模块的路由接口
		router.POST("user/add", v1.AddUser)
		router.GET("users", v1.GetUsers)
		router.GET("user/:username", v1.GetUser)
		router.POST("user/:id", v1.EditUser)
		router.DELETE("user/:id", v1.DeleteUser)
		// 分类 模块的路由接口
		router.POST("category/add", v1.AddCategory)
		router.GET("categorys", v1.GetCategorys)
		router.POST("category/:id", v1.EditCategory)
		router.DELETE("category/:id", v1.DeleteCategory)
		// 文章 模块的路由接口
		router.POST("article/add", v1.AddArticle)
		router.GET("articles", v1.GetArticles)
		router.POST("article/:id", v1.EditArticle)
		router.DELETE("article/:id", v1.DeleteArticle)
		// v1.GET("hello", func(c *gin.Context) {
		// 	c.JSON(http.StatusOK, gin.H{
		// 		"msg": "ok",
		// 	})
		// })
	}
	r.Run(utils.HttpPort)
}
