package rounter

import (
	"exchange_app/controllers"
	"exchange_app/middleware"

	"github.com/gin-gonic/gin"
)


func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Accept, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	api := r.Group("api")
	api.Use(middleware.JWTAuth())
	{
		api.GET("/exchangeRates", controllers.GetExchangeRates)
		api.POST("/exchangeRate", controllers.CreateExchangeRate)
	}

	user := r.Group("user")
	{
		user.POST("/register", controllers.Register)
		user.POST("/login", controllers.Login)
	}

	article := r.Group("article")
	article.Use(middleware.JWTAuth())
	{
		article.POST("/create", controllers.CreateArticle)
		article.GET("/get", controllers.GetArtiles)
		article.DELETE("/delete/:id", controllers.DeleteArtile)
	}

	like := r.Group("like")
	like.Use(middleware.JWTAuth())
	{
		like.POST("/toggle/:id", controllers.ToggleLike)
		like.GET("/get/:id", controllers.GetLikes)
	}

	return r;
}