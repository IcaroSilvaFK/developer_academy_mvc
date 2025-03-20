package models

import (
	"github.com/IcaroSilvaFK/developer_academy_mvc/infra/utils"

	"gorm.io/gorm"
)

type UserModel struct {
	ID         string
	Name       string
	Password   *string
	AvatarUrl  string
	Email      string `gorm:"unique;not null"`
	Bio        string
	Url        string
	Challanges []ChallengeModel        `gorm:"foreignKey:UserId;references:ID"`
	Comments   []ChallengeCommentModel `gorm:"foreignKey:UserId;references:ID"`
	gorm.Model
}

func NewUserModel(
	email, name, avatarUrl, url, bio string, password *string,
) *UserModel {
	return &UserModel{
		ID:        utils.NewId(),
		Email:     email,
		Name:      name,
		AvatarUrl: avatarUrl,
		Url:       url,
		Bio:       bio,
		Password:  password,
	}
}

func (u *UserModel) HashPassword() {
	result := utils.MakeHash(*u.Password)
	u.Password = &result
}

func (u *UserModel) VerifyPassword(password string) bool {
	return utils.VerifyHash(password, *u.Password)
}

func (u *UserModel) TableName() string {
	return "users"
}
