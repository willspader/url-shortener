package service

import "url-shortener/internal/repository"

type Service struct {
	repository *repository.Repository
}

func CreateService(repository *repository.Repository) *Service {
	return &Service{repository: repository}
}
