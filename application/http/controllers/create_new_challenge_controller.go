package controllers

import "github.com/gin-gonic/gin"

type CreateNewChallengeController struct {
}

type CreateNewChallengeControllerInterface interface {
	Index(ctx *gin.Context)
}

func NewCreateNewChallengeController() CreateNewChallengeControllerInterface {
	return &CreateNewChallengeController{}
}

func (cc *CreateNewChallengeController) Index(ctx *gin.Context) {
	ctx.HTML(200, "new_challenge.tmpl", nil)
}
