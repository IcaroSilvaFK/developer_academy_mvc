package controllers

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type ChallengeController struct {
}

type ChallengeControllerInterface interface {
	Index(ctx *gin.Context)
}

func NewChallengeController() ChallengeControllerInterface {
	return &ChallengeController{}
}

func (cc *ChallengeController) Index(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "challenge.gotmpl", nil)
}

func (cc *ChallengeController) Create(ctx *gin.Context) {

	session := sessions.Default(ctx)

	session.Get("")

	ctx.JSON(http.StatusCreated, nil)
}
