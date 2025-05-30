# Script to install required dependencies for the API Gateway

Write-Host "Installing dependencies for API Gateway..." -ForegroundColor Green

# Set the current directory to the API Gateway directory
Set-Location -Path $PSScriptRoot

# Install godotenv for environment variable management
Write-Host "Installing godotenv..." -ForegroundColor Cyan
go get github.com/joho/godotenv

# Install gin framework
Write-Host "Installing gin framework..." -ForegroundColor Cyan
go get github.com/gin-gonic/gin

# Update dependencies
Write-Host "Updating dependencies..." -ForegroundColor Cyan
go mod tidy

Write-Host "Dependencies installed successfully!" -ForegroundColor Green
