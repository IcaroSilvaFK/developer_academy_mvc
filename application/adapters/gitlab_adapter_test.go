package adapters_test

import (
	"context"
	"testing"

	"github.com/IcaroSilvaFK/developer_academy_mvc/application/adapters"
	"github.com/IcaroSilvaFK/developer_academy_mvc/application/dtos"
	"github.com/IcaroSilvaFK/developer_academy_mvc/application/utils"
	"github.com/stretchr/testify/assert"
)

func TestShouldSignInWithGitLab(t *testing.T) {

	clientMock := &mockClient{
		CalledUrls: []string{},
	}

	adapter := adapters.NewGitlabAdapter(clientMock)

	res, err := adapter.SignIn(context.Background(), "123")

	expected := dtos.GithubResponse{
		Name:        "name",
		AvatarUrl:   "avatar",
		Bio:         "bio",
		Email:       "email",
		Url:         "url",
		PublicEmail: "email",
	}

	assert.Nil(t, err)
	assert.Equal(t, expected.Name, res.Name)
	assert.Equal(t, expected.Email, res.Email)
	assert.Equal(t, expected.PublicEmail, res.PublicEmail)
	assert.Equal(t, expected.Bio, res.Bio)
	assert.Equal(t, expected.AvatarUrl, res.AvatarUrl)
	assert.Equal(t, clientMock.CalledUrls[0], "https://gitlab.com/oauth/token")
	assert.Equal(t, clientMock.CalledUrls[1], "https://gitlab.com/api/v4/user")
}

func TestShouldReturnErrorWhenParamsIsMissing(t *testing.T) {
	clientMock := &mockClient{
		CalledUrls: []string{},
	}
	adapter := adapters.NewGithubAdapter(clientMock)

	_, err := adapter.SignIn(context.Background(), "")

	assert.Equal(t, utils.NewBadRequestException(utils.CODE_AUTHETICATION_MISSGIN), err)
	assert.Equal(t, len(clientMock.CalledUrls), 0)
}
