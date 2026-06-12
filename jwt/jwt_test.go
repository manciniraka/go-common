package jwt

import (
	"testing"
	"time"

	jwtgo "github.com/golang-jwt/jwt/v5"
)

func TestGenerateAndParseToken(t *testing.T) {
	secret := "secret"

	inputClaims := jwtgo.MapClaims{
		"user_id": 1,
		"email":   "raka@mail.com",
		"role":    "admin",
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	}

	token, err := GenerateToken(
		inputClaims,
		secret,
	)
	if err != nil {
		t.Fatalf("generate token error: %v", err)
	}

	claims, err := ParseToken(
		token,
		secret,
	)
	if err != nil {
		t.Fatalf("Parse token error: %v", err)
	}

	if claims["user_id"] != float64(1) {
		t.Fatalf("expected user_id 1, got %v", claims["user_id"])
	}

	if claims["email"] != "raka@mail.com" {
		t.Fatalf("expected email, got %v", claims["email"])
	}

	if claims["role"] != "admin" {
		t.Fatalf("expected role, got %v", claims["role"])
	}
}

func TestParseTokenInvalidSecret(t *testing.T) {
	claims := jwtgo.MapClaims{
		"user_id": 1,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	}

	token, err := GenerateToken(
		claims,
		"secret",
	)
	if err != nil {
		t.Fatal(err)
	}

	_, err = ParseToken(
		token,
		"wrong-secret",
	)
	if err == nil {
		t.Fatal("expected error")
	}
}
