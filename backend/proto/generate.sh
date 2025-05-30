#!/bin/bash

# Generate Go code for all services
protoc --go_out=./booking --go_opt=paths=source_relative \
    --go-grpc_out=./booking --go-grpc_opt=paths=source_relative \
    booking/booking.proto

protoc --go_out=./inventory --go_opt=paths=source_relative \
    --go-grpc_out=./inventory --go-grpc_opt=paths=source_relative \
    inventory/inventory.proto

protoc --go_out=./statistics --go_opt=paths=source_relative \
    --go-grpc_out=./statistics --go-grpc_opt=paths=source_relative \
    statistics/statistics.proto

protoc --go_out=./user --go_opt=paths=source_relative \
    --go-grpc_out=./user --go-grpc_opt=paths=source_relative \
    user/user.proto 