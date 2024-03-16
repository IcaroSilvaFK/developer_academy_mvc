package models

import "gorm.io/gorm"

type ChallengeModel struct {
	ID          string
	Title       string
	Description string
	EmbedUrl    string
	UserId      string
	gorm.Model
}

func (c *ChallengeModel) TableName() string {
	return "challenges"
}
