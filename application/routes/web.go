package routes

import (
	"github.com/IcaroSilvaFK/developer_academy_mvc/application/di"
	"github.com/gin-gonic/gin"
)

func NewWebRoutes(en *gin.Engine) {

	lController := di.NewLoginController()
	hController := di.NewHomeController()
	ccController := di.NewCreateNewChallengeController()
	cController := di.NewChallengeController()
	pController := di.NewProfileController()

	en.GET("/", lController.Login)
	en.GET("/home", hController.Index)
	en.GET("/new_challenge", ccController.Index)
	en.GET("/challenge/:id", cController.Index)
	en.GET("/profile/:id", pController.Index)
}
