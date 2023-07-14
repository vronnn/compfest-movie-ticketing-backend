package controller

import (
	"net/http"
	"strconv"

	"github.com/Caknoooo/golang-clean_template/services"
	"github.com/Caknoooo/golang-clean_template/utils"
	"github.com/gin-gonic/gin"
)

type SeederController interface {
	GetAllBank(ctx *gin.Context)
	GetBankByID(ctx *gin.Context)
}

type seederController struct {
	seederService services.SeederService
}

func NewSeederController(ss services.SeederService) SeederController {
	return &seederController{
		seederService: ss,
	}
}

func (sc *seederController) GetAllBank(ctx *gin.Context) {
	banks, err := sc.seederService.GetAllBank()
	if err != nil {
		res := utils.BuildResponseFailed("Gagal Mendapatkan List Bank", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess("Berhasil Mendapatkan List Bank", banks)
	ctx.JSON(http.StatusOK, res)
}

func (sc *seederController) GetBankByID(ctx *gin.Context) {
	id := ctx.Param("id")
	uintID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		res := utils.BuildResponseFailed("Gagal Mendapatkan Bank", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	bank, err := sc.seederService.GetBankByID(uint(uintID))
	if err != nil {
		res := utils.BuildResponseFailed("Gagal Mendapatkan Bank", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess("Berhasil Mendapatkan Bank", bank)
	ctx.JSON(http.StatusOK, res)
}