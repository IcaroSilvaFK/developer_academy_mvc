package models

import (
	"github.com/IcaroSilvaFK/developer_academy_mvc/infra/utils"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ChallengeModel struct {
	ID          string
	Title       string
	Description string
	EmbedUrl    string
	UserId      string
	Comments    []*ChallengeCommentModel `gorm:"foreignKey:ChallengeId;references:ID;OnDelete:CASCADE;"`
	Hint        ChallengeHintsModel      `gorm:"foreignKey:ChallengeId;references:ID;OnDelete:CASCADE;"`
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

func (c *ChallengeModel) BeforeDelete(tx *gorm.DB) (err error) {
	tx.Debug().Clauses(clause.Returning{}).Where("challenge_id = ?", c.ID).Delete(&ChallengeHintsModel{})
	return
}

func (c *ChallengeModel) TableName() string {
	return "challenges"
}
