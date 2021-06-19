package Routes

import (
	"mvc_pattern/Controllers"

	"github.com/gin-gonic/gin"
)

// Setup router
func SetupRouter() *gin.Engine {
	route := gin.Default()

	grp1 := route.Group("/api")
	{
		grp1.GET("users", Controllers.GetUsers)
		grp1.POST("user", Controllers.CreateUser)
		grp1.GET("user/:id", Controllers.GetUserByID)
		grp1.PUT("user/:id", Controllers.UpdateUser)
		grp1.DELETE("user/:id", Controllers.DeleteUser)
	}

	return route
}
