package di

import (
	"github.com/IcaroSilvaFK/developer_academy_mvc/application/http/controllers"
	"github.com/IcaroSilvaFK/developer_academy_mvc/application/services"
	"github.com/IcaroSilvaFK/developer_academy_mvc/infra/database"
	"github.com/IcaroSilvaFK/developer_academy_mvc/infra/repositories"
)

func NewCommentsChallengeController() controllers.CommentsChallengeControllerInterface {

	db := database.GetConnection()

	repo := repositories.NewCommentChallengeRepository(db)

	cache := services.NewCacheService()
	svc := services.NewCommentChallengeServicer(repo, cache)

	return controllers.NewCommentsChallengeController(svc)
}
