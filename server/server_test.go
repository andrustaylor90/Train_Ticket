package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"strings"
	"testing"
	"time"
	"train_ticket_service/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

const bufSize = 1024 * 1024

var lis *bufconn.Listener
var client proto.TrainTicketServiceClient

func init() {
	lis = bufconn.Listen(bufSize)
	s := grpc.NewServer()
	proto.RegisterTrainTicketServiceServer(s, &server{})

	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()
}

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

func setupClient() {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to dial bufnet: %v", err)
	}
	client = proto.NewTrainTicketServiceClient(conn)
}

func TestPurchaseTicket(t *testing.T) {
	setupClient()
	req := &proto.PurchaseRequest{
		From: "London",
		To:   "Paris",
		User: &proto.User{
			FirstName: "John",
			LastName:  "Doe",
			Email:     "john.doe@example.com",
		},
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	res, err := client.PurchaseTicket(ctx, req)
	if err != nil {
		t.Fatalf("PurchaseTicket failed: %v", err)
	}
	if res.ReceiptId == "" {
		t.Errorf("Expected a valid receipt ID, got empty")
	}
}

func TestViewReceipt(t *testing.T) {
	setupClient()
	req := &proto.AuthRequest{
		Token: "john.doe@example.com", // Assuming token directly contains the email
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	res, err := client.ViewReceipt(ctx, req)
	if err != nil {
		t.Fatalf("ViewReceipt failed: %v", err)
	}
	if res.Details == "" {
		t.Errorf("Expected ticket details, got empty")
	}
}

func TestViewAllUsers(t *testing.T) {
	setupClient()
	req := &proto.AuthRequest{
		Token: "admin@example.com", // Assuming an 'admin' token that has the privilege to view all users
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	res, err := client.ViewAllUsers(ctx, req)
	if err != nil {
		t.Fatalf("ViewAllUsers failed: %v", err)
	}
	if res.Details == "" {
		t.Errorf("Expected non-empty user list, got empty")
	}
}

func TestRemoveUserFromTrain(t *testing.T) {
	setupClient()
	userEmail := "john.doe@example.com" // Assuming this user exists and has a ticket
	req := &proto.ModifyUserRequest{
		Token:     "admin@example.com", // Admin token, assuming admin rights are needed to remove a user
		UserEmail: userEmail,
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	res, err := client.RemoveUserFromTrain(ctx, req)
	if err != nil {
		t.Fatalf("RemoveUserFromTrain failed: %v", err)
	}
	if !res.Success {
		t.Errorf("Expected successful removal, got failure")
	}

	// Optionally, try to view the receipt to confirm removal
	viewReq := &proto.AuthRequest{Token: userEmail}
	_, err = client.ViewReceipt(ctx, viewReq)
	if err == nil {
		t.Errorf("Expected error when viewing receipt for removed user, got none")
	}
}

func TestModifyUserSeat(t *testing.T) {
	setupClient()
	userEmail := "jane.doe@example.com" // Assuming this user exists
	newSection := "B"
	newSeat := int32(1)
	req := &proto.ModifyUserRequest{
		Token:      "admin@example.com", // Admin or user token, assuming they have the rights to modify seats
		UserEmail:  userEmail,
		NewSection: newSection,
		NewSeat:    newSeat,
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	res, err := client.ModifyUserSeat(ctx, req)
	if err != nil {
		t.Fatalf("ModifyUserSeat failed: %v", err)
	}
	if !res.Success {
		t.Errorf("Expected successful seat modification, got failure")
	}

	viewReq := &proto.AuthRequest{Token: userEmail}
	viewRes, err := client.ViewReceipt(ctx, viewReq)
	if err != nil {
		t.Fatalf("Error when viewing receipt after modification: %v", err)
	}
	expectedSeat := fmt.Sprintf("%s-%d", newSection, newSeat)
	if !strings.Contains(viewRes.Details, expectedSeat) {
		t.Errorf("Expected new seat '%s', but got: %s", expectedSeat, viewRes.Details)
	}
}
