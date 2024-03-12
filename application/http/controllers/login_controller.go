package controllers

import (
	"net/http"

	"github.com/IcaroSilvaFK/developer_academy_mvc/application/http/views"
	"github.com/IcaroSilvaFK/developer_academy_mvc/application/services"
	"github.com/gin-gonic/gin"
)

type LoginController struct {
	svc services.LoginServiceInterface
}

type LoginControllerInterface interface {
	Login(ctx *gin.Context)
	SignIn(ctx *gin.Context)
}

func NewLoginController(
	svc services.LoginServiceInterface,
) LoginControllerInterface {
	return &LoginController{
		svc,
	}
}

func (c *LoginController) Login(ctx *gin.Context) {
	ctx.HTML(200, "login.tmpl", nil)
}

func (c *LoginController) SignIn(ctx *gin.Context) {
	code := ctx.Param("code")

	if code == "" {
		ctx.JSON(http.StatusNoContent, gin.H{
			"message": "Bad request missing code",
		})
		return
	}

	u, err := c.svc.Login(code)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	r := views.NewLoginResponseView(u)

	ctx.JSON(http.StatusOK, gin.H{
		"user": r,
	})
}
