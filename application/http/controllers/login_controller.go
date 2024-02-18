package controllers

import "github.com/gin-gonic/gin"

type LoginController struct {
}

type LoginControllerInterface interface {
	Login(ctx *gin.Context)
}

func NewLoginController() LoginControllerInterface {
	return &LoginController{}
}

func (c *LoginController) Login(ctx *gin.Context) {
	ctx.HTML(200, "login.tmpl", nil)
}
