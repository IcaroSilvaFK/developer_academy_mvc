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
	urepo := repositories.NewUserRepository(db)
	challengerepo := repositories.NewChallengeRepository(db)
	hintrepo := repositories.NewChallengesHintsRepository(db)

	hintservice := services.NewChallengeHintService(hintrepo)
	iaservice := services.NewAIService()
	usvc := services.NewUserService(urepo)

	chttp := utils.NewHttpClient()

	challengeservice := services.NewChallengeService(challengerepo, hintservice, iaservice)
	sessionservice := services.NewSessionService()

	svc := services.NewAuthService(
		repo,
		chttp,
	)

	return controllers.NewLoginController(svc, usvc, challengeservice, sessionservice)
}
