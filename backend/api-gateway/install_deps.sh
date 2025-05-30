#!/bin/bash

# Script to install required dependencies for the API Gateway

echo "Installing dependencies for API Gateway..."

# Install godotenv for environment variable management
go get github.com/joho/godotenv

# Install gin framework
go get github.com/gin-gonic/gin

# Update dependencies
go mod tidy

echo "Dependencies installed successfully!"
