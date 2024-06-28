package services

import (
	"context"

	"github.com/IcaroSilvaFK/developer_academy_mvc/application/http/views"
	"github.com/IcaroSilvaFK/developer_academy_mvc/application/utils"
	"github.com/IcaroSilvaFK/developer_academy_mvc/infra/models"
	"github.com/IcaroSilvaFK/developer_academy_mvc/infra/repositories"
)

type ChallengeCategoriesUsersService struct {
	repo repositories.ChallengeCategoriesUsersRepositoryInterface
}

type ChallengeCategoriesUsersServiceInterface interface {
	Create(ctx context.Context, input views.ChallengeCategoriesUsersInputView, userId string) *utils.RestErr
	FindByUserId(ctx context.Context, userId string) ([]*models.ChallengeCategoriesUsersModel, *utils.RestErr)
	Delete(ctx context.Context, relationId string) *utils.RestErr
}
