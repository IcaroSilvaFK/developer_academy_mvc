package dtos

type GithubTokenResponse struct {
	AccessToken string `json:"access_token"`
}

type GithubResponse struct {
	AvatarUrl   string `json:"avatar_url,omitempty"`
	Bio         string `json:"bio"`
	Email       string `json:"email"`
	Name        string `json:"name"`
	Url         string `json:"url"`
	PublicEmail string `json:"public_email"`
}
