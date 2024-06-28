package services

import (
	"context"
	"fmt"

	"github.com/IcaroSilvaFK/developer_academy_mvc/application/http/views"
	"github.com/IcaroSilvaFK/developer_academy_mvc/application/utils"
	"github.com/IcaroSilvaFK/developer_academy_mvc/infra/models"
	"github.com/IcaroSilvaFK/developer_academy_mvc/infra/repositories"
	"gorm.io/gorm"
)

type ChallengeCategoriesUsersService struct {
	repo repositories.ChallengeCategoriesUsersRepositoryInterface
}

type ChallengeCategoriesUsersServiceInterface interface {
	Create(ctx context.Context, input views.ChallengeCategoriesUsersInputView, userId string) *utils.RestErr
	FindByUserId(ctx context.Context, userId string) ([]*models.ChallengeCategoriesUsersModel, *utils.RestErr)
	Delete(ctx context.Context, relationId string) *utils.RestErr
}

func NewChallengeCategoriesUsersService(
	repo repositories.ChallengeCategoriesUsersRepositoryInterface,
) ChallengeCategoriesUsersServiceInterface {

	return &ChallengeCategoriesUsersService{repo}
}

// Create implements ChallengeCategoriesUsersServiceInterface.
func (c *ChallengeCategoriesUsersService) Create(ctx context.Context, input views.ChallengeCategoriesUsersInputView, userId string) *utils.RestErr {

	inp := models.NewChallengeCategoriesUsersModel(input.CategoryId, userId)

	if err := c.repo.Create(ctx, inp); err != nil {

		if err == gorm.ErrDuplicatedKey {
			errno := utils.NewBadRequestException(fmt.Sprintf("USER WITH ID %s ALREADY HAS THIS CATEGORY %s FAVORITED", userId, inp.CategoryId))

			return errno
		}

		errno := utils.NewInternalServerError(nil)
		utils.Error("Error on insert data in database", err)

		return errno
	}

	return nil

}

// Delete implements ChallengeCategoriesUsersServiceInterface.
func (c *ChallengeCategoriesUsersService) Delete(ctx context.Context, relationId string) *utils.RestErr {
	if err := c.repo.Delete(ctx, relationId); err != nil {
		if err == gorm.ErrRecordNotFound {
			errno := utils.NewBadRequestException(fmt.Sprintf("THE RELATION WITH ID %s NOT EXISTS", relationId))

			return errno
		}

		errno := utils.NewInternalServerError(nil)

		return errno
	}

	return nil
}

// FindByUserId implements ChallengeCategoriesUsersServiceInterface.
func (c *ChallengeCategoriesUsersService) FindByUserId(ctx context.Context, userId string) ([]*models.ChallengeCategoriesUsersModel, *utils.RestErr) {

	r, err := c.repo.FindByUserId(ctx, userId)

	if err != nil {

		errno := utils.NewInternalServerError(nil)
		return nil, errno
	}

	return r, nil
}
