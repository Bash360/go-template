package modules

import (
	user "github.com/Bash360/go-template/internal/modules/users/handler"
	"github.com/gin-gonic/gin"
)

type RouterOptions struct {
	UserHandler *user.UserHandler
}

func SetUpRouter(options RouterOptions) *gin.Engine {

	router := gin.Default()
	v1 := router.Group("/api/v1")
	//add any app wide middleware

	users := v1.Group("/users")
	{
		users.GET("/", options.UserHandler.GetUsers)
		users.GET("/:id", options.UserHandler.GetUser)
		users.POST("/", options.UserHandler.AddUser)

	}

	return router
}
