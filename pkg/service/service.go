package service

import "github.com/TimmyTurner98/consult/pkg/modules"

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

// Имитируем создание клиента
func (s *Service) CreateClient(client modules.Client) (modules.Client, error) {
	// Здесь мы просто добавляем фиктивный ID и возвращаем клиента как будто он был создан
	client.Id = 1      // Имитируем ID, который бы присвоил сервер БД
	return client, nil // Возвращаем клиента как успешно созданного
}
