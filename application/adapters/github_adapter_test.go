package adapters_test

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/IcaroSilvaFK/developer_academy_mvc/application/adapters"
	"github.com/IcaroSilvaFK/developer_academy_mvc/application/dtos"
	"github.com/IcaroSilvaFK/developer_academy_mvc/application/utils"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

type mockClient struct {
	CalledUrls []string
	ctx        context.Context
}

// CreateUrl implements utils.HttpClientInterface.
func (m *mockClient) CreateUrl(string, map[string]string) string {
	panic("unimplemented")
}

func (m *mockClient) WithContext(ctx context.Context) *utils.HttpClient {

	m.ctx = ctx
	return &utils.HttpClient{}
}

// Get implements utils.HttpClientInterface.
func (m *mockClient) Get(url string, body interface{}, headers map[string]string) (*http.Response, error) {

	m.CalledUrls = append(m.CalledUrls, url)

	bd := dtos.GithubResponse{
		Name:        "name",
		AvatarUrl:   "avatar",
		Bio:         "bio",
		Email:       "email",
		Url:         "url",
		PublicEmail: "email",
	}

	bt, _ := json.Marshal(bd)

	json.Unmarshal(bt, body)

	return nil, nil
}

// Post implements utils.HttpClientInterface.
func (m *mockClient) Post(url string, body interface{}, res interface{}) (*http.Response, error) {

	m.CalledUrls = append(m.CalledUrls, url)

	r := dtos.GithubTokenResponse{
		AccessToken: uuid.NewString(),
	}

	bt, _ := json.Marshal(r)

	json.Unmarshal(bt, res)
	return nil, nil
}

func TestShouldSignInWithGithub(t *testing.T) {
	clientMock := &mockClient{
		CalledUrls: []string{},
	}
	adapter := adapters.NewGithubAdapter(clientMock)

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

	assert.Equal(t, clientMock.CalledUrls[0], "https://github.com/login/oauth/access_token")
	assert.Equal(t, clientMock.CalledUrls[1], "https://api.github.com/user")

}

func TestShouldReturnErrorWhenCodeIsMissing(t *testing.T) {
	clientMock := &mockClient{
		CalledUrls: []string{},
	}
	adapter := adapters.NewGithubAdapter(clientMock)

	_, err := adapter.SignIn(context.Background(), "")

	assert.Equal(t, utils.NewBadRequestException(utils.CODE_AUTHETICATION_MISSGIN), err)
	assert.Equal(t, len(clientMock.CalledUrls), 0)

}
