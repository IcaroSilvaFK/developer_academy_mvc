package controllers

import "github.com/gin-gonic/gin"

type ChallengeController struct {
}

type ChallengeControllerInterface interface {
	Index(ctx *gin.Context)
}

func NewChallengeController() ChallengeControllerInterface {
	return &ChallengeController{}
}

func (cc *ChallengeController) Index(ctx *gin.Context) {
	ctx.HTML(200, "challenge.tmpl", nil)
}
