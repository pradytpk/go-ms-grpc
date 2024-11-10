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


account 
mutation or query → client →(gRpc)→ server → service → repository → database

 go:generate protoc ./account.proto --go_out=plugins=grpc:./pb

 export PATH=$PATH:$(go env GOPATH)/bin  
 protoc -I=. --go_out=. --go-grpc_out=. account.proto

export DATABASE_URL=postgres://admin:adminpassword@localhost/social?sslmode=disable
