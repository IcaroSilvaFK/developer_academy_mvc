package adapters

import (
	"os"

	"github.com/IcaroSilvaFK/developer_academy_mvc/application/dtos"
	"github.com/IcaroSilvaFK/developer_academy_mvc/application/utils"
)

type GithubAuthAdpter struct {
	hc utils.HttpClientInterface
}

func NewGithubAdapter(
	hc utils.HttpClientInterface,
) AdapterAuthInterface {

	return &GithubAuthAdpter{
		hc,
	}
}

func (aa *GithubAuthAdpter) SignIn(code string) (dtos.GithubResponse, error) {
	var res dtos.GithubTokenResponse

	_, err := aa.hc.Post("https://github.com/login/oauth/access_token", map[string]string{
		"client_id":     os.Getenv(utils.GITHUB_CLIENT_ID),
		"client_secret": os.Getenv(utils.GITHUB_CLIENT_SECRET),
		"code":          code,
	}, &res)

	if err != nil {
		return dtos.GithubResponse{}, err
	}

	var u dtos.GithubResponse

	_, err = aa.hc.Get("https://api.github.com/user", &u, map[string]string{
		"Authorization": "token " + res.AccessToken,
	})

	return u, nil
}
