package di

import (
	"github.com/IcaroSilvaFK/developer_academy_mvc/application/http/controllers"
	"github.com/IcaroSilvaFK/developer_academy_mvc/application/services"
	"github.com/IcaroSilvaFK/developer_academy_mvc/application/utils"
	"github.com/IcaroSilvaFK/developer_academy_mvc/infra/database"
	"github.com/IcaroSilvaFK/developer_academy_mvc/infra/repositories"
)

func NewLoginController() controllers.LoginControllerInterface {

	db := database.GetConnection()
	repo := repositories.NewUserRepository(db)
	chttp := utils.NewHttpClient()
	svc := services.NewAuthService(
		repo,
		chttp,
	)

	return controllers.NewLoginController(svc)
}
