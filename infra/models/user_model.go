package models

import (
	"github.com/IcaroSilvaFK/developer_academy_mvc/infra/utils"
	"gorm.io/gorm"
)

type UserModel struct {
	ID        string `json:"id" gorm:"primarykey"`
	Name      string `json:"name"`
	AvatarUrl string `json:"avatar_url"`
	Email     string `json:"email" gorm:"index"`
	Blog      string `json:"blog"`
	Url       string `json:"url"`
	gorm.Model
}

func NewUserModel(
	email, name, avatarUrl, url, blog string,
) *UserModel {
	return &UserModel{
		ID:        utils.NewId(),
		Email:     email,
		Name:      name,
		AvatarUrl: avatarUrl,
		Url:       url,
	}
}

func (u *UserModel) TableName() string {
	return "users"
}
