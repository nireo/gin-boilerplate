package api

import (
	"github.com/gin-gonic/gin"
	"github.com/nireo/gin-boilerplate/api/routes/auth"
)

// ApplyRoutes adds router to gin engine
func ApplyRoutes(r *gin.Engine) {
	routes := r.Group("/api")
	{
		auth.ApplyRoutes(routes)
	}
}
