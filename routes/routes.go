package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/lbodlev888/url_shortener/controllers"
)

func SetupRoutes(r *gin.Engine) {
	r.GET("/", controllers.IndexPage)
	r.GET("/login", controllers.LoginPage)
	r.GET("/register", controllers.RegisterPage)

	api := r.Group("/api")
	{
		api.POST("/login", controllers.LoginUser)
		api.POST("/register", controllers.RegisterUser)

		api.POST("/short", controllers.NewShortUrl)
	}
}
