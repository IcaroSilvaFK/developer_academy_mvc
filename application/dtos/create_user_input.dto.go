package dtos

type CreateUserInputDto struct {
	Email    string `json:"email" validate:"requried,email"`
	Name     string `json:"username" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}
