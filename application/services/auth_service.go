package services

import (
	"context"
	"time"

	"github.com/IcaroSilvaFK/developer_academy_mvc/application/adapters"
	"github.com/IcaroSilvaFK/developer_academy_mvc/application/dtos"
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
	LoginWithPassword(ctx context.Context, dto *dtos.LoginInputDto) (*models.UserModel, *utils.RestErr)
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
			u.Email, u.Name, u.AvatarUrl, u.Url, u.Bio, nil,
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
	case "github":
		{
			utils.Info("github login")
			return s.githubAdapter
		}
	default:
		{
			utils.Info("github login")
			return s.githubAdapter
		}
	}
}

func (a *LoginService) LoginWithPassword(ctx context.Context, dto *dtos.LoginInputDto) (*models.UserModel, *utils.RestErr) {
	user, err := a.ur.FindByEmail(ctx, dto.Email)

	if err != nil {
		message := "Error on find user"
		return nil, utils.NewInternalServerError(&message)
	}

	if !user.VerifyPassword(dto.Password) {
		return nil, utils.NewBadRequestException("Email or password invalid")
	}

	// this is for brute force atack
	time.Sleep(2 * time.Second)

	return user, nil
}
