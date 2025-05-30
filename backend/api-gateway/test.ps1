# Test script for the API Gateway endpoints

Write-Host "Testing API Gateway endpoints..." -ForegroundColor Green

$apiBaseUrl = "http://localhost:8080"
$testToken = "Bearer test-token" # This should be replaced with a real token in production

# Function to make API requests
function Invoke-ApiRequest {
    param (
        [string]$Method,
        [string]$Endpoint,
        [string]$Body = "",
        [bool]$UseAuth = $true
    )
    
    $headers = @{
        "Content-Type" = "application/json"
    }
    
    if ($UseAuth) {
        $headers["Authorization"] = $testToken
    }
    
    $url = "$apiBaseUrl$Endpoint"
    
    Write-Host "Testing $Method $url" -ForegroundColor Cyan
    
    try {
        if ($Body -eq "") {
            $response = Invoke-RestMethod -Method $Method -Uri $url -Headers $headers -ErrorAction Stop
        } else {
            $response = Invoke-RestMethod -Method $Method -Uri $url -Headers $headers -Body $Body -ErrorAction Stop
        }
        Write-Host "✓ Success" -ForegroundColor Green
        return $response
    } catch {
        Write-Host "✗ Failed: $_" -ForegroundColor Red
        return $null
    }
}

# Test public endpoints
Write-Host "`nTesting public endpoints..." -ForegroundColor Yellow

# Test login endpoint
$loginBody = @{
    email = "test@example.com"
    password = "password123"
} | ConvertTo-Json

$loginResponse = Invoke-ApiRequest -Method "POST" -Endpoint "/auth/login" -Body $loginBody -UseAuth $false

# Test authenticated endpoints
Write-Host "`nTesting authenticated endpoints..." -ForegroundColor Yellow

# Test list cars endpoint
$carsResponse = Invoke-ApiRequest -Method "GET" -Endpoint "/cars"

# Test user bookings endpoint
$userBookingsResponse = Invoke-ApiRequest -Method "GET" -Endpoint "/bookings/user?user_id=123"

# Test create booking
$bookingBody = @{
    userId = "123"
    bookings = @(
        @{
            carId = "car-123"
            startDate = "2023-06-01"
            endDate = "2023-06-05"
            pricePerDay = 50.0
            totalDays = 5
        }
    )
} | ConvertTo-Json

$createBookingResponse = Invoke-ApiRequest -Method "POST" -Endpoint "/bookings" -Body $bookingBody

Write-Host "`nAPI Testing completed!" -ForegroundColor Green
