package main

import "context"

type MutationResolver struct {
	server *Server
}

// CreateAccount
func (r *MutationResolver) createAccount(ctx context.Context, in AccountInput) (*Account, error) {

}

// CreateProduct
func (r *MutationResolver) createProduct(ctx context.Context, in ProductInput) (*Product, error) {

}

// CreateOrder
func (r *MutationResolver) createOrder(ctx context.Context, in OrderInput) (*Order, error) {

}
