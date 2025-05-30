# Script to start the API Gateway service

Write-Host "Starting API Gateway..." -ForegroundColor Green

# Set the current directory to the API Gateway directory
Set-Location -Path $PSScriptRoot

# Check if .env file exists
if (-Not (Test-Path ".env")) {
    Write-Host "Warning: .env file not found. Using default configuration." -ForegroundColor Yellow
}

# Build and run the API Gateway
Write-Host "Building and running API Gateway..." -ForegroundColor Cyan
go run main.go

# Note: This script will block until the API Gateway is terminated
