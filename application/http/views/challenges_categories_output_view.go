package views

import (
	"fmt"

	"github.com/IcaroSilvaFK/developer_academy_mvc/application/utils"
	"github.com/IcaroSilvaFK/developer_academy_mvc/infra/models"
)

type ChallengesCategoriesOutputView struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	CreatedAt string `json:"createdAt"`
	UserId    string `json:"userId"`
}

func NewChallengesCategoriesResponseListOutputView(ccs []*models.ChallengesCategoriesModel) []ChallengesCategoriesOutputView {

	var r []ChallengesCategoriesOutputView = []ChallengesCategoriesOutputView{}

	for _, v := range ccs {
		r = append(r, NewChallengeCategoriesResponseOutputView(v))
	}

	return r
}

func NewChallengeCategoriesResponseOutputView(cc *models.ChallengesCategoriesModel) ChallengesCategoriesOutputView {

	fmt.Println(cc)

	return ChallengesCategoriesOutputView{
		ID:        cc.ID,
		Title:     cc.Title,
		CreatedAt: utils.ConvertTimeToText(cc.CreatedAt),
		UserId:    cc.UserId,
	}
}
