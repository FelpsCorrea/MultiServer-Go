package entity

type OrderRepositoryInterface interface {
	Save(order *Order) error
	GetOrder(ID string) (*Order, error)
	ListOrders() ([]Order, error)
}
