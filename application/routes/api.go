package routes

import (
	"net/http"

	"github.com/IcaroSilvaFK/developer_academy_mvc/application/di"
	"github.com/gin-gonic/gin"
)

func NewApiRoutes(engine *gin.Engine) {

	group := engine.Group("/api/v1")

	loginController := di.NewLoginController()
	createChallengeController := di.NewCreateNewChallengeController()

	group.GET("/heath", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"check": true,
		})
	})

	group.GET("/login/:code", loginController.SignIn)

	group.POST("/challenges", createChallengeController.Create)
}
