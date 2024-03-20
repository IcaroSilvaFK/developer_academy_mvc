package views

import (
	"github.com/IcaroSilvaFK/developer_academy_mvc/infra/models"
)

type ResponseChallengeOutputView struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	EmbedUrl    string `json:"embed_url"`
	UserId      string `json:"user_id"`
	CreatedAt   string `json:"created_at"`
}

func NewChallengeResponseOutputList(c []*models.ChallengeModel) []ResponseChallengeOutputView {

	var r []ResponseChallengeOutputView

	for _, challenge := range c {
		r = append(r, NewChallengeResponseOutput(challenge))
	}

	return r
}

func NewChallengeResponseOutput(c *models.ChallengeModel) ResponseChallengeOutputView {

	t, _ := c.CreatedAt.UTC().MarshalText()

	return ResponseChallengeOutputView{
		ID:          c.ID,
		Title:       c.Title,
		Description: c.Description,
		EmbedUrl:    c.EmbedUrl,
		UserId:      c.UserId,
		CreatedAt:   string(t),
	}
}
