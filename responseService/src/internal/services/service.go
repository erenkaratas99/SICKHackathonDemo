package services

import "SICKHackathon/responseService/src/internal/repositories"

type Service struct {
	repo *repositories.Repository
}

func NewService(r *repositories.Repository) *Service {
	return &Service{repo: r}
}
