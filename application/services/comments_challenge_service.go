package services

import (
	"github.com/IcaroSilvaFK/developer_academy_mvc/application/utils"
	"github.com/IcaroSilvaFK/developer_academy_mvc/infra/models"
	"github.com/IcaroSilvaFK/developer_academy_mvc/infra/repositories"
	"gorm.io/gorm"
)

type CommentChallengeService struct {
	repo repositories.CommentChallengeRepositoryInterface
}

type CommentChallengeServiceInterface interface {
	Create(challengeId, userId, comment string) (*models.ChallengeCommentModel, *utils.RestErr)
	FindById(commentId string) (*models.ChallengeCommentModel, *utils.RestErr)
	FindByUserId(userId string) ([]*models.ChallengeCommentModel, *utils.RestErr)
	FindByChallengeId(challengeId string) ([]*models.ChallengeCommentModel, *utils.RestErr)
	Delete(commentId string) *utils.RestErr
}

func NewCommentChallengeServicer(
	repo repositories.CommentChallengeRepositoryInterface,
) CommentChallengeServiceInterface {
	return &CommentChallengeService{
		repo,
	}
}

func (cs *CommentChallengeService) Create(challengeId, userId, comment string) (*models.ChallengeCommentModel, *utils.RestErr) {

	c := models.NewChallengeCommentModel(challengeId, userId, comment)

	err := cs.repo.Create(c)

	if err != nil {

		message := "Error on create comment"

		return nil, utils.NewInternalServerError(&message)
	}

	return c, nil
}
func (cs *CommentChallengeService) FindById(commentId string) (*models.ChallengeCommentModel, *utils.RestErr) {

	r, err := cs.repo.FindById(commentId)

	if err == gorm.ErrRecordNotFound {
		return nil, utils.NewNotFoundException("The id impproved not exists")
	}

	if err != nil {
		message := "Error on find comment"
		return nil, utils.NewInternalServerError(&message)
	}

	return r, nil
}

func (cs *CommentChallengeService) FindByUserId(userId string) ([]*models.ChallengeCommentModel, *utils.RestErr) {

	r, err := cs.repo.FindByUserId(userId)

	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}

	if err != nil {
		message := "Error on find user comments"

		return nil, utils.NewInternalServerError(&message)
	}

	return r, nil
}

func (cs *CommentChallengeService) FindByChallengeId(challengeId string) ([]*models.ChallengeCommentModel, *utils.RestErr) {

	r, err := cs.repo.FindByChallengeId(challengeId)

	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}

	if err != nil {
		message := "Error on find challenge comments"
		return nil, utils.NewInternalServerError(&message)
	}

	return r, nil
}
func (cs *CommentChallengeService) Delete(commentId string) *utils.RestErr {

	err := cs.repo.Delete(commentId)

	if err == gorm.ErrRecordNotFound {
		return utils.NewNotFoundException("Item not found")
	}

	if err != nil {
		message := "Error on delete comment"
		return utils.NewInternalServerError(&message)
	}

	return nil
}
