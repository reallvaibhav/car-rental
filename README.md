🚗 Car Rental System (Microservices Architecture)
This is a cloud-ready, microservices-based Car Rental System designed for scalability, modularity, and maintainability. Built using modern technologies and principles.

📁 Project Structure
api-gateway: Central entry point for routing requests to backend services.

booking-service: Manages car bookings, availability checks, and booking history.

config: Centralized configuration for all services.

docs: Documentation for APIs, architecture, and setup guides.

email-service: Handles email notifications (booking confirmations, reminders).

inventory-service: Maintains vehicle inventory, status, and details.

nats-service: NATS messaging setup for asynchronous communication between services.

proto: Protocol Buffers for gRPC communication.

statistics-service: Collects and analyzes usage/booking data.

test-data: Sample seed data for development and testing.

user-service: Manages user registration, authentication, and roles.

🚀 Tech Stack
Language: Go (Golang)

Communication: gRPC, REST, NATS

Database: (Insert your DB, e.g., PostgreSQL, MongoDB)

API Gateway: (Insert your chosen gateway, e.g., Kong, custom)

Containerization: Docker

Orchestration: (Optional: Kubernetes)

🧩 Features
Modular microservice-based design

User registration & authentication

Real-time vehicle availability

Booking management with notifications

Centralized configuration and communication layer

Statistics and analytics service

🛠️ Getting Started
Clone the repo: git clone https://github.com/reallvaibhav/car-rental


bash
Copy
Edit
git clone https://github.com/yourusername/car-rental-system.git
Configure environment variables (.env)

Run Docker or use docker-compose (if available)

👥 Team Contributions
Member	Contributions
Kumar Vaibhav	Developed booking-service, user-service, nats-service, and test-data. Integrated gRPC with NATS for async messaging. Wrote Docker setup and handled CI/CD pipeline (if any). Also co-developed the api-gateway.
Harshita	Developed inventory-service, email-service, and statistics-service. Managed project documentation in docs, and helped with .env and config setup. Co-developed the api-gateway.

📌 Future Enhancements
Payment gateway integration

Admin dashboard

Role-based access control (RBAC)

Mobile app support
