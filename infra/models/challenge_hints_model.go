package models

import "github.com/IcaroSilvaFK/developer_academy_mvc/infra/utils"

type ChallengeHintsModel struct {
	ID          string
	Text        string
	ChallengeId string
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
