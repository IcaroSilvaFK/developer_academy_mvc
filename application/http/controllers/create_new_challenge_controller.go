package controllers

import (
	"net/http"

	"github.com/IcaroSilvaFK/developer_academy_mvc/application/http/views"
	"github.com/IcaroSilvaFK/developer_academy_mvc/application/services"
	"github.com/IcaroSilvaFK/developer_academy_mvc/application/utils"
	"github.com/gin-gonic/gin"
)

type CreateNewChallengeController struct {
	svc            services.ChallengeServiceInterface
	sessionService services.SessionServiceInterface
}

type CreateNewChallengeControllerInterface interface {
	Index(ctx *gin.Context)
	Create(ctx *gin.Context)
}

func NewCreateNewChallengeController(
	svc services.ChallengeServiceInterface,
	sessionService services.SessionServiceInterface,
) CreateNewChallengeControllerInterface {
	return &CreateNewChallengeController{
		svc, sessionService,
	}
}

func (cc *CreateNewChallengeController) Index(ctx *gin.Context) {

	u := utils.GetCurrentUserInRequestContext(ctx)

	ctx.HTML(http.StatusOK, "new_challenge.gotmpl", gin.H{
		"user": u,
	})
}

func (cc *CreateNewChallengeController) Create(ctx *gin.Context) {

	u := utils.GetCurrentUserInRequestContext(ctx)

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

	if err := cc.svc.Create(c.Title, c.Description, c.EmbedUrl, u.ID); err != nil {
		ctx.JSON(err.Code, err)
		return
	}

	ctx.JSON(http.StatusCreated, nil)
}
