package controller

import (
	"net/http"

	"github.com/Caknoooo/golang-clean_template/dto"
	"github.com/Caknoooo/golang-clean_template/services"
	"github.com/Caknoooo/golang-clean_template/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type TopupController interface {
	CreateTopup(ctx *gin.Context)
	GetAllTopupUser(ctx *gin.Context)
	GetTopupByID(ctx *gin.Context)
}

type topupController struct {
	topupService services.TopupService
	jwtService services.JWTService
}

func NewTopupController(ts services.TopupService, jwt services.JWTService) TopupController {
	return &topupController{
		topupService: ts,
		jwtService: jwt,
	}
}

func (tc *topupController) CreateTopup(ctx *gin.Context) {
	token := ctx.MustGet("token").(string)
	userID, err := tc.jwtService.GetUserIDByToken(token)
	if err != nil {
		res := utils.BuildResponseFailed("Gagal Mendapatkan User", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	var topup dto.TopupCreateDTO
	if err := ctx.ShouldBindJSON(&topup); err != nil {
		res := utils.BuildResponseFailed("Gagal Topup", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	result, err := tc.topupService.CreateTopup(ctx.Request.Context(), topup, userID)
	if err != nil {
		res := utils.BuildResponseFailed("Gagal Topup", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess("Berhasil Topup", result)
	ctx.JSON(http.StatusOK, res)
}

func (tc *topupController) GetAllTopupUser(ctx *gin.Context) {
	token := ctx.MustGet("token").(string)
	userID, err := tc.jwtService.GetUserIDByToken(token)
	if err != nil {
		res := utils.BuildResponseFailed("Gagal Mendapatkan User", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	result, err := tc.topupService.GetAllTopupUser(ctx.Request.Context(), userID)
	if err != nil {
		res := utils.BuildResponseFailed("Gagal Mendapatkan List Topup", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess("Berhasil Mendapatkan List Topup", result)
	ctx.JSON(http.StatusOK, res)
}

func (tc *topupController) GetTopupByID(ctx *gin.Context) {
	id := ctx.Param("id")
	uuid, err := uuid.Parse(id)
	if err != nil {
		res := utils.BuildResponseFailed("Gagal Parse ID", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	result, err := tc.topupService.GetTopupByID(ctx.Request.Context(), uuid)
	if err != nil {
		res := utils.BuildResponseFailed("Gagal Mendapatkan Topup", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess("Berhasil Mendapatkan Topup", result)
	ctx.JSON(http.StatusOK, res)
}