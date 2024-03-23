package controllers

import (
	"net/http"

	"github.com/IcaroSilvaFK/developer_academy_mvc/application/http/views"
	"github.com/IcaroSilvaFK/developer_academy_mvc/application/services"
	"github.com/IcaroSilvaFK/developer_academy_mvc/application/utils"
	"github.com/gin-gonic/gin"
)

type CommentsChallengeController struct {
	svc services.CommentChallengeServiceInterface
}

type CommentsChallengeControllerInterface interface {
	Create(ctx *gin.Context)
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
		ctx.JSON(http.StatusBadRequest, nil)
		return
	}

	var input views.CommentChallengeInputView

	if err := ctx.Bind(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, nil)
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
