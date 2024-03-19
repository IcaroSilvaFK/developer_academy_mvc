package models

import "gorm.io/gorm"

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

func (c *ChallengeModel) TableName() string {
	return "challenges"
}
