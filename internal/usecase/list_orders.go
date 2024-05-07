package usecase

import (
	"github.com/FelpsCorrea/MultiServer-Go/internal/entity"
	"github.com/FelpsCorrea/MultiServer-Go/internal/usecase/dto"
)

type ListOrdersUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewListOrdersUseCase(OrderRepository entity.OrderRepositoryInterface) *ListOrdersUseCase {
	return &ListOrdersUseCase{
		OrderRepository: OrderRepository,
	}
}

func (c *ListOrdersUseCase) Execute() ([]dto.OrderOutputDTO, error) {
	orders, err := c.OrderRepository.ListOrders()

	if err != nil {
		return nil, err
	}

	var ordersDTO []dto.OrderOutputDTO
	for _, order := range orders {
		ordersDTO = append(ordersDTO, dto.OrderOutputDTO{
			ID:         order.ID,
			Price:      order.Price,
			Tax:        order.Tax,
			FinalPrice: order.Price + order.Tax,
		})
	}

	return ordersDTO, nil

}
