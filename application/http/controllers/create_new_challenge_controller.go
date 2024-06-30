package controllers

import (
	"net/http"

	"github.com/IcaroSilvaFK/developer_academy_mvc/application/http/views"
	"github.com/IcaroSilvaFK/developer_academy_mvc/application/services"
	"github.com/IcaroSilvaFK/developer_academy_mvc/application/utils"
	"github.com/gin-gonic/gin"
)

type CreateNewChallengeController struct {
	svc             services.ChallengeServiceInterface
	sessionService  services.SessionServiceInterface
	categoryService services.ChallengesCategoriesServiceInterface
}

type CreateNewChallengeControllerInterface interface {
	Index(ctx *gin.Context)
	Create(ctx *gin.Context)
}

func NewCreateNewChallengeController(
	svc services.ChallengeServiceInterface,
	sessionService services.SessionServiceInterface,
	categoryService services.ChallengesCategoriesServiceInterface,
) CreateNewChallengeControllerInterface {
	return &CreateNewChallengeController{
		svc, sessionService, categoryService,
	}
}

func (cc *CreateNewChallengeController) Index(ctx *gin.Context) {

	u := utils.GetCurrentUserInRequestContext(ctx)

	cats, err := cc.categoryService.GetAll(ctx.Request.Context(), "")

	if err != nil {
		utils.Error("Error on get all categories", err)
	}

	ctx.HTML(http.StatusOK, "new_challenge.gotmpl", gin.H{
		"user":       u,
		"categories": cats,
	})
}

// @Summary	Create new challenge
// @Description	Create new challenge
// @Tags			challenges
// @Accept    json
// @Param		request body views.CreateChallengeInputView required "body"
// @Produce		json
// @Success		201
// @Failure		400	{object}	utils.RestErr
// @Failure		500	{object}	utils.RestErr
// @Router		/challenges [post]
func (cc *CreateNewChallengeController) Create(ctx *gin.Context) {

	u := utils.GetCurrentUserInRequestContext(ctx).(views.UserResponseView)

	if u.ID == "" {
		err := utils.NewBadRequestException("ID is required but is missing in current request")
		ctx.JSON(err.Code, err)
		return
	}

	var c views.CreateChallengeInputView

	if err := ctx.Bind(&c); err != nil {
		erno := utils.NewBadRequestException(err.Error())
		ctx.JSON(erno.Code, erno)
		return
	}

	if err := cc.svc.Create(ctx.Request.Context(), c, u.ID); err != nil {
		ctx.JSON(err.Code, err)
		return
	}

	ctx.JSON(http.StatusCreated, nil)
}
