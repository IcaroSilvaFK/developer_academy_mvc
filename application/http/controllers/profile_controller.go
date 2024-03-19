package controllers

import "github.com/gin-gonic/gin"

type ProfileController struct{}

type ProfileControllerInterface interface {
	Index(ctx *gin.Context)
}

func NewProfileController() ProfileControllerInterface {
	return &ProfileController{}
}

func (pc *ProfileController) Index(ctx *gin.Context) {
	ctx.HTML(200, "profile.gotmpl", nil)
}
