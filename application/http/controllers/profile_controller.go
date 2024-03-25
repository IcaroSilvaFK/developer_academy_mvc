package controllers

import (
	"log/slog"

	"github.com/IcaroSilvaFK/developer_academy_mvc/application/http/views"
	"github.com/IcaroSilvaFK/developer_academy_mvc/application/services"
	"github.com/IcaroSilvaFK/developer_academy_mvc/application/utils"
	"github.com/gin-gonic/gin"
)

type ProfileController struct {
	svc services.ChallengeServiceInterface
}

type ProfileControllerInterface interface {
	Index(ctx *gin.Context)
}

func NewProfileController(
	svc services.ChallengeServiceInterface,
) ProfileControllerInterface {
	return &ProfileController{
		svc,
	}
}

func (pc *ProfileController) Index(ctx *gin.Context) {

	userId := ctx.Param("id")

	u := utils.GetCurrentUserInRequestContext(ctx)

	challenges, err := pc.svc.FindByUserId(userId)

	if err != nil {
		slog.Error(err.Error())
	}

	result := views.NewChallengeResponseOutputList(challenges)

	ctx.HTML(200, "profile.gotmpl", gin.H{
		"user":       u,
		"challenges": result,
	})
}
