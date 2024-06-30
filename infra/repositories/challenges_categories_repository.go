package repositories

import (
	"context"

	"github.com/IcaroSilvaFK/developer_academy_mvc/application/utils"
	"github.com/IcaroSilvaFK/developer_academy_mvc/infra/models"
	"gorm.io/gorm"
)

type ChallengesCategoriesRepository struct {
	db        *gorm.DB
	tableName string
}

type ChallengesCategoriesRepositoryInterface interface {
	Create(ctx context.Context, userId string, title string) (*models.ChallengesCategoriesModel, error)
	GetAll(ctx context.Context, query string) ([]*models.ChallengesCategoriesModel, error)
	GetById(ctx context.Context, id string) (*models.ChallengesCategoriesModel, error)
	Update(ctx context.Context, id string, title string) error
	FindByUserId(ctx context.Context, userId string) ([]*models.ChallengesCategoriesModel, error)
	Delete(ctx context.Context, id string) error
}

func NewChallengesCategoriesRepository(
	db *gorm.DB,
) ChallengesCategoriesRepositoryInterface {

	return &ChallengesCategoriesRepository{db, "challenges_categories_model"}
}

// Create implements ChallengesCategoriesRepositoryInterface.
func (c *ChallengesCategoriesRepository) Create(ctx context.Context, userId string, title string) (*models.ChallengesCategoriesModel, error) {

	cat := models.NewChallengesCategoriesModel(title, userId)

	if err := c.db.WithContext(ctx).Table(c.tableName).Create(&cat).Error; err != nil {
		utils.Error("Error on create category", err)
		return nil, err
	}

	return cat, nil
}

// Delete implements ChallengesCategoriesRepositoryInterface.
func (c *ChallengesCategoriesRepository) Delete(ctx context.Context, id string) error {
	return c.db.WithContext(ctx).Table(c.tableName).Where("id = ?", id).Delete(&models.ChallengesCategoriesModel{
		ID: id,
	}).Error
}

// FindByUserId implements ChallengesCategoriesRepositoryInterface.
func (c *ChallengesCategoriesRepository) FindByUserId(ctx context.Context, userId string) ([]*models.ChallengesCategoriesModel, error) {
	var r []*models.ChallengesCategoriesModel

	if err := c.db.WithContext(ctx).Table(c.tableName).Find(&r, "userId = ?", userId).Error; err != nil {
		return nil, err
	}

	return r, nil
}

// GetAll implements ChallengesCategoriesRepositoryInterface.
func (c *ChallengesCategoriesRepository) GetAll(ctx context.Context, query string) ([]*models.ChallengesCategoriesModel, error) {

	var r []*models.ChallengesCategoriesModel

	if err := c.db.WithContext(ctx).Table(c.tableName).Find(&r).Error; err != nil {
		return nil, err
	}

	return r, nil
}

// GetById implements ChallengesCategoriesRepositoryInterface.
func (c *ChallengesCategoriesRepository) GetById(ctx context.Context, id string) (*models.ChallengesCategoriesModel, error) {
	var r *models.ChallengesCategoriesModel

	if err := c.db.WithContext(ctx).Table(c.tableName).First(&r, id).Error; err != nil {
		return nil, err
	}

	return r, nil
}

// Update implements ChallengesCategoriesRepositoryInterface.
func (c *ChallengesCategoriesRepository) Update(ctx context.Context, id string, title string) error {
	var cat *models.ChallengesCategoriesModel

	if err := c.db.WithContext(ctx).Table(c.tableName).First(&cat, id).Error; err != nil {
		return err
	}

	cat.Title = title

	if err := c.db.Save(cat).Error; err != nil {
		return err
	}

	return nil
}
