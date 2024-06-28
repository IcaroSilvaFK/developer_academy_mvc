package views

type ChallengeCategoriesUsersOutputView struct {
	Id         string `json:"id"`
	Category   string `json:"category"`
	CategoryId string `json:"categoryId"`
	UserId     string `json:"userId"`
}
