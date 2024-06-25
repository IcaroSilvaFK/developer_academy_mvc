package di

import (
	"github.com/IcaroSilvaFK/developer_academy_mvc/application/http/controllers"
	"github.com/IcaroSilvaFK/developer_academy_mvc/application/services"
	"github.com/IcaroSilvaFK/developer_academy_mvc/infra/database"
	"github.com/IcaroSilvaFK/developer_academy_mvc/infra/repositories"
)

func NewChallengesCategoriesController() controllers.ChallengesCategoriesControllerInterface {

	db := database.GetConnection()
	repo := repositories.NewChallengesCategoriesRepository(db)

	return controllers.NewChallengesCategoriesController(
		services.NewChallengesCategoriesService(repo),
	)
}
