package main

import (
	"net/http"

	"github.com/rachanonk/go-jwt/src/controller"
	"github.com/rachanonk/go-jwt/src/middleware"
	"github.com/rachanonk/go-jwt/src/service"

	"github.com/gin-gonic/gin"
)

func main() {
	var loginService service.LoginService = service.StaticLoginService()
	var jwtService service.JWTService = service.JWTAuthService()
	var loginController controller.LoginController = controller.LoginHandler(loginService, jwtService)

	server := gin.Default()

	server.GET("/hello", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"hello": "Hey, there.",
		})
	})

	server.POST("/login", func(ctx *gin.Context) {
		token := loginController.Login(ctx)
		if token != "" {
			ctx.JSON(http.StatusOK, gin.H{
				"token": token,
			})
		} else {
			ctx.JSON(http.StatusUnauthorized, nil)
		}
	})

	api := server.Group("/api")

	api.Use(middleware.AuthorizeJWT())
	{
		api.GET("/test", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{"message": "success"})
		})
	}

	port := "5000"
	server.Run(":" + port)

}
