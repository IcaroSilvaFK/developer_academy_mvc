package services

import (
	"context"
	"fmt"

	"github.com/IcaroSilvaFK/developer_academy_mvc/application/http/views"
	"github.com/IcaroSilvaFK/developer_academy_mvc/application/utils"
	"github.com/IcaroSilvaFK/developer_academy_mvc/infra/repositories"
	"gorm.io/gorm"
)

type ChallengesCategoriesService struct {
	repo repositories.ChallengesCategoriesRepositoryInterface
}

type ChallengesCategoriesServiceInterface interface {
	Create(ctx context.Context, input *views.ChallengesCategoriesInputView) *utils.RestErr
	GetAll(ctx context.Context, query string) ([]views.ChallengesCategoriesOutputView, *utils.RestErr)
	GetById(ctx context.Context, id string) (views.ChallengesCategoriesOutputView, *utils.RestErr)
	Update(ctx context.Context, id string, title string, userId string) *utils.RestErr
	FindByUserId(ctx context.Context, userId string) ([]views.ChallengesCategoriesOutputView, *utils.RestErr)
	Delete(ctx context.Context, id string) *utils.RestErr
}

func NewChallengesCategoriesService(
	repo repositories.ChallengesCategoriesRepositoryInterface,
) ChallengesCategoriesServiceInterface {

	return &ChallengesCategoriesService{repo}
}

func (c *ChallengesCategoriesService) Create(ctx context.Context, input *views.ChallengesCategoriesInputView) *utils.RestErr {

	err := c.repo.Create(ctx, input.Title, input.UserId)

	if err != nil {
		if err == gorm.ErrDuplicatedKey {
			return utils.NewBadRequestException(fmt.Sprintf("THE %s CATEGORY EXISTS.", input.Title))
		}

		return utils.NewInternalServerError(nil)
	}

	return nil
}

// Delete implements ChallengesCategoriesServiceInterface.
func (c *ChallengesCategoriesService) Delete(ctx context.Context, id string) *utils.RestErr {

	if err := c.repo.Delete(ctx, id); err != nil {
		if err == gorm.ErrRecordNotFound {
			m := fmt.Sprintf("THE RECORD WITH ID %s NOT EXISTS IN DATABASE.", id)

			return utils.NewInternalServerError(&m)
		}

		return utils.NewInternalServerError(nil)
	}

	return nil
}

// FindByUserId implements ChallengesCategoriesServiceInterface.
func (c *ChallengesCategoriesService) FindByUserId(ctx context.Context, userId string) ([]views.ChallengesCategoriesOutputView, *utils.RestErr) {
	r, err := c.repo.FindByUserId(ctx, userId)

	if err != nil {

		if err == gorm.ErrRecordNotFound {
			return nil, utils.NewBadRequestException(fmt.Sprintf("THE USER WITH ID %s, NOT HAVE CATEGORIES.", userId))
		}

		return nil, utils.NewInternalServerError(nil)
	}

	return views.NewChallengesCategoriesResponseListOutputView(r), nil
}

// GetAll implements ChallengesCategoriesServiceInterface.
func (c *ChallengesCategoriesService) GetAll(ctx context.Context, query string) ([]views.ChallengesCategoriesOutputView, *utils.RestErr) {
	r, err := c.repo.GetAll(ctx, query)

	if err != nil {
		return nil, utils.NewInternalServerError(nil)
	}

	return views.NewChallengesCategoriesResponseListOutputView(r), nil
}

// GetById implements ChallengesCategoriesServiceInterface.
func (c *ChallengesCategoriesService) GetById(ctx context.Context, id string) (views.ChallengesCategoriesOutputView, *utils.RestErr) {
	r, err := c.repo.GetById(ctx, id)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return views.ChallengesCategoriesOutputView{}, utils.NewBadRequestException(fmt.Sprintf("THE RECORD WITH ID %s, NOT EXISTS.", id))
		}
		return views.ChallengesCategoriesOutputView{}, utils.NewInternalServerError(nil)
	}

	return views.NewChallengeCategoriesResponseOutputView(r), nil
}

// Update implements ChallengesCategoriesServiceInterface.
func (c *ChallengesCategoriesService) Update(ctx context.Context, id string, title string, userId string) *utils.RestErr {

	r, restErr := c.GetById(ctx, id)

	if restErr != nil {

		return restErr
	}

	if r.UserId != userId {
		return utils.NewUnauthorizedException("ONLY THE USER WHO CREATED THE CATEGORY CAN CHANGE IT.")
	}

	err := c.repo.Update(ctx, id, title)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return utils.NewBadRequestException(fmt.Sprintf("THE RECORD WITH ID %s, NOT EXISTS.", id))
		}

		return utils.NewInternalServerError(nil)
	}

	return nil
}
