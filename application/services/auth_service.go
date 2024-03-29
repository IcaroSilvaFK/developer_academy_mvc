package services

import (
	"os"

	"github.com/IcaroSilvaFK/developer_academy_mvc/application/dtos"
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
	Login(code string) (*models.UserModel, *utils.RestErr)
}

func NewAuthService(
	ur repositories.UserRepositoryInterface,
	client utils.HttpClientInterface,
) LoginServiceInterface {

	return &LoginService{
		ur, client,
	}
}

func (a *LoginService) Login(code string) (*models.UserModel, *utils.RestErr) {

	var res dtos.GithubTokenResponse

	_, err := a.client.Post("https://github.com/login/oauth/access_token", map[string]string{
		"client_id":     os.Getenv(utils.GITHUB_CLIENT_ID),
		"client_secret": os.Getenv(utils.GITHUB_CLIENT_SECRET),
		"code":          code,
	}, &res)

	if err != nil {
		return nil, utils.NewForbiddenException("Fail on request access token on github.")
	}

	var u dtos.GithubResponse

	_, err = a.client.Get("https://api.github.com/user", &u, map[string]string{
		"Authorization": "token " + res.AccessToken,
	})

	if err != nil {
		return nil, utils.NewForbiddenException("Fail on get user details in github.")
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
