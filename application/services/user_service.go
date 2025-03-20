package services

import (
	"context"
	"fmt"

	"github.com/IcaroSilvaFK/developer_academy_mvc/application/dtos"
	"github.com/IcaroSilvaFK/developer_academy_mvc/application/utils"
	"github.com/IcaroSilvaFK/developer_academy_mvc/infra/models"
	"github.com/IcaroSilvaFK/developer_academy_mvc/infra/repositories"
	"gorm.io/gorm"
)

type UserService struct {
	ur repositories.UserRepositoryInterface
}

type UserServiceInterface interface {
	GetTenFirstUserAndCount(context.Context) ([]*models.UserModel, int, *utils.RestErr)
	FindAllUsers(context.Context) ([]*models.UserModel, *utils.RestErr)
	FindUserById(context.Context, string) (*models.UserModel, *utils.RestErr)
	Delete(context.Context, string) *utils.RestErr
	CreateUser(context.Context, *dtos.CreateUserInputDto) (*models.UserModel, *utils.RestErr)
}

func NewUserService(
	ur repositories.UserRepositoryInterface,
) UserServiceInterface {

	return &UserService{
		ur,
	}
}

func (us *UserService) CreateUser(ctx context.Context, dto *dtos.CreateUserInputDto) (*models.UserModel, *utils.RestErr) {
	u := models.NewUserModel(dto.Email, dto.Name, "https://cataas.com/cat", "", "Hello, start here with your bio!", &dto.Password)

	err := us.ur.Create(ctx, u)

	if err != nil {
		return nil, utils.NewInternalServerError(nil)
	}

	return u, nil
}

func (us *UserService) GetTenFirstUserAndCount(ctx context.Context) ([]*models.UserModel, int, *utils.RestErr) {

	u, c, err := us.ur.FindFirstTenAndCount(ctx)

	if err != nil {
		message := "Error on get first ten users"
		return nil, 0, utils.NewInternalServerError(&message)
	}

	return u, c, nil
}

func (us *UserService) FindAllUsers(ctx context.Context) ([]*models.UserModel, *utils.RestErr) {

	users, err := us.ur.FindAll(ctx)

	if err != nil {
		message := "Error on find users"
		return nil, utils.NewInternalServerError(&message)
	}

	return users, nil
}
func (us *UserService) FindUserById(ctx context.Context, id string) (*models.UserModel, *utils.RestErr) {

	u, err := us.ur.FindById(ctx, id)

	if err == gorm.ErrRecordNotFound {

		return nil, utils.NewNotFoundException("User not exists")
	}

	if err != nil {
		message := "Error on find user by id"
		return nil, utils.NewInternalServerError(&message)
	}

	return u, nil
}

func (us *UserService) Delete(ctx context.Context, id string) *utils.RestErr {

	err := us.ur.Delete(ctx, id)

	if err == gorm.ErrRecordNotFound {

		return utils.NewNotFoundException(fmt.Sprintf("User id %s not exists", id))
	}

	if err != nil {
		message := fmt.Sprintf("Error on delete user id %s", id)
		return utils.NewInternalServerError(&message)
	}

	return nil
}
