package main

import (
	"context"
	"log"
	"train_ticket_service/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := proto.NewTrainTicketServiceClient(conn)

	// user token
	token := "Bearer token"

	// email used for JWT claims
	userEmail := "user@example.com"

	// Call the methods
	purchaseTicket(client, token)
	viewReceipt(client, token, userEmail)
	viewAllUsers(client, token)
	removeUserFromTrain(client, token, userEmail)
	modifyUserSeat(client, token, userEmail, "A", 1)
}

func purchaseTicket(client proto.TrainTicketServiceClient, token string) {
	md := metadata.New(map[string]string{"authorization": token})
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	req := &proto.PurchaseRequest{
		From: "London",
		To:   "Paris",
		User: &proto.User{
			FirstName: "John",
			LastName:  "Doe",
			Email:     "john.doe@example.com",
		},
	}
	res, err := client.PurchaseTicket(ctx, req)
	if err != nil {
		log.Fatalf("Error when calling PurchaseTicket: %v", err)
	}
	log.Printf("PurchaseTicket Response: %s", res.ReceiptId)
}

func viewReceipt(client proto.TrainTicketServiceClient, token, email string) {
	md := metadata.New(map[string]string{"authorization": token})
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	req := &proto.AuthRequest{Token: email}
	res, err := client.ViewReceipt(ctx, req)
	if err != nil {
		log.Fatalf("Error when calling ViewReceipt: %v", err)
	}
	log.Printf("ViewReceipt Response: %s", res.Details)
}

func viewAllUsers(client proto.TrainTicketServiceClient, token string) {
	md := metadata.New(map[string]string{"authorization": token})
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	req := &proto.AuthRequest{Token: token}
	res, err := client.ViewAllUsers(ctx, req)
	if err != nil {
		log.Fatalf("Error when calling ViewAllUsers: %v", err)
	}
	log.Printf("ViewAllUsers Response: %s", res.Details)
}

func removeUserFromTrain(client proto.TrainTicketServiceClient, token, email string) {
	md := metadata.New(map[string]string{"authorization": token})
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	req := &proto.ModifyUserRequest{
		Token:     token,
		UserEmail: email,
	}
	res, err := client.RemoveUserFromTrain(ctx, req)
	if err != nil {
		log.Fatalf("Error when calling RemoveUserFromTrain: %v", err)
	}
	log.Printf("RemoveUserFromTrain Response: %v", res.Success)
}

func modifyUserSeat(client proto.TrainTicketServiceClient, token, email, section string, seatNumber int32) {
	md := metadata.New(map[string]string{"authorization": token})
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	req := &proto.ModifyUserRequest{
		Token:      token,
		UserEmail:  email,
		NewSection: section,
		NewSeat:    seatNumber,
	}
	res, err := client.ModifyUserSeat(ctx, req)
	if err != nil {
		log.Fatalf("Error when calling ModifyUserSeat: %v", err)
	}
	log.Printf("ModifyUserSeat Response: %v", res.Success)
}
