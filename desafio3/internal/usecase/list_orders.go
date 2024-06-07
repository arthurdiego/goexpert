package usecase

import "github.com/arthurdiego/goexpert/desafio3/internal/entity"

type ListOrdersUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewListOrdersUseCase(OrderRepository entity.OrderRepositoryInterface) *ListOrdersUseCase {
	return &ListOrdersUseCase{OrderRepository: OrderRepository}
}

func (l ListOrdersUseCase) Execute() ([]OrderOutputDTO, error) {
	orders, err := l.OrderRepository.List()
	if err != nil {
		return nil, err
	}

	dto := make([]OrderOutputDTO, len(orders))

	for i := range orders {
		dto[i] = OrderOutputDTO{
			ID:         orders[i].ID,
			Price:      orders[i].Price,
			Tax:        orders[i].Tax,
			FinalPrice: orders[i].FinalPrice,
		}
	}

	return dto, nil
}
