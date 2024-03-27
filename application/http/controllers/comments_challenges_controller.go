package controllers

import (
	"net/http"

	"github.com/IcaroSilvaFK/developer_academy_mvc/application/http/views"
	"github.com/IcaroSilvaFK/developer_academy_mvc/application/services"
	"github.com/IcaroSilvaFK/developer_academy_mvc/application/utils"

	infrautils "github.com/IcaroSilvaFK/developer_academy_mvc/infra/utils"
	"github.com/gin-gonic/gin"
)

type CommentsChallengeController struct {
	svc services.CommentChallengeServiceInterface
}

type CommentsChallengeControllerInterface interface {
	Create(ctx *gin.Context)
	Destroy(ctx *gin.Context)
}

func NewCommentsChallengeController(
	svc services.CommentChallengeServiceInterface,
) CommentsChallengeControllerInterface {

	return &CommentsChallengeController{
		svc,
	}
}

func (cc *CommentsChallengeController) Create(ctx *gin.Context) {

	u := utils.GetCurrentUserInRequestContext(ctx)

	if u.ID == "" {
		err := utils.NewBadRequestException("Current user is not logged.")

		ctx.JSON(err.Code, err)
		return
	}

	var input views.CommentChallengeInputView

	if err := ctx.Bind(&input); err != nil {
		err := utils.NewBadRequestException(err.Error())

		ctx.JSON(err.Code, err)
		return
	}

	c, err := cc.svc.Create(input.ChallengeId, u.ID, input.Comment)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, nil)
		return
	}

	result := views.NewCommentChallengeOutputView(c)

	ctx.JSON(http.StatusCreated, result)
}

func (cc *CommentsChallengeController) Destroy(ctx *gin.Context) {
	id := ctx.Param("id")

	if !infrautils.IsValidId(id) {
		err := utils.NewBadRequestException("The id provided is not valid")
		ctx.JSON(err.Code, err)
		return
	}

	err := cc.svc.Delete(id)

	if err != nil {
		ctx.JSON(err.Code, err)
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
