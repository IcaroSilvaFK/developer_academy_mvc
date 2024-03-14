package services

import (
	"github.com/IcaroSilvaFK/developer_academy_mvc/infra/models"
	"github.com/IcaroSilvaFK/developer_academy_mvc/infra/repositories"
)

type UserService struct {
	ur repositories.UserRepositoryInterface
}

type UserServiceInterface interface {
	GetTenFirstUserAndCount() ([]*models.UserModel, int, error)
}

func NewUserService(
	ur repositories.UserRepositoryInterface,
) UserServiceInterface {

	return &UserService{
		ur,
	}
}

func (us *UserService) GetTenFirstUserAndCount() ([]*models.UserModel, int, error) {
	return us.ur.FindFirstTenAndCount()
}
