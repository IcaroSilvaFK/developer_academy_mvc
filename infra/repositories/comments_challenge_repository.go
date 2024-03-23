package repositories

import (
	"github.com/IcaroSilvaFK/developer_academy_mvc/infra/models"
	"gorm.io/gorm"
)

type CommentChallengeRepository struct {
	db *gorm.DB
}

type CommentChallengeRepositoryInterface interface {
	Create(*models.ChallengeCommentModel) error
	FindById(commentId string) (*models.ChallengeCommentModel, error)
	FindByUserId(userId string) ([]*models.ChallengeCommentModel, error)
	FindByChallengeId(challengeId string) ([]*models.ChallengeCommentModel, error)
	Delete(commentId string) error
}

func NewCommentChallengeRepository(
	db *gorm.DB,
) CommentChallengeRepositoryInterface {

	return &CommentChallengeRepository{db}
}

func (cc *CommentChallengeRepository) Create(c *models.ChallengeCommentModel) error {

	// c := models.NewChallengeCommentMode(challengeId, userId, comment)

	err := cc.db.Model(&models.ChallengeCommentModel{}).Create(c).Error

	return err
}
func (cc *CommentChallengeRepository) FindById(commentId string) (*models.ChallengeCommentModel, error) {

	var c models.ChallengeCommentModel

	err := cc.db.Model(&models.ChallengeCommentModel{}).Find(&c, "id = ?", commentId).Error

	if err != nil {
		return nil, err
	}
	return &c, nil
}

func (cc *CommentChallengeRepository) FindByUserId(userId string) ([]*models.ChallengeCommentModel, error) {
	var c []*models.ChallengeCommentModel

	err := cc.db.Model(&models.ChallengeCommentModel{}).Find(&c, "user_id = ?", userId).Error
	if err != nil {
		return nil, err
	}

	return c, nil
}
func (cc *CommentChallengeRepository) FindByChallengeId(challengeId string) ([]*models.ChallengeCommentModel, error) {

	var c []*models.ChallengeCommentModel

	err := cc.db.Model(&models.ChallengeCommentModel{}).Find(&c, "challenge_id = ?", challengeId).Error

	if err != nil {
		return nil, err
	}

	return c, nil
}

func (cc *CommentChallengeRepository) Delete(commentId string) error {

	err := cc.db.Model(&models.ChallengeCommentModel{}).Delete(commentId).Error

	if err != nil {
		return err
	}

	return nil
}
