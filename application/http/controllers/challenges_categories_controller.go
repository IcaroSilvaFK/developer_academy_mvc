package controllers

import (
	"net/http"

	"github.com/IcaroSilvaFK/developer_academy_mvc/application/http/views"
	"github.com/IcaroSilvaFK/developer_academy_mvc/application/services"
	"github.com/IcaroSilvaFK/developer_academy_mvc/application/utils"
	infrautils "github.com/IcaroSilvaFK/developer_academy_mvc/infra/utils"
	"github.com/gin-gonic/gin"
)

type ChallengesCategoriesController struct {
	svc services.ChallengesCategoriesServiceInterface
}

type ChallengesCategoriesControllerInterface interface {
	Create(ctx *gin.Context)
	FindAll(ctx *gin.Context)
	FindById(ctx *gin.Context)
	FindByUserId(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

func NewChallengesCategoriesController(
	svc services.ChallengesCategoriesServiceInterface,
) ChallengesCategoriesControllerInterface {
	return &ChallengesCategoriesController{
		svc,
	}
}

// @Summary	Create new challenge category
// @Description	Create new challenge category
// @Tags			challenges categories
// @Accept    json
// @Param		request body views.ChallengesCategoriesInputView required "body"
// @Produce		json
// @Success		204
// @Failure		400	{object}	utils.RestErr
// @Failure		500	{object}	utils.RestErr
// @Router		/challenges/categories [post]
func (c *ChallengesCategoriesController) Create(ctx *gin.Context) {
	goContext := ctx.Request.Context()
	u := utils.GetCurrentUserInRequestContext(ctx).(views.UserResponseView)

	var input *views.ChallengesCategoriesInputView

	if err := ctx.Bind(&input); err != nil {
		errRes := utils.NewBadRequestException(err.Error())

		ctx.JSON(errRes.Code, errRes)
		return
	}

	input.UserId = u.ID

	err := c.svc.Create(goContext, input)

	if err != nil {
		ctx.JSON(err.Code, err)
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}

// @Summary	Delete challenge category
// @Description	Delete challenge category
// @Tags			challenges categories
// @Produce		json
// @Success		204
// @Failure		400	{object}	utils.RestErr
// @Failure		500	{object}	utils.RestErr
// @Router		/challenges/categories/{id} [delete]
func (c *ChallengesCategoriesController) Delete(ctx *gin.Context) {

	id := ctx.Param("id")
	goContext := ctx.Request.Context()

	if !infrautils.IsValidId(id) {
		err := utils.NewBadRequestException(
			"The param id is mandatory in request and valid uuid",
		)

		ctx.JSON(err.Code, err)
		return
	}

	if err := c.svc.Delete(goContext, id); err != nil {
		ctx.JSON(err.Code, err)
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}

// @Summary	Find all challenge category
// @Description	Find all challenge category
// @Tags			challenges categories
// @Produce		json
// @Success		200 {array}	 views.ChallengesCategoriesOutputView
// @Failure		400	{object}	utils.RestErr
// @Failure		500	{object}	utils.RestErr
// @Router		/challenges/categories [get]
func (c *ChallengesCategoriesController) FindAll(ctx *gin.Context) {
	goContext := ctx.Request.Context()

	q := ctx.Param("query")

	r, err := c.svc.GetAll(goContext, q)

	if err != nil {
		ctx.JSON(err.Code, err)
		return
	}

	ctx.JSON(http.StatusOK, r)
}

// @Summary	Find by id challenge category
// @Description	Find by id challenge category
// @Tags			challenges categories
// @Produce		json
// @Success		200 {object}	 views.ChallengesCategoriesOutputView
// @Failure		400	{object}	utils.RestErr
// @Failure		500	{object}	utils.RestErr
// @Router		/challenges/categories/{id} [get]
func (c *ChallengesCategoriesController) FindById(ctx *gin.Context) {
	id := ctx.Param("id")
	goContext := ctx.Request.Context()

	if !infrautils.IsValidId(id) {
		err := utils.NewBadRequestException("The id provided is invalid or bad formed please provide valid id.")

		ctx.JSON(err.Code, err)
		return
	}

	r, err := c.svc.GetById(goContext, id)

	if err != nil {
		ctx.JSON(err.Code, err)
		return
	}

	ctx.JSON(http.StatusOK, r)
}

// @Summary	Find by user id challenge category
// @Description	Find by user id challenge category
// @Tags			challenges categories
// @Produce		json
// @Success		204
// @Failure		400	{object}	utils.RestErr
// @Failure		500	{object}	utils.RestErr
// @Router		/challenges/categories/users/{id} [get]
func (c *ChallengesCategoriesController) FindByUserId(ctx *gin.Context) {
	userId := ctx.Param("userId")
	goContext := ctx.Request.Context()

	if !infrautils.IsValidId(userId) {
		err := utils.NewBadRequestException("The id provided is invalid")

		ctx.JSON(err.Code, err)
		return
	}

	r, err := c.svc.FindByUserId(goContext, userId)

	if err != nil {
		ctx.JSON(http.StatusOK, r)

		return
	}
}

// @Summary	Find by user id challenge category
// @Description	Find by user id challenge category
// @Tags			challenges categories
// @Accept    json
// @Param		request body views.ChallengesCategoriesInputView required "body"
// @Produce		json
// @Success		204
// @Failure		400	{object}	utils.RestErr
// @Failure		500	{object}	utils.RestErr
// @Router		/challenges/categories/{id} [put]
func (c *ChallengesCategoriesController) Update(ctx *gin.Context) {
	id := ctx.Param("id")
	goContext := ctx.Request.Context()
	u := utils.GetCurrentUserInRequestContext(ctx).(views.UserResponseView)

	if !infrautils.IsValidId(id) {
		err := utils.NewBadRequestException("The id provided is invalid")

		ctx.JSON(err.Code, err)
		return
	}

	var body *views.ChallengesCategoriesInputView

	if err := ctx.Bind(&body); err != nil {

		err := utils.NewBadRequestException(err.Error())

		ctx.JSON(err.Code, err)
		return
	}

	err := c.svc.Update(goContext, id, body.Title, u.ID)

	if err != nil {
		ctx.JSON(err.Code, err)
		return
	}

	ctx.JSON(http.StatusNoContent, nil)

}
