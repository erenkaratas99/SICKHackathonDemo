package handlers

import (
	"SICKHackathon/SAPService/src/internal/repositories"
	"SICKHackathon/SAPService/src/internal/services"
	"SICKHackathon/shared/types"
	"fmt"
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
	g := e.Group("/sap-service")
	g.POST("/pushmsg", h.TakePushedMSG)
}

func (h *Handler) TakePushedMSG(c echo.Context) error {
	cbResp := types.MsgCommModel{}
	err := c.Bind(&cbResp)
	if err != nil {
		fmt.Println("err : ", err, "\n")
		return customErrors.BindErr
	}
	csv, err := h.service.PreNLPService(&cbResp)
	if err != nil {
		return err
	}

	return c.JSON(200, csv)
}
