package handlers

import (
	"SICKHackathon/responseService/src/internal/kafka/consumer"
	"SICKHackathon/responseService/src/internal/repositories"
	"SICKHackathon/responseService/src/internal/services"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	repo     *repositories.Repository
	echo     *echo.Echo
	service  *services.Service
	consumer *consumer.KafkaConsumerHandler
}

func NewHandler(repo *repositories.Repository, service *services.Service, echo *echo.Echo, consumer *consumer.KafkaConsumerHandler) *Handler {
	return &Handler{repo: repo, service: service, echo: echo, consumer: consumer}
}

func (h *Handler) InitEndpoints() {
	e := h.echo
	g := e.Group("/response")
	g.POST("", h.ConsumeHandler)

}

func (h *Handler) ConsumeHandler(c echo.Context) error {

	return nil
}
