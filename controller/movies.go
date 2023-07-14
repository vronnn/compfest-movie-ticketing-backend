package controller

import (
	"net/http"

	"github.com/Caknoooo/golang-clean_template/dto"
	"github.com/Caknoooo/golang-clean_template/services"
	"github.com/Caknoooo/golang-clean_template/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type MovieController interface {
	CreateMovie(ctx *gin.Context)
	GetAllMovie(ctx *gin.Context)
	GetMovieByID(ctx *gin.Context)
}

type movieController struct {
	jwtService services.JWTService
	movieService services.MovieService
	userService services.UserService
}

func NewMovieController(ms services.MovieService, jwt services.JWTService, us services.UserService) MovieController {
	return &movieController{
		jwtService: jwt,
		movieService: ms,
		userService: us,
	}
}

func (mc *movieController) CreateMovie(ctx *gin.Context) {
	token := ctx.MustGet("token").(string)
	userID, err := mc.jwtService.GetUserIDByToken(token)
	if err != nil {
		res := utils.BuildResponseFailed("Gagal Mendapatkan User", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	data, err := mc.userService.GetUserByID(ctx.Request.Context(), userID)
	if err != nil {
		res := utils.BuildResponseFailed("Gagal Mendapatkan User", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	if data.Role != "Admin" {
		res := utils.BuildResponseFailed("Anda Tidak Memiliki Akses", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusUnauthorized, res)
		return
	}

	var movie dto.MovieCreateDTO
	if err := ctx.ShouldBindJSON(&movie); err != nil {
		res := utils.BuildResponseFailed("Gagal Mendapatkan Request", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	result, err := mc.movieService.CreateMovies(ctx.Request.Context(), movie)
	if err != nil {
		res := utils.BuildResponseFailed("Gagal Menambahkan Movie", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess("Berhasil Menambahkan Movie", result)
	ctx.JSON(http.StatusOK, res)
}

func (mc *movieController) GetAllMovie(ctx *gin.Context) {
	result, err := mc.movieService.GetAllMovies(ctx.Request.Context())
	if err != nil {
		res := utils.BuildResponseFailed("Gagal Mendapatkan List Movie", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := utils.BuildResponseSuccess("Berhasil Mendapatkan List Movie", result)
	ctx.JSON(http.StatusOK, res)
}

func (mc *movieController) GetMovieByID(ctx *gin.Context) {
	id := ctx.Param("id")
	uuid, err := uuid.Parse(id)
	if err != nil {
		res := utils.BuildResponseFailed("Gagal Parse ID", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	result, err := mc.movieService.GetMovieByID(ctx.Request.Context(), uuid)
	if err != nil {
		res := utils.BuildResponseFailed("Gagal Mendapatkan Movie", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess("Berhasil Mendapatkan Movie", result)
	ctx.JSON(http.StatusOK, res)
}