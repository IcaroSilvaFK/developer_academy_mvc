package repositories

import (
	"context"
	"fmt"

	"github.com/IcaroSilvaFK/developer_academy_mvc/infra/models"
	"gorm.io/gorm"
)

type ChallengeCategoriesUsersRepository struct {
	db *gorm.DB
}

type ChallengeCategoriesUsersRepositoryInterface interface {
	Create(ctx context.Context, input *models.ChallengeCategoriesUsersModel) error
	FindByUserId(ctx context.Context, userId string) ([]*models.ChallengeCategoriesUsersModel, error)
	Delete(ctx context.Context, relationId string) error
}

func NewChallengeCategoriesUser(
	db *gorm.DB,
) ChallengeCategoriesUsersRepositoryInterface {
	return &ChallengeCategoriesUsersRepository{
		db,
	}
}

// Create implements ChallengeCategoriesUsersInterface.
func (c *ChallengeCategoriesUsersRepository) Create(ctx context.Context, input *models.ChallengeCategoriesUsersModel) error {
	return c.makeQuery(ctx).Create(input).Error
}

// Delete implements ChallengeCategoriesUsersInterface.
func (c *ChallengeCategoriesUsersRepository) Delete(ctx context.Context, relationId string) error {
	return c.makeQuery(ctx).Delete(relationId).Error
}

// FindByUserId implements ChallengeCategoriesUsersInterface.
func (c *ChallengeCategoriesUsersRepository) FindByUserId(ctx context.Context, userId string) ([]*models.ChallengeCategoriesUsersModel, error) {
	var r []*models.ChallengeCategoriesUsersModel

	c.makeQuery(ctx).Joins("left join challenges_categories_model as cat on challenge_categories_users_models.category_id = cat.id").Find(r, "userId = ?", userId)

	fmt.Println(r)

	if err := c.makeQuery(ctx).Find(r, "userId = ?", userId).Error; err != nil {

		return nil, err
	}

	return r, nil
}

func (c *ChallengeCategoriesUsersRepository) makeQuery(ctx context.Context) *gorm.DB {

	return c.db.WithContext(ctx).Model(&models.ChallengeCategoriesUsersModel{})
}
