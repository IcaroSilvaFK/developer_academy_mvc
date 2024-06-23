package repositories

import (
	"context"

	"github.com/IcaroSilvaFK/developer_academy_mvc/infra/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

type UserRepositoryInterface interface {
	Create(ctx context.Context, user *models.UserModel) error
	FindByEmail(ctx context.Context, email string) (*models.UserModel, error)
	FindFirstTenAndCount(ctx context.Context) ([]*models.UserModel, int, error)
	FindById(ctx context.Context, id string) (*models.UserModel, error)
	FindAll(ctx context.Context) ([]*models.UserModel, error)
	Delete(ctx context.Context, id string) error
	Count(ctx context.Context) (int64, error)
}

func NewUserRepository(db *gorm.DB) UserRepositoryInterface {

	return &UserRepository{
		db,
	}
}

func (ur *UserRepository) Create(ctx context.Context, user *models.UserModel) error {
	return ur.db.WithContext(ctx).Create(user).Error
}

func (ur *UserRepository) FindByEmail(ctx context.Context, email string) (*models.UserModel, error) {
	var user models.UserModel
	return &user, ur.db.WithContext(ctx).First(&user, "email = ?", email).Error
}

func (ur *UserRepository) FindById(ctx context.Context, id string) (*models.UserModel, error) {
	var user models.UserModel
	return &user, ur.db.WithContext(ctx).First(&user, "id = ?", id).Error
}

func (ur *UserRepository) FindAll(ctx context.Context) ([]*models.UserModel, error) {
	var users []*models.UserModel
	return users, ur.db.WithContext(ctx).Find(&users).Error
}

func (ur *UserRepository) Delete(ctx context.Context, id string) error {
	return ur.db.WithContext(ctx).Delete(&models.UserModel{}, "id = ?", id).Error
}

func (ur *UserRepository) FindFirstTenAndCount(ctx context.Context) ([]*models.UserModel, int, error) {

	var count int64
	var users []*models.UserModel
	quantityLimitUsers := 10

	tx := ur.db.WithContext(ctx).Begin()

	tx.Model(&models.UserModel{}).Limit(quantityLimitUsers).Find(&users)
	tx.Model(&models.UserModel{}).Count(&count)

	err := tx.Commit().Error

	return users, int(count) - quantityLimitUsers, err
}

func (ur *UserRepository) Count(ctx context.Context) (int64, error) {

	var c int64

	return c, ur.db.WithContext(ctx).Model(&models.UserModel{}).Count(&c).Error
}
