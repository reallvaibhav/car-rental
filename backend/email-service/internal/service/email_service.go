package service

import (
	"fmt"
	"log"

	"gopkg.in/gomail.v2"
)

type EmailService struct {
	dialer *gomail.Dialer
	from   string
}

type EmailConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	From     string
}

func NewEmailService(config EmailConfig) *EmailService {
	dialer := gomail.NewDialer(config.Host, config.Port, config.Username, config.Password)

	return &EmailService{
		dialer: dialer,
		from:   config.From,
	}
}

func (s *EmailService) SendBookingConfirmation(to string, bookingData map[string]interface{}) error {
	subject := "Booking Confirmation - Car Rental"
	body := fmt.Sprintf(`
		<h2>Booking Confirmation</h2>
		<p>Dear Customer,</p>
		<p>Your booking has been confirmed with the following details:</p>
		<ul>
			<li>Booking ID: %v</li>
			<li>Car: %v %v</li>
			<li>Start Date: %v</li>
			<li>End Date: %v</li>
			<li>Total Price: $%.2f</li>
		</ul>
		<p>Thank you for choosing our service!</p>
	`,
		bookingData["booking_id"],
		bookingData["car_make"],
		bookingData["car_model"],
		bookingData["start_date"],
		bookingData["end_date"],
		bookingData["total_price"],
	)

	return s.sendEmail(to, subject, body)
}

func (s *EmailService) SendBookingStatusUpdate(to string, bookingID string, status string) error {
	subject := "Booking Status Update - Car Rental"
	body := fmt.Sprintf(`
		<h2>Booking Status Update</h2>
		<p>Dear Customer,</p>
		<p>Your booking (ID: %s) status has been updated to: <strong>%s</strong></p>
		<p>If you have any questions, please contact our support team.</p>
	`, bookingID, status)

	return s.sendEmail(to, subject, body)
}

func (s *EmailService) SendWelcomeEmail(to string, name string) error {
	subject := "Welcome to Car Rental Service"
	body := fmt.Sprintf(`
		<h2>Welcome to Car Rental!</h2>
		<p>Dear %s,</p>
		<p>Thank you for registering with our car rental service. We're excited to have you on board!</p>
		<p>With our service, you can:</p>
		<ul>
			<li>Browse our extensive collection of cars</li>
			<li>Make bookings easily</li>
			<li>Track your rental history</li>
			<li>And much more!</li>
		</ul>
		<p>If you need any assistance, don't hesitate to contact our support team.</p>
	`, name)

	return s.sendEmail(to, subject, body)
}

func (s *EmailService) sendEmail(to, subject, body string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", s.from)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	if err := s.dialer.DialAndSend(m); err != nil {
		log.Printf("Failed to send email: %v", err)
		return fmt.Errorf("failed to send email: %v", err)
	}

	return nil
}
