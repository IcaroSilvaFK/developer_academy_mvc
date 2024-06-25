package services

import (
	"context"

	"github.com/IcaroSilvaFK/developer_academy_mvc/application/http/views"
	"github.com/IcaroSilvaFK/developer_academy_mvc/application/utils"
	"github.com/IcaroSilvaFK/developer_academy_mvc/infra/models"
	"github.com/IcaroSilvaFK/developer_academy_mvc/infra/repositories"
	infrautils "github.com/IcaroSilvaFK/developer_academy_mvc/infra/utils"
	"go.uber.org/zap"

	"gorm.io/gorm"
)

type ChallengeService struct {
	repo      repositories.ChallengeRepositoryInterface
	hinrepo   ChallengeHintServiceInterface
	iaservice AIServiceInterface
	cache     CacheServiceInterface
	ck        string
}

type ChallengeServiceInterface interface {
	Create(ctx context.Context, input views.CreateChallengeInputView, userId string) *utils.RestErr
	FindAll(ctx context.Context, page *int) ([]*models.ChallengeModel, *utils.RestErr)
	FindById(ctx context.Context, id string) (*models.ChallengeModel, *utils.RestErr)
	FindByUserId(ctx context.Context, id string) ([]*models.ChallengeModel, *utils.RestErr)
	CountChallenges(context.Context) (int, *utils.RestErr)
	Delete(ctx context.Context, id string) *utils.RestErr
}

func NewChallengeService(
	repo repositories.ChallengeRepositoryInterface,
	hinrepo ChallengeHintServiceInterface,
	iaservice AIServiceInterface,
	cacheSvc CacheServiceInterface,
) ChallengeServiceInterface {

	return &ChallengeService{
		repo, hinrepo, iaservice, cacheSvc, "challenges",
	}
}

func (c *ChallengeService) Create(ctx context.Context, input views.CreateChallengeInputView, userId string) *utils.RestErr {

	if !c.iaservice.VerifyIfIsValidChallenge(input.Title) {
		return utils.NewBadRequestException("Te request contains params inappropriate")
	}
	cm := models.NewChallengeModel(input.Title, input.Description, input.EmbedUrl, userId, input.Categories)

	err := c.repo.Create(ctx, cm)

	if err != nil {
		message := "Error on create new challenge"
		return utils.NewInternalServerError(&message)
	}

	hint, err := c.iaservice.GenerateHintFromChallenge(input.Title)

	if err != nil {
		message := "Error on generate hint from challenge"
		return utils.NewInternalServerError(&message)
	}

	if hinErr := c.hinrepo.Create(ctx, cm.ID, hint); hinErr != nil {

		return hinErr
	}

	if err := c.cache.Delete(c.ck); err != nil {
		utils.Error("Error on delete cache from challenges", err)
	}

	return nil
}

func (c *ChallengeService) FindAll(ctx context.Context, page *int) ([]*models.ChallengeModel, *utils.RestErr) {

	var res []*models.ChallengeModel

	if err := c.cache.Get(c.ck, &res); err != nil {
		utils.Error("Erro on get in cache", err)
	}

	if len(res) > 0 {
		utils.Info("Response from cache", zap.String("redis", "ok"))
		return res, nil
	}

	res, err := c.repo.GetAll(ctx, page)

	if err != nil {
		message := "Error on get all challenges please try again later"
		return nil, utils.NewInternalServerError(&message)
	}

	if err := c.cache.Set(c.ck, res); err != nil {
		utils.Error("Error on insert items in cache", err)
	}

	return res, nil
}

func (c *ChallengeService) FindById(ctx context.Context, id string) (*models.ChallengeModel, *utils.RestErr) {

	if !infrautils.IsValidId(id) {
		return nil, utils.NewBadRequestException("ID provided is invalid")
	}

	r, err := c.repo.GetById(ctx, id)

	if err == gorm.ErrRecordNotFound {
		return nil, utils.NewNotFoundException("This challenge not exists")
	}

	if err != nil {
		message := "Error on find challenge"
		return nil, utils.NewInternalServerError(&message)
	}
	return r, nil
}

func (c *ChallengeService) FindByUserId(ctx context.Context, id string) ([]*models.ChallengeModel, *utils.RestErr) {

	if !infrautils.IsValidId(id) {
		return nil, utils.NewBadRequestException("Please provide a valid id")
	}

	r, err := c.repo.GetByUserId(ctx, id)

	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}

	if err != nil {
		message := "Error on find user"
		return nil, utils.NewInternalServerError(&message)
	}

	return r, nil
}

func (c *ChallengeService) CountChallenges(ctx context.Context) (int, *utils.RestErr) {
	r, err := c.repo.CountChallenges(ctx)

	if err != nil {
		message := "Error on count challenges"
		return 0, utils.NewInternalServerError(&message)
	}

	return r, nil
}

func (c *ChallengeService) Delete(ctx context.Context, id string) *utils.RestErr {

	if !infrautils.IsValidId(id) {
		return utils.NewBadRequestException("ID provided is invalid")
	}

	err := c.repo.Delete(ctx, id)

	if err == gorm.ErrRecordNotFound {
		return utils.NewNotFoundException("Id provide not exists in challenges")
	}

	if err != nil {
		message := "Error on delete challenge"
		return utils.NewInternalServerError(&message)
	}

	if err := c.cache.Delete(c.ck); err != nil {
		utils.Error("Failed to delete challanges to cache", err)
	}

	return nil
}
