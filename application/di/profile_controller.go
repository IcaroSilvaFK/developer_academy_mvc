package di

import (
	"github.com/IcaroSilvaFK/developer_academy_mvc/application/http/controllers"
	"github.com/IcaroSilvaFK/developer_academy_mvc/application/services"
	"github.com/IcaroSilvaFK/developer_academy_mvc/infra/database"
	"github.com/IcaroSilvaFK/developer_academy_mvc/infra/repositories"
)

func NewProfileController() controllers.ProfileControllerInterface {

	db := database.GetConnection()
	repo := repositories.NewChallengeRepository(db)
	hintRepo := repositories.NewChallengesHintsRepository(db)
	hintSvc := services.NewChallengeHintService(hintRepo)
	aiservice := services.NewAIService()
	svc := services.NewChallengeService(repo, hintSvc, aiservice)

	return controllers.NewProfileController(svc)
}
