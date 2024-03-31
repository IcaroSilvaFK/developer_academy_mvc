package controllers

import (
	"net/http"

	"github.com/IcaroSilvaFK/developer_academy_mvc/application/http/views"
	"github.com/IcaroSilvaFK/developer_academy_mvc/application/services"
	"github.com/IcaroSilvaFK/developer_academy_mvc/application/utils"
	infrautils "github.com/IcaroSilvaFK/developer_academy_mvc/infra/utils"
	"github.com/gin-gonic/gin"
)

type ProfileController struct {
	svc  services.ChallengeServiceInterface
	usvc services.UserServiceInterface
}

type ProfileControllerInterface interface {
	Index(*gin.Context)
	FindByUserId(*gin.Context)
	FindAllUsers(*gin.Context)
	Delete(*gin.Context)
}

func NewProfileController(
	svc services.ChallengeServiceInterface,
	usvc services.UserServiceInterface,
) ProfileControllerInterface {
	return &ProfileController{
		svc, usvc,
	}
}

func (pc *ProfileController) Index(ctx *gin.Context) {

	userId := ctx.Param("id")

	u := utils.GetCurrentUserInRequestContext(ctx)

	challenges, err := pc.svc.FindByUserId(userId)

	if err != nil {
		utils.Error("Error while search user", err)
		ctx.Redirect(http.StatusPermanentRedirect, "/error")
		return
	}

	result := views.NewChallengeResponseOutputList(challenges)

	ctx.HTML(200, "profile.gotmpl", gin.H{
		"user":       u,
		"challenges": result,
	})
}

func (pc *ProfileController) FindByUserId(ctx *gin.Context) {

	id := ctx.Param("id")

	if !infrautils.IsValidId(id) {
		err := utils.NewBadRequestException(utils.INVALID_ID_MESSAGE)
		ctx.JSON(err.Code, err)
		return
	}

	u, err := pc.usvc.FindUserById(id)

	if err != nil {
		ctx.JSON(err.Code, err)
		return
	}

	ctx.JSON(http.StatusOK, u)
}

func (pc *ProfileController) FindAllUsers(ctx *gin.Context) {

	users, err := pc.usvc.FindAllUsers()

	if err != nil {
		ctx.JSON(err.Code, err)
		return
	}

	ctx.JSON(http.StatusOK, users)
}

func (pc *ProfileController) Delete(ctx *gin.Context) {

	id := ctx.Param("id")

	if !infrautils.IsValidId(id) {
		err := utils.NewBadRequestException(utils.INVALID_ID_MESSAGE)

		ctx.JSON(err.Code, err)
		return
	}

	err := pc.usvc.Delete(id)

	if err != nil {
		ctx.JSON(err.Code, err)
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
