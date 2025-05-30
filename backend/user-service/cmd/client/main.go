package main

import (
	"context"
	"fmt"
	"log"

	pb_user "proto/user"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// Connect to the server
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	// Create a client
	client := pb_user.NewUserServiceClient(conn)
	ctx := context.Background()

	// Test Register
	fmt.Println("\n=== Testing Register ===")
	registerResp, err := client.Register(ctx, &pb_user.RegisterRequest{
		Email:    "test@example.com",
		Password: "password123",
		Name:     "Test User",
	})
	if err != nil {
		log.Printf("Register failed: %v", err)
	} else {
		fmt.Printf("Register successful! Token: %s\n", registerResp.Token)
	}

	// Test Login
	fmt.Println("\n=== Testing Login ===")
	loginResp, err := client.Login(ctx, &pb_user.LoginRequest{
		Email:    "test@example.com",
		Password: "password123",
	})
	if err != nil {
		log.Printf("Login failed: %v", err)
	} else {
		fmt.Printf("Login successful! Token: %s\n", loginResp.Token)
	}

	// Test GetUserByID
	fmt.Println("\n=== Testing GetUserByID ===")
	userResp, err := client.GetUserByID(ctx, &pb_user.UserIDRequest{
		UserId: "test_user_id", // This will be mocked in test mode
	})
	if err != nil {
		log.Printf("GetUserByID failed: %v", err)
	} else {
		fmt.Printf("User found: %+v\n", userResp)
	}

	// Test ValidateToken
	fmt.Println("\n=== Testing ValidateToken ===")
	validateResp, err := client.ValidateToken(ctx, &pb_user.TokenRequest{
		Token: loginResp.GetToken(),
	})
	if err != nil {
		log.Printf("ValidateToken failed: %v", err)
	} else {
		fmt.Printf("Token validation result: %+v\n", validateResp)
	}

	// Test UpdateProfile
	fmt.Println("\n=== Testing UpdateProfile ===")
	updateResp, err := client.UpdateProfile(ctx, &pb_user.UpdateProfileRequest{
		UserId: "test_user_id",
		Name:   "Updated Test User",
	})
	if err != nil {
		log.Printf("UpdateProfile failed: %v", err)
	} else {
		fmt.Printf("Profile updated: %+v\n", updateResp)
	}

	// Test DeleteUser
	fmt.Println("\n=== Testing DeleteUser ===")
	deleteResp, err := client.DeleteUser(ctx, &pb_user.UserIDRequest{
		UserId: "test_user_id",
	})
	if err != nil {
		log.Printf("DeleteUser failed: %v", err)
	} else {
		fmt.Printf("User deleted: %+v\n", deleteResp)
	}
}
