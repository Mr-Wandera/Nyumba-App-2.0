package auth

import (
	"testing"
	"time"
)

func TestHashAndVerifyPassword(t *testing.T) {
	password := "SecretPass123!"

	hash, err := HashPassword(password)
	if err != nil {
		t.Fatalf("Failed to hash password: %v", err)
	}

	if hash == "" {
		t.Fatal("Expected non-empty password hash")
	}

	if !VerifyPassword(password, hash) {
		t.Error("VerifyPassword returned false for correct password")
	}

	if VerifyPassword("WrongPassword", hash) {
		t.Error("VerifyPassword returned true for incorrect password")
	}
}

func TestGenerateAndVerifyToken(t *testing.T) {
	userID := 42
	email := "testuser@example.com"
	role := "landlord"

	token, err := GenerateToken(userID, email, role, 1*time.Hour)
	if err != nil {
		t.Fatalf("Failed to generate token: %v", err)
	}

	claims, err := VerifyToken(token)
	if err != nil {
		t.Fatalf("Failed to verify token: %v", err)
	}

	if claims.UserID != userID || claims.Email != email || claims.Role != role {
		t.Errorf("Claims mismatch: got %+v", claims)
	}
}

func TestExpiredToken(t *testing.T) {
	token, err := GenerateToken(1, "expired@example.com", "tenant", -1*time.Minute)
	if err != nil {
		t.Fatalf("Failed to generate expired token: %v", err)
	}

	_, err = VerifyToken(token)
	if err == nil {
		t.Error("Expected error for expired token, got nil")
	}
}

func TestTamperedToken(t *testing.T) {
	token, _ := GenerateToken(1, "valid@example.com", "tenant", 1*time.Hour)
	tamperedToken := token + "tamper"

	_, err := VerifyToken(tamperedToken)
	if err == nil {
		t.Error("Expected error for tampered token, got nil")
	}
}
