package order

import (
	"context"
	"time"
)

type Order struct {
	ID         string
	CreatedAt  time.Time
	TotalPrice float64
	AccountID  string
	Products   []OrderedProduct
}

type OrderedProduct struct {
	ID          string
	Name        string
	Description string
	Price       float64
	Quantity    uint32
}

type Service interface {
	PostOrder(ctx context.Context, accountID string, products []OrderedProduct) (*Order, error)
	GetOrdersForAccount(ctx context.Context, accountID string) ([]Order, error)
}

type orderService struct {
	repository Repository
}

// GetOrdersForAccount implements Service.
func (o *orderService) GetOrdersForAccount(ctx context.Context, accountID string) ([]Order, error) {
	panic("unimplemented")
}

// PostOrder implements Service.
func (o *orderService) PostOrder(ctx context.Context, accountID string, products []OrderedProduct) (*Order, error) {
	panic("unimplemented")
}

func NewService(r Repository) Service {
	return &orderService{r}
}
