# Use the updated Go version 1.22.7 with Alpine
FROM golang:1.23.3-alpine3.20 AS build

# Install necessary tools
RUN apk --no-cache add gcc g++ make ca-certificates

# Set the working directory
WORKDIR /go/src/github.com/pradytpk/go-ms-grpc

# Copy dependency files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY vendor vendor
COPY account account
COPY catalog catalog
COPY order order
COPY graphql graphql

# Build the application targeting the graphql module
RUN go build -mod=vendor -o /go/bin/app ./graphql

# Use a minimal Alpine image for the runtime environment
FROM alpine:3.11

# Set the working directory
WORKDIR /usr/bin

# Copy the built application from the build stage
COPY --from=build /go/bin .

# Expose the application's port
EXPOSE 8080

# Set the command to run the application
CMD ["app"]
