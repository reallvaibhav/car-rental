#!/usr/bin/env pwsh

# This script tests the API Gateway endpoints using PowerShell

# Set base URL for API Gateway
$baseUrl = "http://localhost:8080"
$token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySWQiOiJ0ZXN0LXVzZXItaWQiLCJlbWFpbCI6InRlc3RAZXhhbXBsZS5jb20iLCJyb2xlIjoidXNlciJ9.aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"

Write-Host "========================================"
Write-Host "API Gateway Test Script"
Write-Host "========================================"

function Test-Endpoint {
    param (
        [string]$Method,
        [string]$Endpoint,
        [string]$Description,
        [object]$Body = $null,
        [bool]$RequiresAuth = $true
    )
    
    Write-Host "`nTesting: $Description"
    Write-Host "Endpoint: $Method $Endpoint"
    
    $headers = @{
        "Content-Type" = "application/json"
    }
    
    if ($RequiresAuth) {
        $headers["Authorization"] = "Bearer $token"
    }
    
    $params = @{
        Method = $Method
        Uri = "$baseUrl$Endpoint"
        Headers = $headers
    }
    
    if ($Body -ne $null -and $Method -ne "GET") {
        $params["Body"] = ($Body | ConvertTo-Json)
    }
    
    try {
        $response = Invoke-RestMethod @params -ErrorAction Stop
        Write-Host "Status: Success" -ForegroundColor Green
        Write-Host "Response: " -NoNewline
        $response | ConvertTo-Json -Depth 3 | Write-Host
    }
    catch {
        Write-Host "Status: Failed" -ForegroundColor Red
        Write-Host "Error: $($_.Exception.Message)"
        
        if ($_.Exception.Response) {
            $result = $_.Exception.Response.GetResponseStream()
            $reader = New-Object System.IO.StreamReader($result)
            $reader.BaseStream.Position = 0
            $reader.DiscardBufferedData()
            $responseBody = $reader.ReadToEnd()
            Write-Host "Response Body: $responseBody"
        }
    }
    
    Write-Host "----------------------------------------"
}

# Test Authentication Endpoints
Test-Endpoint -Method "POST" -Endpoint "/auth/register" -Description "Register a new user" -RequiresAuth $false -Body @{
    username = "testuser"
    email = "test@example.com"
    password = "password123"
}

Test-Endpoint -Method "POST" -Endpoint "/auth/login" -Description "Login user" -RequiresAuth $false -Body @{
    email = "test@example.com"
    password = "password123"
}

# Test Car Endpoints
Test-Endpoint -Method "GET" -Endpoint "/cars" -Description "List all cars"

Test-Endpoint -Method "POST" -Endpoint "/cars" -Description "Create a new car" -Body @{
    make = "Toyota"
    model = "Camry"
    year = 2022
    price_per_day = 50.0
    color = "Blue"
    category = "Sedan"
    available = $true
    owner_id = "test-user-id"
    location = "New York"
}

Test-Endpoint -Method "GET" -Endpoint "/cars/car123" -Description "Get car details"

Test-Endpoint -Method "PUT" -Endpoint "/cars/car123" -Description "Update car details" -Body @{
    price_per_day = 55.0
    available = $false
}

# Test Booking Endpoints
Test-Endpoint -Method "POST" -Endpoint "/bookings" -Description "Create a booking" -Body @{
    userId = "test-user-id"
    bookings = @(
        @{
            carId = "car123"
            startDate = "2025-06-01"
            endDate = "2025-06-05"
            pricePerDay = 50.0
            totalDays = 5
        }
    )
}

Test-Endpoint -Method "GET" -Endpoint "/bookings/booking123" -Description "Get booking details"

Test-Endpoint -Method "GET" -Endpoint "/bookings/user?user_id=test-user-id" -Description "Get user bookings"

Test-Endpoint -Method "GET" -Endpoint "/bookings/fleet-owner?owner_id=test-user-id" -Description "Get fleet owner bookings"

Test-Endpoint -Method "PUT" -Endpoint "/bookings/booking123/status" -Description "Update booking status" -Body @{
    status = "COMPLETED"
}

Test-Endpoint -Method "DELETE" -Endpoint "/bookings/booking123" -Description "Cancel booking"

# Test Statistics Endpoints
Test-Endpoint -Method "GET" -Endpoint "/statistics/bookings?time_range=monthly" -Description "Get booking statistics"

Test-Endpoint -Method "GET" -Endpoint "/statistics/cars?time_range=monthly&category=Sedan" -Description "Get car statistics"

Test-Endpoint -Method "GET" -Endpoint "/statistics/revenue?time_range=yearly" -Description "Get revenue statistics"

Test-Endpoint -Method "GET" -Endpoint "/statistics/popular-locations?limit=5" -Description "Get popular locations"

Test-Endpoint -Method "GET" -Endpoint "/statistics/users?time_range=monthly" -Description "Get user statistics"

# Test Metrics Endpoint
Test-Endpoint -Method "GET" -Endpoint "/metrics" -Description "Get API metrics"

Write-Host "`nTest script completed!"
