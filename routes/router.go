package routes

import (
	v1 "ginblog/api/v1"
	"ginblog/middleware"
	"ginblog/utils"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.New()
	r.Use(middleware.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.Cors())

	auth := r.Group("api/v1")
	auth.Use(middleware.JwtToken())
	// auth.Use(middleware.Cors())
	{
		// 用户 模块的路由接口
		// auth.POST("user/add", v1.AddUser)
		auth.PUT("user/:id", v1.EditUser)
		auth.DELETE("user/:id", v1.DeleteUser)
		auth.GET("user/:id", v1.GetUser)
		// 分类 模块的路由接口
		auth.POST("category/add", v1.AddCategory)
		auth.POST("category/:id", v1.EditCategory)
		auth.DELETE("category/:id", v1.DeleteCategory)
		// 文章 模块的路由接口
		auth.POST("article/add", v1.AddArticle)
		auth.POST("article/:id", v1.EditArticle)
		auth.DELETE("article/:id", v1.DeleteArticle)
		// v1.GET("hello", func(c *gin.Context) {
		// 	c.JSON(http.StatusOK, gin.H{
		// 		"msg": "ok",
		// 	})
		// })
	}
	router := r.Group("api/v1")
	// router.Use(middleware.Cors())
	{
		router.POST("user/add", v1.AddUser)
		router.GET("users", v1.GetUsers)
		router.GET("categorys", v1.GetCategorys)
		router.GET("articles", v1.GetArticles)
		router.POST("login", v1.Login)
		// router.GET("article/list/:id",v1.GetArticleI)
	}
	_ = r.Run(utils.HttpPort)
}
