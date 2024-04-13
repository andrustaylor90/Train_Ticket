// auth.go
package main

import (
	"context"
	"fmt"
	"strings"

	"github.com/golang-jwt/jwt/v4"
	"google.golang.org/grpc/metadata"
)

// Helper function to parse and validate JWT token from metadata
func extractAndValidateTokenFromMetadata(ctx context.Context) (string, string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", "", fmt.Errorf("failed to retrieve metadata from context")
	}

	// Retrieve the authorization token from the metadata
	tokenString, exists := md["authorization"]
	if !exists || len(tokenString) == 0 {
		return "", "", fmt.Errorf("authorization token not provided")
	}

	// Trim the 'Bearer ' prefix, if present
	tokenStr := strings.TrimPrefix(tokenString[0], "Bearer ")

	// Parse the JWT without validating the signature
	token, _, err := new(jwt.Parser).ParseUnverified(tokenStr, jwt.MapClaims{})
	if err != nil {
		return "", "", fmt.Errorf("invalid JWT token: %v", err)
	}

	// Extract the claims and validate them
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", "", fmt.Errorf("error parsing token claims")
	}

	// Check for an email claim which identifies the user
	email, ok := claims["email"].(string)
	if !ok {
		return "", "", fmt.Errorf("email not present in token claims")
	}

	// Check for a role claim which identifies the user
	role, ok := claims["role"].(string)
	if !ok {
		return "", "", fmt.Errorf("role not present in token claims")
	}

	return email, role, nil
}

func parseTokenFromContext(ctx context.Context) (string, error) {
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		if len(md["authorization"]) > 0 {
			return md["authorization"][0], nil
		}
	}
	return "", fmt.Errorf("no auth token found")
}

func authenticate(ctx context.Context) (*jwt.Token, error) {
	tokenString, err := parseTokenFromContext(ctx)
	if err != nil {
		return nil, err
	}

	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	token, _, err := new(jwt.Parser).ParseUnverified(tokenString, &jwt.RegisteredClaims{})
	if err != nil {
		return nil, fmt.Errorf("invalid token: %v", err)
	}

	return token, nil
}
