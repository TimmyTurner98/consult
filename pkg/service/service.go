package service

import (
	"github.com/TimmyTurner98/consult/pkg/modules"
	"github.com/TimmyTurner98/consult/pkg/repository"
)

type Client interface {
	CreateClient(client modules.Client) (modules.Client, error)
}

type Service struct {
	repos *repository.Repository
}

func NewService(repos *repository.Repository) *Service {
	return &Service{repos: repos}
}
