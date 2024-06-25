package views

type ChallengesCategoriesInputView struct {
	Title  string `json:"title" biding:"required"`
	UserId string
}
