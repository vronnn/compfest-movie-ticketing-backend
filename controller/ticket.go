package controller

import (
	"net/http"

	"github.com/Caknoooo/golang-clean_template/dto"
	"github.com/Caknoooo/golang-clean_template/services"
	"github.com/Caknoooo/golang-clean_template/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type TicketController interface {
	CreateTicket(ctx *gin.Context)
	GetAllTicketMovie(ctx *gin.Context)
	GetTicketUser(ctx *gin.Context)
	SendSchedule(ctx *gin.Context)
}

type ticketController struct {
	ticketService services.TicketService
	jwtService services.JWTService
	Jam string
	Studio string
}

func NewTicketController(ticketService services.TicketService, jwtService services.JWTService) TicketController {
	return &ticketController{
		ticketService: ticketService,
		jwtService: jwtService,
		Jam: "0",
		Studio: "0",
	}
}

func (tc *ticketController) CreateTicket(ctx *gin.Context) {
	token := ctx.MustGet("token").(string)
	userID, err := tc.jwtService.GetUserIDByToken(token)
	if err != nil {
		res := utils.BuildResponseFailed("Gagal Mendapatkan User", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	var ticketDTO dto.TicketCreateDTO
	if err := ctx.ShouldBind(&ticketDTO); err != nil {
		res := utils.BuildResponseFailed("Gagal Mendapatkan Request", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	if ticketDTO.Amount > 6 {
		res := utils.BuildResponseFailed("User Hanya Dapat Melakukan Pembelian Maksimal 6 Ticket", "", utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	result, err := tc.ticketService.CreateTicket(ctx.Request.Context(), ticketDTO, userID)
	if err != nil {
		res := utils.BuildResponseFailed("Gagal Melakukan Pembelian Ticketing", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess("Berhasil Melakukan Pembelian Ticket", result)
	ctx.JSON(http.StatusOK, res)
}

func (tc *ticketController) GetAllTicketMovie(ctx *gin.Context) {
	id := ctx.Param("movie_id")
	uuid, err := uuid.Parse(id)
	if err != nil {
		res := utils.BuildResponseFailed("Gagal Mendapatkan Movie ID", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	var ticket dto.GetAllTicketMovie
	ticket.MovieID = uuid

	result, err := tc.ticketService.GetAllTicketMovie(ctx.Request.Context(), ticket, tc.Jam, tc.Studio)
	if err != nil {
		res := utils.BuildResponseFailed("Gagal Mendapatkan Ticket", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess("Berhasil Mendapatkan Ticket", result)
	ctx.JSON(http.StatusOK, res)
}

func (tc *ticketController) GetTicketUser(ctx *gin.Context) {
	token := ctx.MustGet("token").(string)
	userID, err := tc.jwtService.GetUserIDByToken(token)
	if err != nil {
		res := utils.BuildResponseFailed("Gagal Mendapatkan User", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	result, err := tc.ticketService.GetTicketUser(ctx.Request.Context(), userID)
	if err != nil {
		res := utils.BuildResponseFailed("Gagal Mendapatkan Ticket", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess("Berhasil Mendapatkan Ticket", result)
	ctx.JSON(http.StatusOK, res)
}

func (tc *ticketController) SendSchedule(ctx *gin.Context) {
	var schedule dto.SendScheduleMovie
	if err := ctx.ShouldBindJSON(&schedule); err != nil {
		res := utils.BuildResponseFailed("Gagal Mendapatkan Request", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	tc.Jam = schedule.Jam
	tc.Studio = schedule.Studio

	res := utils.BuildResponseSuccess("Berhasil Mengirim Schedule", utils.EmptyObj{})
	ctx.JSON(http.StatusOK, res)
}