package models

type ChallengeHintsModel struct {
	ID          string
	Text        string
	ChallengeId string
}

func (ch *ChallengeHintsModel) TableName() string {
	return "challenge_hints"
}
