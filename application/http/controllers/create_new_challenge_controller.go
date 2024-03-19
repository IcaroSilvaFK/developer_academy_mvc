package controllers

import (
	"net/http"

	"github.com/IcaroSilvaFK/developer_academy_mvc/application/http/views"
	"github.com/IcaroSilvaFK/developer_academy_mvc/application/services"
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
	ctx.HTML(http.StatusOK, "new_challenge.gotmpl", nil)
}

func (cc *CreateNewChallengeController) Create(ctx *gin.Context) {

	var u views.UserResponseView

	cc.sessionService.Get(ctx, "user", &u)

	if u.ID == "" {
		//TODO padronizar as respostas de erro
		ctx.JSON(http.StatusBadRequest, gin.H{
			"err": "Usuário inválido",
		})
		return
	}

	var c views.CreateChallengeInputView

	if err := ctx.Bind(&c); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	if err := cc.svc.Create(c.Title, c.Description, c.EmbedUrl, u.ID); err != nil {

		ctx.JSON(http.StatusInternalServerError, nil)
		return
	}

	ctx.JSON(http.StatusCreated, nil)
}
