package services

import (
	"os"

	"github.com/IcaroSilvaFK/developer_academy_mvc/application/dtos"
	"github.com/IcaroSilvaFK/developer_academy_mvc/application/utils"
	"github.com/IcaroSilvaFK/developer_academy_mvc/infra/models"
	"github.com/IcaroSilvaFK/developer_academy_mvc/infra/repositories"
)

type AuthService struct {
	ur     repositories.UserRepositoryInterface
	client utils.HttpClientInterface
}

type AuthServiceInterface interface {
	Login(code string) (*models.UserModel, error)
}

func NewAuthService(
	ur repositories.UserRepositoryInterface,
	client utils.HttpClientInterface,
) AuthServiceInterface {

	return &AuthService{
		ur, client,
	}
}

func (a *AuthService) Login(code string) (*models.UserModel, error) {

	var res dtos.GithubTokenResponse

	_, err := a.client.Post("https://github.com/login/oauth/access_token", map[string]string{
		"client_id":     os.Getenv(utils.GITHUB_CLIENT_ID),
		"client_secret": os.Getenv(utils.GITHUB_CLIENT_SECRET),
		"code":          code,
	}, &res)

	if err != nil {
		return nil, err
	}

	var u dtos.GithubResponse

	_, err = a.client.Get("https://api.github.com/user", &u, map[string]string{
		"Authorization": "token " + res.AccessToken,
	})

	if err != nil {
		return nil, err
	}

	return nil, nil
}
