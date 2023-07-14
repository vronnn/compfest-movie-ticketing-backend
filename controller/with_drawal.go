package controller

import (
	"net/http"

	"github.com/Caknoooo/golang-clean_template/dto"
	"github.com/Caknoooo/golang-clean_template/services"
	"github.com/Caknoooo/golang-clean_template/utils"
	"github.com/gin-gonic/gin"
)

type WithDrawalController interface {
	CreateWithDrawal(ctx *gin.Context)
	GetAllWithDrawalUser(ctx *gin.Context)
}

type withDrawalController struct {
	withDrawalService services.WithDrawalService
	jwtService services.JWTService
}

func NewWithDrawalController(ws services.WithDrawalService, jwt services.JWTService) WithDrawalController {
	return &withDrawalController{
		withDrawalService: ws,
		jwtService: jwt,
	}
}

func (wc *withDrawalController) CreateWithDrawal(ctx *gin.Context) {
	token := ctx.MustGet("token").(string)	
	userID, err := wc.jwtService.GetUserIDByToken(token)
	if err != nil {
		res := utils.BuildResponseFailed("Gagal Mendapatkan User", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	var withDrawal dto.CreateWithDrawalDTO
	if err := ctx.ShouldBindJSON(&withDrawal); err != nil {
		res := utils.BuildResponseFailed("Gagal Mendapatkan Request Dari Body", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	result, err := wc.withDrawalService.CreateWithDrawal(ctx.Request.Context(), withDrawal, userID)
	if err != nil {
		res := utils.BuildResponseFailed("Gagal Melakukan Penarikan", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess("Berhasil Melakukan Penarikan", result)
	ctx.JSON(http.StatusOK, res)
}

func (wc *withDrawalController) GetAllWithDrawalUser(ctx *gin.Context) {
	token := ctx.MustGet("token").(string)
	userID, err := wc.jwtService.GetUserIDByToken(token)
	if err != nil {
		res := utils.BuildResponseFailed("Gagal Mendapatkan User", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	result, err := wc.withDrawalService.GetAllWithDrawalUser(ctx.Request.Context(), userID)
	if err != nil {
		res := utils.BuildResponseFailed("Gagal Mendapatkan History With Drawal", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess("Berhasil Mendapatkan History With Drawal", result)
	ctx.JSON(http.StatusOK, res)
}