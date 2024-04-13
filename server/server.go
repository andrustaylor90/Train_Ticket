package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"sync"
	"train_ticket_service/proto"

	"github.com/golang-jwt/jwt/v4"
	"google.golang.org/grpc"
)

var mu sync.Mutex         // Protects tickets
var tickets sync.Map      // A simple map to store user email to ticket mapping
var seatCounter int32 = 0 // Atomic counter for seat assignments

func authInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	token, err := authenticate(ctx)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("could not parse claims")
	}

	if role, ok := claims["role"].(string); ok && (role == "admin" || role == "user") {
		return handler(ctx, req)
	}

	return nil, fmt.Errorf("unauthorized")
}

type server struct {
	proto.UnimplementedTrainTicketServiceServer
}

func (s *server) PurchaseTicket(ctx context.Context, req *proto.PurchaseRequest) (*proto.PurchaseReply, error) {
	userEmail, _, err := extractAndValidateTokenFromMetadata(ctx)
	if err != nil {
		return nil, err
	}

	// Ensure the train isn't overbooked
	mu.Lock()
	if seatCounter >= 20 {
		mu.Unlock()
		return nil, fmt.Errorf("all seats on the train are booked")
	}
	seatCounter++
	seat := seatCounter
	mu.Unlock()

	// Assign seats to section A or B
	section := "A"
	seatNumber := seat
	if seat > 10 {
		section = "B"
		seatNumber = seat - 10
	}
	ticketID := fmt.Sprintf("%s-%d", section, seatNumber)

	// Use userEmail as the key
	tickets.Store(userEmail, ticketID)

	return &proto.PurchaseReply{ReceiptId: ticketID}, nil
}

func (s *server) ViewReceipt(ctx context.Context, req *proto.AuthRequest) (*proto.ReceiptReply, error) {
	userEmail, _, err := extractAndValidateTokenFromMetadata(ctx)
	if err != nil {
		return nil, err
	}

	// Retrieve the ticket information associated with the user's email
	if ticketInfo, ok := tickets.Load(userEmail); ok {
		return &proto.ReceiptReply{Details: fmt.Sprintf("Ticket Info: %s", ticketInfo)}, nil
	}

	return nil, fmt.Errorf("no ticket found for user with email: %s", userEmail)
}

func (s *server) RemoveUserFromTrain(ctx context.Context, req *proto.ModifyUserRequest) (*proto.ModifyUserReply, error) {
	userEmail, role, err := extractAndValidateTokenFromMetadata(ctx)
	if err != nil {
		return nil, err
	}

	if role == "user" && userEmail != req.UserEmail {
		return &proto.ModifyUserReply{Success: false}, fmt.Errorf("you're not allowed to remove the user with the email: %s", req.UserEmail)
	}

	mu.Lock()
	defer mu.Unlock()

	_, exists := tickets.Load(req.UserEmail)
	if !exists {
		return &proto.ModifyUserReply{Success: false}, fmt.Errorf("no ticket found for email: %s", req.UserEmail)
	}

	tickets.Delete(req.UserEmail)

	return &proto.ModifyUserReply{Success: true}, nil
}

func (s *server) ModifyUserSeat(ctx context.Context, req *proto.ModifyUserRequest) (*proto.ModifyUserReply, error) {
	_, err := authenticate(ctx) // Authenticates the user or admin
	if err != nil {
		return nil, err
	}

	mu.Lock()
	defer mu.Unlock()

	newSeatID := fmt.Sprintf("section-%s-seat-%d", req.NewSection, req.NewSeat)
	_, exists := tickets.Load(req.UserEmail)
	if !exists {
		return &proto.ModifyUserReply{Success: false}, fmt.Errorf("no ticket found for email: %s", req.UserEmail)
	}

	tickets.Store(req.UserEmail, newSeatID)

	return &proto.ModifyUserReply{Success: true}, nil
}

func (s *server) ViewAllUsers(ctx context.Context, req *proto.AuthRequest) (*proto.UserListReply, error) {
	_, role, err := extractAndValidateTokenFromMetadata(ctx) // Authenticates and checks if admin
	if err != nil {
		return nil, err
	}

	if role == "user" {
		return &proto.UserListReply{Details: "Only Admins can see list of users"}, nil
	}

	mu.Lock()
	defer mu.Unlock()

	var details string
	tickets.Range(func(key, value interface{}) bool {
		details += fmt.Sprintf("User: %v, Seat: %v\n", key, value)
		return true
	})

	return &proto.UserListReply{Details: details}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(authInterceptor),
	)
	proto.RegisterTrainTicketServiceServer(grpcServer, &server{})
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
