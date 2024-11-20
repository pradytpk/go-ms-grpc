#  Go Microservice 
## Tech Stacks
- gRPC
- GO
- Docker
- Graphql
- Postgres
- Elasticsearch!

```
go get github.com/99designs/gqlgen
```


## Account code flow
mutation or query → client →(gRpc)→ server → service → repository → database

```
go:generate protoc ./account.proto --go_out=plugins=grpc:./pb

go:generate protoc ./catalog.proto --go_out=plugins=grpc:./pb

export PATH=$PATH:$(go env GOPATH)/bin  

protoc -I=. --go_out=. --go-grpc_out=. account.proto

protoc -I=. --go_out=. --go-grpc_out=. catalog.proto

protoc -I=. --go_out=. --go-grpc_out=. order.proto

export DATABASE_URL=postgres://admin:adminpassword@localhost/social?sslmode=disable
```

```
curl -X PUT "localhost:9200/products" -H 'Content-Type: application/json' -d '{
  "mappings": {
    "properties": {
      "id": {
        "type": "keyword"
      },
      "name": {
        "type": "text"
      },
      "description": {
        "type": "text"
      },
      "price": {
        "type": "float"
      }
    }
  }
}'

```