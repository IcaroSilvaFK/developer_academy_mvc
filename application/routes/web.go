package routes

import (
	"github.com/IcaroSilvaFK/developer_academy_mvc/application/di"
	"github.com/gin-gonic/gin"
)

func NewWebRoutes(en *gin.Engine) {

	lController := di.NewLoginController()

	en.GET("/", lController.Login)
}
