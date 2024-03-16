package models

import (
	"github.com/IcaroSilvaFK/developer_academy_mvc/infra/utils"
	"gorm.io/gorm"
)

type UserModel struct {
	ID         string
	Name       string
	AvatarUrl  string
	Email      string `gorm:"unique;not null"`
	Bio        string
	Url        string
	Challanges []ChallengeModel `gorm:"foreignKey:UserId;references:ID"`
	gorm.Model
}

func NewUserModel(
	email, name, avatarUrl, url, bio string,
) *UserModel {
	return &UserModel{
		ID:        utils.NewId(),
		Email:     email,
		Name:      name,
		AvatarUrl: avatarUrl,
		Url:       url,
		Bio:       bio,
	}
}

func (u *UserModel) TableName() string {
	return "users"
}
