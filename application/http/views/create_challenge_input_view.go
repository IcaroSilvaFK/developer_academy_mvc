package views

type CreateChallengeInputView struct {
	Title       string   `json:"title" biding:"required"`
	Description string   `json:"description" biding:"required"`
	EmbedUrl    string   `json:"embed_url" biding:"required"`
	Rating      int      `json:"rating" biding:"required"`
	Categories  []string `json:"categories" biding:"required"`
}
