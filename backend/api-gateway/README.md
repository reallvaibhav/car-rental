# Car Rental API Gateway

The API Gateway serves as a unified entry point for all client requests in the Car Rental system, routing them to the appropriate microservices.

## Features

- **Single Entry Point**: Provides a unified API for the frontend
- **Authentication**: Handles JWT-based authentication and authorization
- **Request Routing**: Routes requests to the appropriate microservices
- **Error Handling**: Standardized error responses
- **Metrics**: Basic monitoring of API usage

## Architecture

The API Gateway connects to the following microservices:
- User Service: Handles user authentication and management
- Inventory Service: Manages car inventory and availability
- Booking Service: Processes car bookings and reservations
- Statistics Service: Provides system-wide statistics and analytics

## Setup and Installation

1. Configure the environment variables in the `.env` file:
   ```
   # Service URLs
   USER_SERVICE_ADDR=localhost:50051
   INVENTORY_SERVICE_ADDR=localhost:50052
   BOOKING_SERVICE_ADDR=localhost:50053
   STATISTICS_SERVICE_ADDR=localhost:50054

   # JWT Configuration
   JWT_SECRET=your-secret-key
   JWT_EXPIRATION=24h

   # Server Configuration
   SERVER_PORT=8080
   SERVER_HOST=0.0.0.0
   DEBUG_MODE=true

   # CORS Configuration
   CORS_ALLOWED_ORIGINS=*
   ```

2. Install dependencies:
   ```
   # Windows
   .\install_deps.ps1

   # Linux/Mac
   ./install_deps.sh
   ```

3. Start the service:
   ```
   # Windows
   .\start.ps1

   # Linux/Mac
   go run main.go
   ```

4. Test the API endpoints:
   ```
   # Windows
   .\test_api.ps1
   ```

## API Endpoints

### Authentication
- `POST /auth/register`: Register a new user
- `POST /auth/login`: Login and get JWT token

### Cars
- `GET /cars`: Get list of available cars
- `GET /cars/:id`: Get details of a specific car
- `POST /cars`: Add a new car (fleet owner only)
- `PUT /cars/:id`: Update car information
- `DELETE /cars/:id`: Remove a car from inventory

### Bookings
- `POST /bookings`: Create a new booking
- `GET /bookings/:id`: Get details of a specific booking
- `GET /bookings/user`: Get bookings for current user
- `GET /bookings/fleet-owner`: Get bookings for cars owned by fleet owner
- `PUT /bookings/:id/status`: Update booking status
- `DELETE /bookings/:id`: Cancel a booking

### Statistics
- `GET /statistics/bookings`: Get booking statistics with time range filtering
- `GET /statistics/cars`: Get car statistics and utilization data
- `GET /statistics/revenue`: Get revenue statistics and trends
- `GET /statistics/popular-locations`: Get most popular rental locations
- `GET /statistics/users`: Get user activity statistics

### System
- `GET /metrics`: Get API usage metrics and service status

## Docker Deployment

The API Gateway can be deployed using Docker:

```
docker build -t car-rental-api-gateway .
docker run -p 8080:8080 --env-file .env car-rental-api-gateway
```

## Troubleshooting

If you encounter issues connecting to the microservices, check:
1. The service addresses in your `.env` file
2. That all microservices are running
3. Network connectivity between containers if using Docker

## Error Handling

The API Gateway provides standardized error responses in the following format:

```json
{
  "error": "Detailed error message"
}
```

Common HTTP status codes:
- 200: Success
- 400: Bad Request - Invalid input or parameters
- 401: Unauthorized - Authentication required
- 403: Forbidden - Insufficient permissions
- 404: Not Found - Resource doesn't exist
- 500: Internal Server Error - Server-side error

The gateway includes comprehensive error handling middleware that catches panics and returns proper error responses.

## Authentication

All endpoints except for `/auth/register` and `/auth/login` require authentication using JWT tokens.
Include an Authorization header with a Bearer token:

```
Authorization: Bearer <jwt_token>
```

## Development

To add new API endpoints, update the following files:
1. `main.go`: Add the route
2. `handlers/handler.go`: Implement the handler method
3. Ensure proper authentication middleware is applied
