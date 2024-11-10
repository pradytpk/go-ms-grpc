package main

import "context"

type queryResolver struct {
	server *Server
}

// Accounts

func (r *queryResolver) Accounts(ctx context.Context, pagination *PaginationInput, id *string) ([]*Account, error) {

	return nil, nil
}

// Products
func (r *queryResolver) Products(ctx context.Context, pagination *PaginationInput, query, id *string) ([]*Product, error) {

	return nil, nil
}
