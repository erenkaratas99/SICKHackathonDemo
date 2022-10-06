package handlers

import (
	"SICKHackathon/chatbotService/src/internal/entities"
	"SICKHackathon/chatbotService/src/internal/repositories"
	"SICKHackathon/chatbotService/src/internal/services"
	"SICKHackathon/shared/types"
	"github.com/erenkaratas99/COApiCore/pkg/customErrors"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	repo    *repositories.Repository
	echo    *echo.Echo
	service *services.Service
}

func NewHandler(repo *repositories.Repository, service *services.Service, echo *echo.Echo) *Handler {
	return &Handler{repo: repo, service: service, echo: echo}
}

func (h *Handler) InitEndpoints() {
	e := h.echo
	g := e.Group("/ask")
	g.POST("", h.SendMsg)
	//g.GET("/last", h.GetLastAsked)
}

func (h *Handler) SendMsg(c echo.Context) error {
	messageReq := entities.MsgRequestModel{}
	err := c.Bind(&messageReq)
	if err != nil {
		return customErrors.BindErr
	}
	msgComm := types.MsgCommModel{
		MsgBody: messageReq.MsgBody,
		Name:    messageReq.Name,
		SName:   messageReq.SName,
	}
	err = h.service.SendMessageService(&msgComm)
	return err
}

//func (h *Handler) GetLastAsked(c echo.Context) error {
//	return nil
//}
