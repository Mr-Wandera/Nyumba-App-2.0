package auth

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"crypto/subtle"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

var (
	ErrInvalidToken = errors.New("invalid or expired token")
)

// getJWTSecret fetches the secret key for signing tokens
func getJWTSecret() []byte {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		secret = "nyumba_default_secure_jwt_secret_key_2026_change_in_prod"
	}
	return []byte(secret)
}

// HashPassword generates a salt and hashes password using HMAC-SHA256 with salt
func HashPassword(password string) (string, error) {
	salt := make([]byte, 16)
	if _, err := rand.Read(salt); err != nil {
		return "", fmt.Errorf("failed to generate salt: %w", err)
	}

	mac := hmac.New(sha256.New, getJWTSecret())
	mac.Write(salt)
	mac.Write([]byte(password))
	hash := mac.Sum(nil)

	// Format: salt_hex:hash_hex
	return fmt.Sprintf("%s:%s", hex.EncodeToString(salt), hex.EncodeToString(hash)), nil
}

// VerifyPassword checks if raw password matches stored salt:hash string
func VerifyPassword(password, storedHash string) bool {
	parts := strings.Split(storedHash, ":")
	if len(parts) != 2 {
		return false
	}

	salt, err := hex.DecodeString(parts[0])
	if err != nil {
		return false
	}

	expectedHash, err := hex.DecodeString(parts[1])
	if err != nil {
		return false
	}

	mac := hmac.New(sha256.New, getJWTSecret())
	mac.Write(salt)
	mac.Write([]byte(password))
	computedHash := mac.Sum(nil)

	return subtle.ConstantTimeCompare(computedHash, expectedHash) == 1
}

// TokenClaims represents session payload encoded in token
type TokenClaims struct {
	UserID int    `json:"uid"`
	Email  string `json:"email"`
	Role   string `json:"role"`
	Exp    int64  `json:"exp"`
}

// GenerateToken creates an HMAC-SHA256 signed session token (JWT-like format)
func GenerateToken(userID int, email, role string, duration time.Duration) (string, error) {
	claims := TokenClaims{
		UserID: userID,
		Email:  email,
		Role:   role,
		Exp:    time.Now().Add(duration).Unix(),
	}

	claimsBytes, err := json.Marshal(claims)
	if err != nil {
		return "", err
	}

	encodedPayload := base64.RawURLEncoding.EncodeToString(claimsBytes)

	mac := hmac.New(sha256.New, getJWTSecret())
	mac.Write([]byte(encodedPayload))
	signature := base64.RawURLEncoding.EncodeToString(mac.Sum(nil))

	return fmt.Sprintf("%s.%s", encodedPayload, signature), nil
}

// VerifyToken validates and parses signed session token
func VerifyToken(tokenStr string) (*TokenClaims, error) {
	parts := strings.Split(tokenStr, ".")
	if len(parts) != 2 {
		return nil, ErrInvalidToken
	}

	payloadPart, sigPart := parts[0], parts[1]

	mac := hmac.New(sha256.New, getJWTSecret())
	mac.Write([]byte(payloadPart))
	expectedSig := base64.RawURLEncoding.EncodeToString(mac.Sum(nil))

	if subtle.ConstantTimeCompare([]byte(sigPart), []byte(expectedSig)) != 1 {
		return nil, ErrInvalidToken
	}

	payloadBytes, err := base64.RawURLEncoding.DecodeString(payloadPart)
	if err != nil {
		return nil, ErrInvalidToken
	}

	var claims TokenClaims
	if err := json.Unmarshal(payloadBytes, &claims); err != nil {
		return nil, ErrInvalidToken
	}

	if time.Now().Unix() > claims.Exp {
		return nil, ErrInvalidToken
	}

	return &claims, nil
}
