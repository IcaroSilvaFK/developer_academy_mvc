package services

import (
	"context"

	"github.com/IcaroSilvaFK/developer_academy_mvc/application/adapters"
	"github.com/IcaroSilvaFK/developer_academy_mvc/application/utils"
	"github.com/IcaroSilvaFK/developer_academy_mvc/infra/models"
	"github.com/IcaroSilvaFK/developer_academy_mvc/infra/repositories"
	"gorm.io/gorm"
)

type LoginService struct {
	ur            repositories.UserRepositoryInterface
	githubAdapter adapters.AdapterAuthInterface
	gitlabAdapter adapters.AdapterAuthInterface
}

type LoginServiceInterface interface {
	Login(ctx context.Context, code, provider string) (*models.UserModel, *utils.RestErr)
}

func NewAuthService(
	ur repositories.UserRepositoryInterface,
	githubAdapter adapters.AdapterAuthInterface,
	gitlabAdapter adapters.AdapterAuthInterface,
) LoginServiceInterface {

	return &LoginService{
		ur, githubAdapter, gitlabAdapter,
	}
}

func (a *LoginService) Login(ctx context.Context, code, provider string) (*models.UserModel, *utils.RestErr) {

	prov := a.getProvider(provider)

	u, restErr := prov.SignIn(ctx, code)

	if u.Email == "" {
		return nil, utils.NewForbiddenException("Fail on request access user in auth")
	}

	if restErr != nil {
		return nil, restErr
	}

	uExists, err := a.ur.FindByEmail(ctx, u.Email)

	if err != nil && err != gorm.ErrRecordNotFound {
		message := "Error on find user"
		return nil, utils.NewInternalServerError(&message)
	}
	if err == gorm.ErrRecordNotFound {
		uExists = models.NewUserModel(
			u.Email, u.Name, u.AvatarUrl, u.Url, u.Bio,
		)
		err = a.ur.Create(ctx, uExists)

		if err != nil {
			message := "Error on create user"
			return nil, utils.NewInternalServerError(&message)
		}
	}

	return uExists, nil
}

func (s *LoginService) getProvider(provider string) adapters.AdapterAuthInterface {
	switch provider {
	case "gitlab":
		{
			utils.Info("gitlab login")
			return s.gitlabAdapter
		}
	default:
		{
			utils.Info("github login")
			return s.githubAdapter
		}
	}
}
