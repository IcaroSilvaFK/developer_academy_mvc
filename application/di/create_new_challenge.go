package di

import (
	"github.com/IcaroSilvaFK/developer_academy_mvc/application/http/controllers"
	"github.com/IcaroSilvaFK/developer_academy_mvc/application/services"
	"github.com/IcaroSilvaFK/developer_academy_mvc/infra/database"
	"github.com/IcaroSilvaFK/developer_academy_mvc/infra/repositories"
)

func NewCreateNewChallengeController() controllers.CreateNewChallengeControllerInterface {

	db := database.GetConnection()
	repo := repositories.NewChallengeRepository(db)
	cService := services.NewChallengeService(repo)
	sessionService := services.NewSessionService()

	return controllers.NewCreateNewChallengeController(cService, sessionService)
}
