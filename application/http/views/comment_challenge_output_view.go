package views

import "github.com/IcaroSilvaFK/developer_academy_mvc/infra/models"

type CommentChallengeOutputView struct {
	ID          string `json:"id"`
	Comment     string `json:"comment"`
	UserId      string `json:"user_id"`
	ChallengeId string `json:"challenge_id"`
	CreatedAt   string `json:"created_at"`
}

func NewCommentChallengeOutputView(c *models.ChallengeCommentModel) CommentChallengeOutputView {

	created, _ := c.CreatedAt.UTC().MarshalText()

	return CommentChallengeOutputView{
		ID:          c.ID,
		Comment:     c.Comment,
		UserId:      c.UserId,
		ChallengeId: c.ChallengeId,
		CreatedAt:   string(created),
	}
}

func NewCommentChallengeListOutputView(list []models.ChallengeCommentModel) []CommentChallengeOutputView {

	r := make([]CommentChallengeOutputView, 0)

	for _, comment := range list {
		r = append(r, NewCommentChallengeOutputView(&comment))
	}

	return r
}
