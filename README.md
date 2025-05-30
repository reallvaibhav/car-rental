# 🚗 Car Rental System (Microservices Architecture)

This is a microservices-based **Car Rental System** designed for scalability, modularity, and maintainability.

---

## 📁 Project Structure

- **`api-gateway`**: Central entry point for routing requests to backend services.
- **`booking-service`**: Manages car bookings, availability checks, and booking history.
- **`config`**: Centralized configuration for all services.
- **`docs`**: Documentation for APIs, architecture, and setup guides.
- **`email-service`**: Handles email notifications (booking confirmations, reminders).
- **`inventory-service`**: Maintains vehicle inventory, status, and details.
- **`nats-service`**: NATS messaging setup for asynchronous communication between services.
- **`proto`**: Protocol Buffers for gRPC communication.
- **`statistics-service`**: Collects and analyzes usage/booking data.
- **`test-data`**: Sample seed data for development and testing.
- **`user-service`**: Manages user registration, authentication, and roles.

---

## 🚀 Tech Stack

- **Language**: Go (Golang)
- **Communication**: gRPC, REST, NATS
- **Database**: MongoDB
- **API Gateway**: localhost/8080
- **Containerization**: Docker

---

## 🧩 Features

- Modular microservice-based design
- User registration & authentication
- Real-time vehicle availability
- Booking management with notifications
- Centralized configuration and communication layer
- Statistics and analytics service

---

## 🛠️ Getting Started

```bash
# Clone the repository
git clone https://github.com/reallvaibhav/car-rental

# Navigate to the project directory
cd car-rental

# Configure environment variables
cp .env.example .env
# Edit .env with your own configuration

# Run services using Docker (if available)
docker-compose up --build



👥 Team Contributions
Member	Contributions
Kumar Vaibhav	Developed booking-service, user-service, nats-service, and test-data. Integrated gRPC with NATS. Set up Docker and CI/CD. Co-developed api-gateway.
Harshita	Developed inventory-service, email-service, and statistics-service. Handled documentation in docs, and contributed to .env and config. Co-developed api-gateway.


