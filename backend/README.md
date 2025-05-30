# Car Rental System

A microservices-based car rental system built with Go and gRPC.

## Services

- API Gateway
- User Service
- Inventory Service
- Booking Service
- Statistics Service

## Prerequisites

- Go 1.24.1 or later
- Protocol Buffers compiler (protoc)
- Docker and Docker Compose (for containerization)

## Setup

1. Clone the repository:
```bash
git clone https://github.com/reallvaibhav/car-rental.git
cd car-rental
```

2. Install dependencies:
```bash
go mod tidy
```

3. Generate proto files:
```bash
cd proto
./generate.sh
```

4. Run the services:
```bash
docker-compose up
```

## API Documentation

The API Gateway exposes the following endpoints:

### Authentication
- POST /auth/register - Register a new user
- POST /auth/login - Login user

### Cars
- POST /cars - Create a new car
- GET /cars - List all cars
- GET /cars/:id - Get car details
- PUT /cars/:id - Update car details
- DELETE /cars/:id - Delete a car

### Bookings
- POST /bookings - Create a new booking
- GET /bookings/user - Get user's bookings
- GET /bookings/fleet-owner - Get fleet owner's bookings
- DELETE /bookings/:id - Cancel a booking

### Statistics
- GET /metrics - Get system metrics

## License

MIT 