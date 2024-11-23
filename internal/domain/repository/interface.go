package repository

import "github.com/wrferreira1003/Desafio-Clean-Architecture/internal/domain/entities"

type OrderRepositoryInterface interface {
	Save(order *entities.Order) error   // Método para salvar o pedido no banco de dados
	FindAll() ([]entities.Order, error) // Método para listar todos os pedidos
}
