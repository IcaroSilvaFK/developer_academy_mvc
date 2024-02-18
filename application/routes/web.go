package routes

import (
	"github.com/IcaroSilvaFK/developer_academy_mvc/application/http/controllers"
	"github.com/gin-gonic/gin"
)

func NewWebRoutes(en *gin.Engine) {

	lController := controllers.NewLoginController()

	en.GET("/", lController.Login)
}
