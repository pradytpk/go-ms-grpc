package catalog

import (
	"context"

	"github.com/pradytpk/go-ms-grpc/catalog/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	conn    *grpc.ClientConn
	service pb.CatalogServiceClient
}

func NewClient(url string) (*Client, error) {
	conn, err := grpc.NewClient(url, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	c := pb.NewCatalogServiceClient(conn)
	return &Client{conn: conn, service: c}, nil
}

func (c *Client) Close() {
	c.conn.Close()
}

func (c *Client) PostProduct(ctx context.Context, name, description string, price float64) (*Product, error) {
	r, err := c.service.PostProduct(ctx, &pb.PostProductRequest{
		Name:        name,
		Description: description,
		Price:       price,
	})
	if err != nil {
		return nil, err
	}
	return &Product{
		ID:          r.Product.Id,
		Name:        r.Product.Name,
		Description: r.Product.Description,
		Price:       r.Product.Price,
	}, nil

}

func (c *Client) GetProduct(ctx context.Context, id string) (*Product, error) {
	r, err := c.service.GetProduct(ctx, &pb.GetProductRequest{Id: id})
	if err != nil {
		return nil, err
	}
	return &Product{
		ID:          r.Product.Id,
		Name:        r.Product.Name,
		Description: r.Product.Description,
		Price:       r.Product.Price,
	}, nil

}

func (c *Client) GetProducts(ctx context.Context, ids []string, query string, skip, take uint64) ([]Product, error) {
	r, err := c.service.GetProducts(ctx,
		&pb.GetProductsRequest{
			Ids:   ids,
			Skip:  skip,
			Take:  take,
			Query: query,
		})
	if err != nil {
		return nil, err
	}
	products := []Product{}
	for _, v := range r.Products {
		products = append(products, Product{
			ID:          v.Id,
			Name:        v.Name,
			Description: v.Description,
			Price:       v.Price,
		})
	}
	return products, nil
}
