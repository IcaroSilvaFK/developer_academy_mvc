package adapters

import (
	"context"
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

func (aa *GithubAuthAdpter) SignIn(ctx context.Context, code string) (dtos.GithubResponse, *utils.RestErr) {
	if code == "" {
		return dtos.GithubResponse{}, utils.NewBadRequestException(utils.CODE_AUTHETICATION_MISSGIN)
	}

	var res dtos.GithubTokenResponse

	_, err := aa.hc.WithContext(ctx).Post(utils.GITHUB_AUTH_URL, map[string]string{
		"client_id":     os.Getenv(utils.GITHUB_CLIENT_ID),
		"client_secret": os.Getenv(utils.GITHUB_CLIENT_SECRET),
		"code":          code,
	}, &res)

	if err != nil {
		return dtos.GithubResponse{}, utils.NewBadRequestException(err.Error())
	}

	var u dtos.GithubResponse

	_, err = aa.hc.WithContext(ctx).Get(utils.GITHUB_USER_API, &u, map[string]string{
		"Authorization": "token " + res.AccessToken,
	})

	if err != nil {
		return dtos.GithubResponse{}, utils.NewBadRequestException(err.Error())
	}

	return u, nil
}
