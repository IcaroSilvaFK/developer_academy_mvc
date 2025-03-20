package controllers

import (
	"net/http"

	"github.com/IcaroSilvaFK/developer_academy_mvc/application/dtos"
	"github.com/IcaroSilvaFK/developer_academy_mvc/application/http/views"
	"github.com/IcaroSilvaFK/developer_academy_mvc/application/services"
	"github.com/IcaroSilvaFK/developer_academy_mvc/application/utils"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	svc services.UserServiceInterface
}

type UserControllerInterface interface {
	CreateUser(*gin.Context)
}

func NewUserController(
	svc services.UserServiceInterface,
) UserControllerInterface {
	return &UserController{
		svc,
	}
}

func (u *UserController) CreateUser(c *gin.Context) {
	ctx := c.Request.Context()

	input := new(dtos.CreateUserInputDto)

	if err := c.ShouldBindJSON(input); err != nil {
		errno := utils.NewBadRequestException(err.Error())

		c.JSON(errno.Code, errno)
		return
	}

	r, err := u.svc.CreateUser(ctx, input)

	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	user := views.NewUserResponseView(r)

	c.JSON(http.StatusCreated, user)
	return
}
