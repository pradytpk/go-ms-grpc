package order

import (
	"context"
	"time"

	"github.com/segmentio/ksuid"
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
	return o.repository.GetOrdersForAccount(ctx, accountID)
}

// PostOrder implements Service.
func (o *orderService) PostOrder(ctx context.Context, accountID string, products []OrderedProduct) (*Order, error) {
	or := &Order{
		ID:        ksuid.New().String(),
		CreatedAt: time.Now().UTC(),
		AccountID: accountID,
		Products:  products,
	}
	or.TotalPrice = 0.0
	for _, v := range products {
		or.TotalPrice += v.Price * float64(v.Quantity)
	}
	err := o.repository.PutOrder(ctx, *or)
	if err != nil {
		return nil, err
	}
	return or, nil
}

func NewService(r Repository) Service {
	return &orderService{r}
}
