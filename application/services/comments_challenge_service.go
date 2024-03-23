package services

import (
	"github.com/IcaroSilvaFK/developer_academy_mvc/infra/models"
	"github.com/IcaroSilvaFK/developer_academy_mvc/infra/repositories"
)

type CommentChallengeService struct {
	repo repositories.CommentChallengeRepositoryInterface
}

type CommentChallengeServiceInterface interface {
	Create(challengeId, userId, comment string) (*models.ChallengeCommentModel, error)
	FindById(commentId string) (*models.ChallengeCommentModel, error)
	FindByUserId(userId string) ([]*models.ChallengeCommentModel, error)
	FindByChallengeId(challengeId string) ([]*models.ChallengeCommentModel, error)
	Delete(commentId string) error
}

func NewCommentChallengeServicer(
	repo repositories.CommentChallengeRepositoryInterface,
) CommentChallengeServiceInterface {
	return &CommentChallengeService{
		repo,
	}
}

func (cs *CommentChallengeService) Create(challengeId, userId, comment string) (*models.ChallengeCommentModel, error) {

	c := models.NewChallengeCommentModel(challengeId, userId, comment)

	err := cs.repo.Create(c)

	return c, err
}
func (cs *CommentChallengeService) FindById(commentId string) (*models.ChallengeCommentModel, error) {
	return cs.repo.FindById(commentId)
}

func (cs *CommentChallengeService) FindByUserId(userId string) ([]*models.ChallengeCommentModel, error) {

	return cs.repo.FindByUserId(userId)
}

func (cs *CommentChallengeService) FindByChallengeId(challengeId string) ([]*models.ChallengeCommentModel, error) {
	return cs.repo.FindByChallengeId(challengeId)
}
func (cs *CommentChallengeService) Delete(commentId string) error {
	return cs.repo.Delete(commentId)
}
