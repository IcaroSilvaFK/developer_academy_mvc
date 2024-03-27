package services

import (
	"github.com/IcaroSilvaFK/developer_academy_mvc/application/utils"
	"github.com/IcaroSilvaFK/developer_academy_mvc/infra/models"
	"github.com/IcaroSilvaFK/developer_academy_mvc/infra/repositories"
)

type UserService struct {
	ur repositories.UserRepositoryInterface
}

type UserServiceInterface interface {
	GetTenFirstUserAndCount() ([]*models.UserModel, int, *utils.RestErr)
}

func NewUserService(
	ur repositories.UserRepositoryInterface,
) UserServiceInterface {

	return &UserService{
		ur,
	}
}

func (us *UserService) GetTenFirstUserAndCount() ([]*models.UserModel, int, *utils.RestErr) {

	m, i, err := us.ur.FindFirstTenAndCount()

	if err != nil {
		message := "Error on get first ten users"
		return nil, 0, utils.NewInternalServerError(&message)
	}

	return m, i, nil
}
