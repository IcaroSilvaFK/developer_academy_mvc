package di

import (
	"github.com/IcaroSilvaFK/developer_academy_mvc/application/http/controllers"
	"github.com/IcaroSilvaFK/developer_academy_mvc/application/services"
	"github.com/IcaroSilvaFK/developer_academy_mvc/infra/database"
	"github.com/IcaroSilvaFK/developer_academy_mvc/infra/repositories"
)

func NewHomeController() controllers.HomeControllerInterface {

	db := database.GetConnection()
	repo := repositories.NewChallengeRepository(db)
	svc := services.NewChallengeService(repo)

	return controllers.NewHomeController(svc)
}
