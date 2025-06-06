﻿# Car Rental System

A modern, microservices-based car rental system built with Go, React, and gRPC. The system allows users to rent cars, manage fleet inventory, and track bookings with real-time statistics.

## Project Structure

```
car-rental/
├── backend/              # Go microservices
│   ├── api-gateway/     # API Gateway service
│   ├── booking-service/ # Booking management service
│   ├── email-service/   # Email notification service
│   ├── inventory-service/# Car inventory service
│   ├── nats-service/    # Message queue service
│   ├── statistics-service/# Analytics service
│   └── user-service/    # User authentication service
└── frontend/            # React/TypeScript frontend
    └── src/
        ├── components/  # React components
        ├── pages/       # Page components
        └── services/    # API services
```

## Features

### Backend (Kumar Vaibhav)
- Microservices Architecture with gRPC Communication
- JWT-based Authentication
- Real-time Statistics and Analytics
- Message Queue Integration with NATS
- Docker Containerization
- MongoDB Database Integration
- Comprehensive API Gateway

### Frontend (Harshita)
- Modern React/TypeScript Implementation
- Responsive UI with Tailwind CSS
- User Authentication Flow
- Car Booking Interface
- Fleet Management Dashboard
- Real-time Statistics Display
- Form Validations
- Error Handling

## Prerequisites

- Go 1.24.1 or later
- Node.js 18.x or later
- Docker and Docker Compose
- MongoDB
- Protocol Buffers compiler (protoc)

## Setup Instructions

1. Clone the repository:
```bash
git clone https://github.com/reallvaibhav/car-rental.git
cd car-rental
```

2. Start the Backend Services:
```bash
cd backend
docker-compose up
```

3. Start the Frontend:
```bash
cd frontend
npm install
npm run dev
```

## API Documentation

### Authentication
- POST /auth/register - Register a new user
- POST /auth/login - Login user

### Cars
- POST /cars - Create a new car (Fleet Owner only)
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
- GET /statistics/bookings - Get booking statistics
- GET /statistics/cars - Get car statistics
- GET /statistics/revenue - Get revenue statistics
- GET /statistics/popular-locations - Get popular locations
- GET /statistics/users - Get user activity statistics

## Contributors

### Kumar Vaibhav
- Backend Architecture and Implementation
- Microservices Development
- Database Design
- API Gateway Implementation
- Docker Configuration
- gRPC Integration
- Testing and Documentation

### Harshita
- Frontend Development
- UI/UX Design
- React Components
- State Management
- API Integration
- User Authentication Flow
- Responsive Design
- Error Handling
- API Gateway Implementation
