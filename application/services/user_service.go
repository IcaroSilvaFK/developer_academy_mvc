package services

import (
	"fmt"

	"github.com/IcaroSilvaFK/developer_academy_mvc/application/utils"
	"github.com/IcaroSilvaFK/developer_academy_mvc/infra/models"
	"github.com/IcaroSilvaFK/developer_academy_mvc/infra/repositories"
	"gorm.io/gorm"
)

type UserService struct {
	ur repositories.UserRepositoryInterface
}

type UserServiceInterface interface {
	GetTenFirstUserAndCount() ([]*models.UserModel, int, *utils.RestErr)
	FindAllUsers() ([]*models.UserModel, *utils.RestErr)
	FindUserById(string) (*models.UserModel, *utils.RestErr)
	Delete(string) *utils.RestErr
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

func (us *UserService) FindAllUsers() ([]*models.UserModel, *utils.RestErr) {

	users, err := us.ur.FindAll()

	if err != nil {
		message := "Error on find users"
		return nil, utils.NewInternalServerError(&message)
	}

	return users, nil
}
func (us *UserService) FindUserById(id string) (*models.UserModel, *utils.RestErr) {

	u, err := us.ur.FindById(id)

	if err == gorm.ErrRecordNotFound {

		return nil, utils.NewNotFoundException("User not exists")
	}

	if err != nil {
		message := "Error on find user by id"
		return nil, utils.NewInternalServerError(&message)
	}

	return u, nil
}

func (us *UserService) Delete(id string) *utils.RestErr {

	err := us.ur.Delete(id)

	if err == gorm.ErrRecordNotFound {

		return utils.NewNotFoundException(fmt.Sprintf("User id %s not exists", id))
	}

	if err != nil {
		message := fmt.Sprintf("Error on delete user id %s", id)
		return utils.NewInternalServerError(&message)

	}

	return nil
}
