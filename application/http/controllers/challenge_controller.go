package controllers

import (
	"net/http"

	"github.com/IcaroSilvaFK/developer_academy_mvc/application/services"
	"github.com/IcaroSilvaFK/developer_academy_mvc/infra/utils"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type ChallengeController struct {
	svc services.ChallengeServiceInterface
}

type ChallengeControllerInterface interface {
	Index(ctx *gin.Context)
}

func NewChallengeController(
	svc services.ChallengeServiceInterface,
) ChallengeControllerInterface {
	return &ChallengeController{
		svc,
	}
}

func (cc *ChallengeController) Index(ctx *gin.Context) {

	id := ctx.Param("id")

	if id == "" || !utils.IsValidId(id) {
		//TODO implment redirect to error page
		ctx.JSON(http.StatusBadRequest, nil)
		return
	}

	c, err := cc.svc.FindById(id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, nil)
		return
	}

	d, _ := c.CreatedAt.UTC().MarshalText()

	ctx.HTML(http.StatusOK, "challenge.gotmpl", gin.H{
		"challenge": c,
		"created":   string(d),
	})
}

func (cc *ChallengeController) Create(ctx *gin.Context) {

	session := sessions.Default(ctx)

	session.Get("")

	ctx.JSON(http.StatusCreated, nil)
}
