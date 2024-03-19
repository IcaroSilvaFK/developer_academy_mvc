package models

import "gorm.io/gorm"

type ChallengeCommentModel struct {
	ID          string
	Comment     string
	UserId      string
	ChallengeId string
	gorm.Model
}

func (cm *ChallengeCommentModel) TableName() string {
	return "challenge_comment"
}
