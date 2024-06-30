package models

import (
	"github.com/IcaroSilvaFK/developer_academy_mvc/infra/utils"
	"gorm.io/gorm"
)

type ChallengesCategoriesModel struct {
	Title  string `gorm:"uniqueIndex"`
	ID     string
	UserId string

	gorm.Model
}

func NewChallengesCategoriesModel(
	title string,
	userId string,
) *ChallengesCategoriesModel {

	return &ChallengesCategoriesModel{
		ID:     utils.NewId(),
		Title:  title,
		UserId: userId,
	}
}

func (cc *ChallengesCategoriesModel) TableName() string {
	return "challenges_categories_model"
}
