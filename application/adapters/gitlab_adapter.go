package adapters

import (
	"context"
	"os"

	"github.com/IcaroSilvaFK/developer_academy_mvc/application/dtos"
	"github.com/IcaroSilvaFK/developer_academy_mvc/application/utils"
)

type AdapterAuthInterface interface {
	SignIn(ctx context.Context, code string) (dtos.GithubResponse, *utils.RestErr)
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

func (ga *GitlabAuthAdapter) SignIn(ctx context.Context, code string) (dtos.GithubResponse, *utils.RestErr) {

	if code == "" {
		return dtos.GithubResponse{}, utils.NewBadRequestException(utils.CODE_AUTHETICATION_MISSGIN)
	}

	var res dtos.GithubTokenResponse

	_, err := ga.hc.WithContext(ctx).Post(utils.GITLAB_TOKEN_URL, map[string]string{
		"client_id":     os.Getenv(utils.GITLAB_APP_ID),
		"client_secret": os.Getenv(utils.GITLAB_SECRET),
		"code":          code,
		"redirect_uri":  os.Getenv(utils.GITLAB_REDIRECT_URI),
		"grant_type":    "authorization_code",
	}, &res)

	if err != nil {

		return dtos.GithubResponse{}, utils.NewBadRequestException(err.Error())
	}

	var u dtos.GithubResponse

	_, err = ga.hc.WithContext(ctx).Get(utils.GITLAB_USER_API, &u, map[string]string{
		"Authorization": "Bearer " + res.AccessToken,
	})

	if err != nil {
		return dtos.GithubResponse{}, utils.NewBadRequestException(err.Error())
	}

	return u, nil
}
