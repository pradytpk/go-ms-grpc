package catalog

import (
	"context"

	"github.com/segmentio/ksuid"
)

type Product struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

type Service interface {
	PostProduct(ctx context.Context, name, description string, price float64) (*Product, error)
	GetProduct(ctx context.Context, id string) (*Product, error)
	GetProducts(ctx context.Context, skip uint64, take uint64) ([]Product, error)
	GetProductsByID(ctx context.Context, ids []string) ([]Product, error)
	SearchProducts(ctx context.Context, query string, skip uint64, take uint64) ([]Product, error)
}

type catalogService struct {
	repository Repository
}

// GetProduct implements Service.
func (c *catalogService) GetProduct(ctx context.Context, id string) (*Product, error) {
	panic("unimplemented")
}

// GetProducts implements Service.
func (c *catalogService) GetProducts(ctx context.Context, skip uint64, take uint64) ([]Product, error) {
	if take > 100 || (skip == 0 && take == 0) {
		take = 100
	}
	return c.repository.ListProducts(ctx, skip, take)
}

// GetProductsByID implements Service.
func (c *catalogService) GetProductsByID(ctx context.Context, ids []string) ([]Product, error) {
	return c.repository.ListProductsWithIDs(ctx, ids)
}

// PostProduct implements Service.
func (c *catalogService) PostProduct(ctx context.Context, name string, description string, price float64) (*Product, error) {
	p := &Product{
		Name:        name,
		Description: description,
		Price:       price,
		ID:          ksuid.New().String(),
	}
	if err := c.repository.PutProduct(ctx, *p); err != nil {
		return nil, err
	}
	return p, nil
}

// SearchProducts implements Service.
func (c *catalogService) SearchProducts(ctx context.Context, query string, skip uint64, take uint64) ([]Product, error) {
	if take > 100 || (skip == 0 && take == 0) {
		take = 100
	}
	return c.repository.SearchProducts(ctx, query, skip, take)
}

func NewService(r Repository) Service {
	return &catalogService{r}
}
