package models

import (
	"github.com/IcaroSilvaFK/developer_academy_mvc/infra/utils"
	"gorm.io/gorm"
)

type ChallengeCommentModel struct {
	ID          string
	Comment     string
	UserId      string
	ChallengeId string
	gorm.Model
}

func NewChallengeCommentModel(
	challengeId, userId, comment string,
) *ChallengeCommentModel {

	return &ChallengeCommentModel{
		ID:          utils.NewId(),
		Comment:     comment,
		UserId:      userId,
		ChallengeId: challengeId,
	}
}

func (cm *ChallengeCommentModel) TableName() string {
	return "challenge_comment"
}
