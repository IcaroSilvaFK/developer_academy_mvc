package repositories

import (
	"context"

	"github.com/IcaroSilvaFK/developer_academy_mvc/infra/models"
	"gorm.io/gorm"
)

type CommentChallengeRepository struct {
	db *gorm.DB
}

type CommentChallengeRepositoryInterface interface {
	Create(ctx context.Context, m *models.ChallengeCommentModel) error
	FindById(ctx context.Context, commentId string) (*models.ChallengeCommentModel, error)
	FindByUserId(ctx context.Context, userId string) ([]*models.ChallengeCommentModel, error)
	FindByChallengeId(ctx context.Context, challengeId string) ([]*models.ChallengeCommentModel, error)
	Delete(ctx context.Context, commentId string) error
}

func NewCommentChallengeRepository(
	db *gorm.DB,
) CommentChallengeRepositoryInterface {

	return &CommentChallengeRepository{db}
}

func (cc *CommentChallengeRepository) Create(ctx context.Context, c *models.ChallengeCommentModel) error {

	err := cc.db.WithContext(ctx).Model(&models.ChallengeCommentModel{}).Create(c).Error

	return err
}
func (cc *CommentChallengeRepository) FindById(ctx context.Context, commentId string) (*models.ChallengeCommentModel, error) {

	var c models.ChallengeCommentModel

	err := cc.db.WithContext(ctx).Model(&models.ChallengeCommentModel{}).Find(&c, "id = ?", commentId).Error

	if err != nil {
		return nil, err
	}
	return &c, nil
}

func (cc *CommentChallengeRepository) FindByUserId(ctx context.Context, userId string) ([]*models.ChallengeCommentModel, error) {
	var c []*models.ChallengeCommentModel

	err := cc.db.WithContext(ctx).Model(&models.ChallengeCommentModel{}).Find(&c, "user_id = ?", userId).Error
	if err != nil {
		return nil, err
	}

	return c, nil
}
func (cc *CommentChallengeRepository) FindByChallengeId(ctx context.Context, challengeId string) ([]*models.ChallengeCommentModel, error) {

	var c []*models.ChallengeCommentModel

	err := cc.db.WithContext(ctx).Model(&models.ChallengeCommentModel{}).Find(&c, "challenge_id = ?", challengeId).Error

	if err != nil {
		return nil, err
	}

	return c, nil
}

func (cc *CommentChallengeRepository) Delete(ctx context.Context, commentId string) error {

	err := cc.db.WithContext(ctx).Model(&models.ChallengeCommentModel{}).Where("id = ?", commentId).Delete(commentId).Error

	if err != nil {
		return err
	}

	return nil
}
