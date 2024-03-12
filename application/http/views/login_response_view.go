package views

import (
	"time"

	"github.com/IcaroSilvaFK/developer_academy_mvc/infra/models"
)

type LoginResponseView struct {
	CreateAt  time.Time `json:"create_at"`
	ID        string    `json:"id"`
	UpdatedAt time.Time `json:"updated_at"`
	AvatarUrl string    `json:"avatar_url"`
	Bio       string    `json:"bio"`
	Email     string    `json:"email"`
	Name      string    `json:"name"`
	Url       string    `json:"url"`
}

func NewLoginResponseView(u *models.UserModel) *LoginResponseView {

	return &LoginResponseView{
		AvatarUrl: u.AvatarUrl,
		Bio:       u.Bio,
		Email:     u.Email,
		Name:      u.Name,
		Url:       u.Url,
		ID:        u.ID,
		CreateAt:  u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}
