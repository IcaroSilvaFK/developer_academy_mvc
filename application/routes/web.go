package routes

import (
	"github.com/IcaroSilvaFK/developer_academy_mvc/application/di"
	"github.com/gin-gonic/gin"
)

func NewWebRoutes(en *gin.Engine) {

	lController := di.NewLoginController()
	hController := di.NewHomeController()
	ccController := di.CreateNewChallengeController()

	en.GET("/", lController.Login)
	en.GET("/home", hController.Index)
	en.GET("/new_challenge", ccController.Index)
}
