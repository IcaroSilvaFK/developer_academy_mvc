package services

import (
	"github.com/IcaroSilvaFK/developer_academy_mvc/application/adapters"
	"github.com/IcaroSilvaFK/developer_academy_mvc/application/utils"
	"github.com/IcaroSilvaFK/developer_academy_mvc/infra/models"
	"github.com/IcaroSilvaFK/developer_academy_mvc/infra/repositories"
	"gorm.io/gorm"
)

type LoginService struct {
	ur     repositories.UserRepositoryInterface
	client utils.HttpClientInterface
}

type LoginServiceInterface interface {
	Login(code, provider string) (*models.UserModel, *utils.RestErr)
}

func NewAuthService(
	ur repositories.UserRepositoryInterface,
	client utils.HttpClientInterface,
) LoginServiceInterface {

	return &LoginService{
		ur, client,
	}
}

func (a *LoginService) Login(code, provider string) (*models.UserModel, *utils.RestErr) {

	prov := a.instaceProvider(provider)

	u, err := prov.SignIn(code)

	if u.Email == "" {
		return nil, utils.NewForbiddenException("Fail on request access user in auth")
	}

	if err != nil {
		return nil, utils.NewForbiddenException("Fail on request access in user account.")
	}

	uExists, err := a.ur.FindByEmail(u.Email)

	if err != nil && err != gorm.ErrRecordNotFound {
		message := "Error on find user"
		return nil, utils.NewInternalServerError(&message)
	}
	if err == gorm.ErrRecordNotFound {
		uExists = models.NewUserModel(
			u.Email, u.Name, u.AvatarUrl, u.Url, u.Bio,
		)
		err = a.ur.Create(uExists)

		if err != nil {
			message := "Error on create user"
			return nil, utils.NewInternalServerError(&message)
		}
	}

	return uExists, nil
}

func (s *LoginService) instaceProvider(provider string) adapters.AdapterAuthInterface {

	httpClient := utils.NewHttpClient()

	switch provider {
	case "gitlab":
		{
			utils.Info("gitlab login")
			return adapters.NewGitlabAdapter(httpClient)
		}
	default:
		{
			utils.Info("github login")
			return adapters.NewGithubAdapter(httpClient)
		}
	}
}
