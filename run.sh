#!/bin/bash

# Run the product service server
go run services/product/server/main.go
SERVER_PID=$!

# Run the product service client
go run services/product/main.go

kill $SERVER_PID