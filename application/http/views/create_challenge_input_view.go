package views

type CreateChallengeInputView struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	EmbedUrl    string `json:"embed_url" validate:"required"`
}

