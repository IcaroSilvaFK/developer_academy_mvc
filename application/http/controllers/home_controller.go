package controllers

import (
	"fmt"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type HomeController struct{}

type HomeControllerInterface interface {
	Index(ctx *gin.Context)
}

func NewHomeController() HomeControllerInterface {
	return &HomeController{}
}

func (hc *HomeController) Index(ctx *gin.Context) {

	session := sessions.Default(ctx)

	fmt.Println(session.Get("user"))

	ctx.HTML(200, "home.gotmpl", nil)
}
