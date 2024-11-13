package catalog

import (
	"context"
	"encoding/json"
	"errors"
	"log"

	elastic "gopkg.in/olivere/elastic.v5"
)

var (
	ErrNotFound = errors.New("entity not found")
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
	SearchProducts(ctx context.Context, query string, skip uint64, take uint64) ([]Product, error)
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
	res, err := e.client.Search().Index(catalog).Type(product).Query(elastic.NewMatchAllQuery()).From(int(skip)).Size(int(take)).Do(ctx)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	products := []Product{}

	for _, hit := range res.Hits.Hits {
		p := productDocument{}
		if err = json.Unmarshal(*hit.Source, &p); err == nil {
			products = append(products, Product{
				ID:          hit.Id,
				Name:        p.Name,
				Description: p.Description,
				Price:       p.Price,
			})
		}
	}
	return products, nil
}

// ListsProductsWithIDs implements Repository.
func (e *elasticRepository) ListsProductsWithIDs(ctx context.Context, ids []string) ([]Product, error) {
	items := []*elastic.MultiGetItem{}
	for _, id := range ids {
		items = append(items,
			elastic.NewMultiGetItem().Index(catalog).Type(product).Id(id))
	}
	res, err := e.client.MultiGet().Add(items...).Do(ctx)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	products := []Product{}
	for _, doc := range res.Docs {
		p := productDocument{}
		if err = json.Unmarshal(*doc.Source, &p); err == nil {
			products = append(products, Product{
				ID:          doc.Id,
				Name:        p.Name,
				Description: p.Description,
				Price:       p.Price,
			})
		}
	}
	return products, nil
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
func (e *elasticRepository) SearchProducts(ctx context.Context, query string, skip uint64, take uint64) ([]Product, error) {
	res, err := e.client.Search().
		Index(catalog).
		Type(product).
		Query(elastic.NewMultiMatchQuery(query, "name", "description")).
		From(int(skip)).Size(int(take)).
		Do(ctx)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	products := []Product{}
	for _, hits := range res.Hits.Hits {
		p := productDocument{}
		if err = json.Unmarshal(*hits.Source, &p); err == nil {
			products = append(products, Product{
				ID:          hits.Id,
				Name:        p.Name,
				Description: p.Description,
				Price:       p.Price,
			})
		}
	}
	return products, nil
}
