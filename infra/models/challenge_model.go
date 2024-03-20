package models

import (
	"github.com/IcaroSilvaFK/developer_academy_mvc/infra/utils"
	"gorm.io/gorm"
)

type ChallengeModel struct {
	ID          string
	Title       string
	Description string
	EmbedUrl    string
	UserId      string
	Comments    []ChallengeCommentModel `gorm:"foreignKey:ChallengeId;references:ID"`
	Hint        ChallengeHintsModel     `gorm:"foreignKey:ChallengeId;references:ID"`
	gorm.Model
}

func NewChallengeModel(title, description, embedUrl, userId string) *ChallengeModel {

	return &ChallengeModel{
		ID:          utils.NewId(),
		Title:       title,
		Description: description,
		EmbedUrl:    embedUrl,
		UserId:      userId,
	}
}

func (c *ChallengeModel) TableName() string {
	return "challenges"
}
