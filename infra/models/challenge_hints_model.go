package models

import (
	"github.com/IcaroSilvaFK/developer_academy_mvc/infra/utils"
	"gorm.io/gorm"
)

type ChallengeHintsModel struct {
	ID          string
	Text        string
	ChallengeId string
	gorm.Model
}

func NewChallengeHintsModel(challengeId, text string) *ChallengeHintsModel {
	return &ChallengeHintsModel{
		ID:          utils.NewId(),
		Text:        text,
		ChallengeId: challengeId,
	}
}

func (ch *ChallengeHintsModel) TableName() string {
	return "challenge_hints"
}
