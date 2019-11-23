package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/nireo/gin-boilerplate/middlewares"
)

func ApplyRoutes(r *gin.RouterGroup) {
	auth := r.Group("/auth")
	{
		auth.POST("/register", register)
		auth.POST("/login", login)
		auth.DELETE("/", middlewares.Authorized, remove)
	}
}