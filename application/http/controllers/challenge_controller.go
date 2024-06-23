package controllers

import (
	"net/http"

	"github.com/IcaroSilvaFK/developer_academy_mvc/application/http/views"
	"github.com/IcaroSilvaFK/developer_academy_mvc/application/services"
	apputils "github.com/IcaroSilvaFK/developer_academy_mvc/application/utils"
	"github.com/IcaroSilvaFK/developer_academy_mvc/infra/utils"

	"github.com/gin-gonic/gin"
)

type ChallengeController struct {
	svc services.ChallengeServiceInterface
}

type ChallengeControllerInterface interface {
	Index(ctx *gin.Context)
	Destroy(ctx *gin.Context)
	GetAllChallenges(ctx *gin.Context)
	FindById(ctx *gin.Context)
	FindUserId(ctx *gin.Context)
}

func NewChallengeController(
	svc services.ChallengeServiceInterface,
) ChallengeControllerInterface {
	return &ChallengeController{
		svc,
	}
}

func (cc *ChallengeController) Index(ctx *gin.Context) {

	id := ctx.Param("id")

	if !utils.IsValidId(id) {
		ctx.Redirect(http.StatusPermanentRedirect, "/error")
		return
	}

	c, err := cc.svc.FindById(ctx.Request.Context(), id)

	if err != nil {
		ctx.Redirect(http.StatusPermanentRedirect, "/error")
		return
	}

	d, _ := c.CreatedAt.UTC().MarshalText()
	u := apputils.GetCurrentUserInRequestContext(ctx)

	countComments := len(c.Comments)

	ctx.HTML(http.StatusOK, "challenge.gotmpl", gin.H{
		"challenge":        c,
		"created":          string(d),
		"quantityComments": countComments,
		"comments":         views.NewCommentChallengeListOutputView(c.Comments),
		"id":               id,
		"user":             u,
	})
}

// @Summary	Create new challenge comment
// @Description	Create new challenge comment
// @Tags			challenges
// @Param		  id path string true "the id from challenge"
// @Produce		json
// @Success		204
// @Failure		400	{object}	utils.RestErr
// @Failure		500	{object}	utils.RestErr
// @Router		/challenges/{id} [delete]
func (cc *ChallengeController) Destroy(ctx *gin.Context) {
	id := ctx.Param("id")

	if !utils.IsValidId(id) {
		err := apputils.NewBadRequestException(apputils.INVALID_ID_MESSAGE)
		ctx.JSON(err.Code, err)
		return
	}

	if err := cc.svc.Delete(ctx.Request.Context(), id); err != nil {
		ctx.JSON(err.Code, err)
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}

// @Summary	Find all challenges
// @Description	Find all challenges
// @Tags			challenges
// @Produce		json
// @Success		200 {array}  	views.ResponseChallengeOutputView
// @Failure		400	{object}	utils.RestErr
// @Failure		500	{object}	utils.RestErr
// @Router		/challenges [get]
func (cc *ChallengeController) GetAllChallenges(ctx *gin.Context) {

	challenges, err := cc.svc.FindAll(ctx.Request.Context(), nil)

	if err != nil {
		ctx.JSON(err.Code, err)
		return
	}

	r := views.NewChallengeResponseOutputList(challenges)

	ctx.JSON(http.StatusOK, r)
}

// @Summary	Find all challenges
// @Description	Find all challenges
// @Tags			challenges
// @Param     id path string true "The id from challegen"
// @Produce		json
// @Success		200 {object}  	views.ResponseChallengeOutputView
// @Failure		400	{object}	utils.RestErr
// @Failure		500	{object}	utils.RestErr
// @Router		/challenges/{id} [get]
func (cc *ChallengeController) FindById(ctx *gin.Context) {

	id := ctx.Param("id")

	if !utils.IsValidId(id) {
		err := apputils.NewBadRequestException(apputils.INVALID_ID_MESSAGE)

		ctx.JSON(err.Code, err)
		return
	}

	c, err := cc.svc.FindById(ctx.Request.Context(), id)

	if err != nil {
		ctx.JSON(err.Code, err)
		return
	}

	r := views.NewChallengeResponseOutput(c)

	ctx.JSON(http.StatusOK, r)
}

// @Summary	Find all challenges
// @Description	Find all challenges
// @Tags			challenges
// @Param     userId path string true "The id from challegen"
// @Produce		json
// @Success		200 {object}  	views.ResponseChallengeOutputView
// @Failure		400	{object}	utils.RestErr
// @Failure		500	{object}	utils.RestErr
// @Router		/challenges/users/{userId} [get]
func (cc *ChallengeController) FindUserId(ctx *gin.Context) {

	id := ctx.Param("userId")

	if !utils.IsValidId(id) {
		err := apputils.NewBadRequestException(apputils.INVALID_ID_MESSAGE)

		ctx.JSON(err.Code, err)
		return
	}

	c, err := cc.svc.FindByUserId(ctx.Request.Context(), id)

	if err != nil {

		ctx.JSON(err.Code, err)
		return
	}

	r := views.NewChallengeResponseOutputList(c)

	ctx.JSON(http.StatusOK, r)

}
