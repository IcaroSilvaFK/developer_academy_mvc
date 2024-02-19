package repositories

import (
	"github.com/IcaroSilvaFK/developer_academy_mvc/infra/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

type UserRepositoryInterface interface {
	Create(user *models.UserModel) error
	FindByEmail(email string) (*models.UserModel, error)
	FindById(id string) (*models.UserModel, error)
	FindAll() (*[]models.UserModel, error)
	Delete(id string) error
	Count() (int64, error)
}

func NewUserRepository(db *gorm.DB) UserRepositoryInterface {

	return &UserRepository{
		db,
	}
}

func (ur *UserRepository) Create(user *models.UserModel) error {
	return ur.db.Create(user).Error
}

func (ur *UserRepository) FindByEmail(email string) (*models.UserModel, error) {
	var user models.UserModel
	return &user, ur.db.First(&user, "email = ?", email).Error
}

func (ur *UserRepository) FindById(id string) (*models.UserModel, error) {
	var user models.UserModel
	return &user, ur.db.First(&user, "id = ?", id).Error
}

func (ur *UserRepository) FindAll() (*[]models.UserModel, error) {
	var users []models.UserModel
	return &users, ur.db.Find(&users).Error
}

func (ur *UserRepository) Delete(id string) error {
	return ur.db.Delete(&models.UserModel{}, "id = ?", id).Error
}

func (ur *UserRepository) Count() (int64, error) {

	var c int64

	return c, ur.db.Model(&models.UserModel{}).Count(&c).Error
}
