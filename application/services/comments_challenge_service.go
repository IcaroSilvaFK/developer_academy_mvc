package services

import (
	"context"
	"fmt"

	"github.com/IcaroSilvaFK/developer_academy_mvc/application/utils"
	"github.com/IcaroSilvaFK/developer_academy_mvc/infra/models"
	"github.com/IcaroSilvaFK/developer_academy_mvc/infra/repositories"
	"gorm.io/gorm"
)

type CommentChallengeService struct {
	repo         repositories.CommentChallengeRepositoryInterface
	cache        CacheServiceInterface
	challengeKey string
}

type CommentChallengeServiceInterface interface {
	Create(ctx context.Context, challengeId, userId, comment string) (*models.ChallengeCommentModel, *utils.RestErr)
	FindById(ctx context.Context, commentId string) (*models.ChallengeCommentModel, *utils.RestErr)
	FindByUserId(ctx context.Context, userId string) ([]*models.ChallengeCommentModel, *utils.RestErr)
	FindByChallengeId(ctx context.Context, challengeId string) ([]*models.ChallengeCommentModel, *utils.RestErr)
	Delete(ctx context.Context, commentId string) *utils.RestErr
}

func NewCommentChallengeServicer(
	repo repositories.CommentChallengeRepositoryInterface,
	cache CacheServiceInterface,
) CommentChallengeServiceInterface {
	return &CommentChallengeService{
		repo, cache, "challenges",
	}
}

func (cs *CommentChallengeService) Create(ctx context.Context, challengeId, userId, comment string) (*models.ChallengeCommentModel, *utils.RestErr) {

	c := models.NewChallengeCommentModel(challengeId, userId, comment)
	cacheKey := fmt.Sprintf("%s-%s", cs.challengeKey, challengeId)

	if err := cs.cache.Delete(cacheKey); err != nil {
		utils.Error(fmt.Sprintf("Error on delete %s challenge from cache", challengeId), err)
	}
	err := cs.repo.Create(ctx, c)

	if err != nil {

		message := "Error on create comment"

		return nil, utils.NewInternalServerError(&message)
	}

	return c, nil
}
func (cs *CommentChallengeService) FindById(ctx context.Context, commentId string) (*models.ChallengeCommentModel, *utils.RestErr) {

	r, err := cs.repo.FindById(ctx, commentId)

	if err == gorm.ErrRecordNotFound {
		return nil, utils.NewNotFoundException(
			fmt.Sprintf("The record with id %s not exists", commentId),
		)
	}

	if err != nil {
		message := "Error on find comment"
		return nil, utils.NewInternalServerError(&message)
	}

	return r, nil
}

func (cs *CommentChallengeService) FindByUserId(ctx context.Context, userId string) ([]*models.ChallengeCommentModel, *utils.RestErr) {

	r, err := cs.repo.FindByUserId(ctx, userId)

	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}

	if err != nil {
		message := "Error on find user comments"

		return nil, utils.NewInternalServerError(&message)
	}

	return r, nil
}

func (cs *CommentChallengeService) FindByChallengeId(ctx context.Context, challengeId string) ([]*models.ChallengeCommentModel, *utils.RestErr) {

	r, err := cs.repo.FindByChallengeId(ctx, challengeId)

	if err == gorm.ErrRecordNotFound {
		return nil, utils.NewNotFoundException(
			fmt.Sprintf("The record with id %s not exists", challengeId),
		)
	}

	if err != nil {
		message := "Error on find challenge comments"
		return nil, utils.NewInternalServerError(&message)
	}

	return r, nil
}
func (cs *CommentChallengeService) Delete(ctx context.Context, commentId string) *utils.RestErr {

	err := cs.repo.Delete(ctx, commentId)

	if err == gorm.ErrRecordNotFound {
		return utils.NewNotFoundException(
			fmt.Sprintf("The record with id %s not exists", commentId),
		)
	}

	if err != nil {
		message := "Error on delete comment"
		return utils.NewInternalServerError(&message)
	}

	return nil
}
