package usecase

import (
	dto "github.com/wrferreira1003/Desafio-Clean-Architecture/internal/Dto"
	repository "github.com/wrferreira1003/Desafio-Clean-Architecture/internal/domain/repository"
)

type ListOrderUseCase struct {
	OrderRepository repository.OrderRepositoryInterface
}

func NewListOrderUseCase(orderRepository repository.OrderRepositoryInterface) *ListOrderUseCase {
	return &ListOrderUseCase{
		OrderRepository: orderRepository,
	}
}

func (l *ListOrderUseCase) Execute() ([]dto.OrderOutputDTO, error) {
	orders, err := l.OrderRepository.FindAll()
	if err != nil {
		return nil, err
	}

	var ordersOutput []dto.OrderOutputDTO
	for _, order := range orders {
		ordersOutput = append(ordersOutput, dto.OrderOutputDTO{
			ID:    order.ID,
			Price: order.Price,
			Tax:   order.Tax,
		})
	}
	return ordersOutput, nil
}
