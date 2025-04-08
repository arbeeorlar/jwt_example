package routes

import (
	"github.com/arbeeorlar/jwt-example/controllers"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.Engine) {
	router.POST("/login", controllers.LoginController)
	router.GET("/premium", controllers.PremiumController)
	router.POST("/signup", controllers.SignUpController)
	router.GET("/signout", controllers.SignOutController)
	router.GET("/home", controllers.HomeController)
}
