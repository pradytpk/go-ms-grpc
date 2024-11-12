package catalog

import (
	"context"
	"encoding/json"
	"errors"

	elastic "gopkg.in/olivere/elastic.v5"
)

var (
	ErrNotFound = errors.New("Entity not found")
)

const (
	catalog = "catalog"
	product = "product"
)

type Repository interface {
	Close()
	PutProduct(ctx context.Context, p Product) error
	GetProductsByID(ctx context.Context, id string) (*Product, error)
	ListsProducts(ctx context.Context, skip, take uint64) ([]Product, error)
	ListsProductsWithIDs(ctx context.Context, ids []string) ([]Product, error)
	SearchProducts(ctx context.Context, qurey string, skip uint64, take uint64) ([]Product, error)
}

type elasticRepository struct {
	client *elastic.Client
}

type productDocument struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

func NewElasticRepository(url string) (Repository, error) {
	client, err := elastic.NewClient(
		elastic.SetURL(url),
		elastic.SetSniff(false),
	)
	if err != nil {
		return nil, err
	}
	return &elasticRepository{client: client}, nil
}

// Close implements Repository.
func (e *elasticRepository) Close() {
	panic("unimplemented")
}

// GetProductsByID implements Repository.
func (e *elasticRepository) GetProductsByID(ctx context.Context, id string) (*Product, error) {
	res, err := e.client.Get().
		Index(catalog).
		Type(product).
		Id(id).
		Do(ctx)
	if err != nil {
		return nil, err
	}
	if !res.Found {
		return nil, ErrNotFound
	}
	p := productDocument{}
	if err = json.Unmarshal(*res.Source, &p); err != nil {
		return nil, err
	}

	return &Product{
		ID:          id,
		Name:        p.Name,
		Description: p.Description,
		Price:       p.Price,
	}, nil
}

// ListsProducts implements Repository.
func (e *elasticRepository) ListsProducts(ctx context.Context, skip uint64, take uint64) ([]Product, error) {
	panic("unimplemented")
}

// ListsProductsWithIDs implements Repository.
func (e *elasticRepository) ListsProductsWithIDs(ctx context.Context, ids []string) ([]Product, error) {
	panic("unimplemented")
}

// PutProduct implements Repository.
func (e *elasticRepository) PutProduct(ctx context.Context, p Product) error {
	_, err := e.client.Index().
		Index(catalog).
		Type(product).
		Id(p.ID).
		BodyJson(productDocument{
			Name:        p.Name,
			Description: p.Description,
			Price:       p.Price,
		}).
		Do(ctx)
	return err
}

// SearchProducts implements Repository.
func (e *elasticRepository) SearchProducts(ctx context.Context, qurey string, skip uint64, take uint64) ([]Product, error) {
	panic("unimplemented")
}
