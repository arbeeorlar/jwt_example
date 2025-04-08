package routes

import (
	"github.com/arbeeorlar/jwt-example/controllers"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.Engine) {

	public := router.Group("/auth")
	{
		public.POST("/login", controllers.LoginController)
		public.POST("/signup", controllers.SignUpController)
	}

	private := router.Group("/user")
	private.Use(controllers.AuthenticationMiddleware())
	{
		private.GET("/premium", controllers.PremiumController)
		private.GET("/signout", controllers.SignOutController)
		private.GET("/home", controllers.HomeController)
	}

}
