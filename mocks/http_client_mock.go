package mocks_test

import (
	"encoding/json"
	"net/http"

	"github.com/IcaroSilvaFK/developer_academy_mvc/application/dtos"
	"github.com/google/uuid"
)

type MockClient struct {
	CalledUrls []string
}

func NewMockClientHttp() *MockClient {

	return &MockClient{
		CalledUrls: []string{},
	}
}

// CreateUrl implements utils.HttpClientInterface.
func (m *MockClient) CreateUrl(string, map[string]string) string {
	panic("unimplemented")
}

// Get implements utils.HttpClientInterface.
func (m *MockClient) Get(url string, body interface{}, headers map[string]string) (*http.Response, error) {

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
func (m *MockClient) Post(url string, body interface{}, res interface{}) (*http.Response, error) {

	m.CalledUrls = append(m.CalledUrls, url)

	r := dtos.GithubTokenResponse{
		AccessToken: uuid.NewString(),
	}

	bt, _ := json.Marshal(r)

	json.Unmarshal(bt, res)
	return nil, nil
}
