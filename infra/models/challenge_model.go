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
	Comments    []*ChallengeCommentModel     `gorm:"foreignKey:ChallengeId;references:ID;OnDelete:CASCADE;"`
	Hint        ChallengeHintsModel          `gorm:"foreignKey:ChallengeId;references:ID;OnDelete:CASCADE;"`
	Categories  []*ChallengesCategoriesModel `gorm:"many2many:challenges_categories;"`
	gorm.Model
}

func NewChallengeModel(
	title, description, embedUrl, userId string,
	categoriesId []string,
) *ChallengeModel {

	c := &ChallengeModel{
		ID:          utils.NewId(),
		Title:       title,
		Description: description,
		EmbedUrl:    embedUrl,
		UserId:      userId,
	}

	if len(categoriesId) > 0 {
		for _, cat := range categoriesId {
			c.Categories = append(c.Categories, &ChallengesCategoriesModel{
				ID: cat,
			})
		}
	}

	return c
}

func (c *ChallengeModel) BeforeDelete(tx *gorm.DB) (err error) {
	tx.Debug().Clauses(clause.Returning{}).Where("challenge_id = ?", c.ID).Delete(&ChallengeHintsModel{})
	return
}

func (c *ChallengeModel) TableName() string {
	return "challenges"
}
