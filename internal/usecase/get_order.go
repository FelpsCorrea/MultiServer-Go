package usecase

import (
	"github.com/FelpsCorrea/MultiServer-Go/internal/entity"
	"github.com/FelpsCorrea/MultiServer-Go/internal/usecase/dto"
)

type GetOrderUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewGetOrderUseCase(OrderRepository entity.OrderRepositoryInterface) *GetOrderUseCase {
	return &GetOrderUseCase{
		OrderRepository: OrderRepository,
	}
}

func (c *GetOrderUseCase) Execute(input dto.GetOrderInputDTO) (dto.OrderOutputDTO, error) {

	order, err := c.OrderRepository.GetOrder(input.ID)

	if err != nil {
		return dto.OrderOutputDTO{}, err
	}

	return dto.OrderOutputDTO{
		ID:         order.ID,
		Price:      order.Price,
		Tax:        order.Tax,
		FinalPrice: order.Price + order.Tax,
	}, nil

}
