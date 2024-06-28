package models

import (
	"github.com/IcaroSilvaFK/developer_academy_mvc/infra/utils"
	"gorm.io/gorm"
)

type ChallengeCategoriesUsersModel struct {
	ID         string
	UserId     string `gorm:"index:idx_catmember"`
	CategoryId string `gorm:"index:idx_catmember"`
	ChallengesCategoriesModel

	gorm.Model
}

func NewChallengeCategoriesUsersModel(
	catId, userId string,
) *ChallengeCategoriesUsersModel {

	return &ChallengeCategoriesUsersModel{
		ID:         utils.NewId(),
		CategoryId: catId,
		UserId:     userId,
	}
}
