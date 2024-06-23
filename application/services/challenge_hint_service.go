package services

import (
	"context"
	"fmt"

	"github.com/IcaroSilvaFK/developer_academy_mvc/application/utils"
	"github.com/IcaroSilvaFK/developer_academy_mvc/infra/models"
	"github.com/IcaroSilvaFK/developer_academy_mvc/infra/repositories"
)

type ChallengeHintService struct {
	repo repositories.ChallengesHintsRepositoryInterface
}

type ChallengeHintServiceInterface interface {
	Create(ctx context.Context, challengeId, text string) *utils.RestErr
	FindByChallengeId(ctx context.Context, challengeId string) (*models.ChallengeHintsModel, *utils.RestErr)
	FindById(ctx context.Context, id string) (*models.ChallengeHintsModel, *utils.RestErr)
	Delete(ctx context.Context, id string) *utils.RestErr
}

func NewChallengeHintService(
	repo repositories.ChallengesHintsRepositoryInterface,
) ChallengeHintServiceInterface {

	return &ChallengeHintService{
		repo,
	}
}

func (ch *ChallengeHintService) Create(ctx context.Context, challengeId, text string) *utils.RestErr {

	h := models.NewChallengeHintsModel(challengeId, text)

	err := ch.repo.Create(ctx, h)

	if err != nil {
		m := "ERROR ON CREATE NEW HINT"
		return utils.NewInternalServerError(&m)
	}

	return nil
}

func (ch *ChallengeHintService) FindByChallengeId(
	ctx context.Context,
	challengeId string,
) (*models.ChallengeHintsModel, *utils.RestErr) {

	h, err := ch.repo.GetByChallengeId(ctx, challengeId)

	if err != nil {
		m := "ERROR ON FIND CHALLENGE HINT"
		return nil, utils.NewInternalServerError(&m)
	}

	if h == nil {
		return nil, utils.NewNotFoundException("THE CHALLENGE HINT NOT EXISTS")
	}

	return h, nil
}

func (ch *ChallengeHintService) FindById(ctx context.Context, id string) (*models.ChallengeHintsModel, *utils.RestErr) {

	h, err := ch.repo.GetById(ctx, id)

	if err != nil {
		m := "ERROR ON FIND HINT BY ID"
		return nil, utils.NewInternalServerError(&m)
	}

	return h, nil
}

func (ch *ChallengeHintService) Delete(ctx context.Context, id string) *utils.RestErr {

	h, _ := ch.repo.GetById(ctx, id)

	if h == nil {
		return utils.NewNotFoundException(fmt.Sprintf("THE HINT WITH ID: %s NOT  EXISTS", id))
	}

	err := ch.repo.Delete(ctx, id)

	if err != nil {
		m := "ERROR ON DELETE HINT"
		return utils.NewInternalServerError(&m)
	}

	return nil
}
