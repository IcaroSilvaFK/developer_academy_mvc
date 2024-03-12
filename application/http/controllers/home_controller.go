package controllers

import "github.com/gin-gonic/gin"

type HomeController struct{}

type HomeControllerInterface interface {
	Index(ctx *gin.Context)
}

func NewHomeController() HomeControllerInterface {
	return &HomeController{}
}

func (hc *HomeController) Index(ctx *gin.Context) {
	ctx.HTML(200, "home.tmpl", nil)
}
