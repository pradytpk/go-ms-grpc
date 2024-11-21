#  Go Microservice with GraphQL API(gRPC)
This project demonstrates a microservices architecture using gRPC for inter-service communication and GraphQL as the API gateway. It includes services for account management, product catalog, and order processing.

## Project Structure

The project consists of the following main components:

- Account Service
- Catalog Service
- Order Service
- GraphQL API Gateway

Each service has its own database:
- Account and Order services use PostgreSQL
- Catalog service uses Elasticsearch

## Tech Stacks
- gRPC (v1.68.0)
- Golang (golang:1.23.3)
- Docker (26.0.0)
- Graphql
- Postgres (postgres:10.3)
- Elasticsearch (elasticsearch:8.6.0)

## Code flow
mutation or query → client →(gRpc)→ server → service → repository → database

## Frequently used Commmands
```
protoc -I=. --go_out=. --go-grpc_out=. account.proto

protoc -I=. --go_out=. --go-grpc_out=. catalog.proto

protoc -I=. --go_out=. --go-grpc_out=. order.proto

```
## Getting Started

1. Clone the repository:
   ```
   git clone <repository-url>
   cd <project-directory>
   ```

2. Start the services using Docker Compose:
   ```
   docker-compose up -d --build
   ```

3. Access the GraphQL playground at `http://localhost:8000/playground`

## GraphQL API 

### Query Accounts

```graphql
query {
  accounts {
    id
    name
  }
}
```

### Create an Account

```graphql
mutation {
  createAccount(account: {name: "New Account"}) {
    id
    name
  }
}
```

### Query Products

```graphql
query {
  products {
    id
    name
    price
  }
}
```

### Create a Product

```graphql
mutation {
  createProduct(product: {name: "New Product", description: "A new product", price: 19.99}) {
    id
    name
    price
  }
}
```

### Create an Order

```graphql
mutation {
  createOrder(order: {accountId: "account_id", products: [{id: "product_id", quantity: 2}]}) {
    id
    totalPrice
    products {
      name
      quantity
    }
  }
}
```

### Query Account with Orders

```graphql
query {
  accounts(id: "account_id") {
    name
    orders {
      id
      createdAt
      totalPrice
      products {
        name
        quantity
        price
      }
    }
  }
}
```

## Advanced Queries

### Pagination and Filtering

```graphql
query {
  products(pagination: {skip: 0, take: 5}, query: "search_term") {
    id
    name
    description
    price
  }
}
```

### Calculate Total Spent by an Account

```graphql
query {
  accounts(id: "account_id") {
    name
    orders {
      totalPrice
    }
  }
}
```