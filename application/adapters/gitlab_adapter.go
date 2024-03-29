package adapters

import (
	"os"

	"github.com/IcaroSilvaFK/developer_academy_mvc/application/dtos"
	"github.com/IcaroSilvaFK/developer_academy_mvc/application/utils"
)

type AdapterAuthInterface interface {
	SignIn(code string) (dtos.GithubResponse, error)
}

type GitlabAuthAdapter struct {
	hc utils.HttpClientInterface
}

func NewGitlabAdapter(
	hc utils.HttpClientInterface,
) AdapterAuthInterface {
	return &GitlabAuthAdapter{
		hc,
	}
}

func (ga *GitlabAuthAdapter) SignIn(code string) (dtos.GithubResponse, error) {

	var res dtos.GithubTokenResponse

	_, err := ga.hc.Post("https://gitlab.com/oauth/token", map[string]string{
		"client_id":     os.Getenv(utils.GITLAB_APP_ID),
		"client_secret": os.Getenv(utils.GITLAB_SECRET),
		"code":          code,
		"redirect_uri":  os.Getenv(utils.GITLAB_REDIRECT_URI),
		"grant_type":    "authorization_code",
	}, &res)

	if err != nil {

		return dtos.GithubResponse{}, err
	}

	var u dtos.GithubResponse

	_, err = ga.hc.Get("https://gitlab.com/api/v4/user", &u, map[string]string{
		"Authorization": "Bearer " + res.AccessToken,
	})

	if err != nil {
		return dtos.GithubResponse{}, err
	}

	return u, nil
}
