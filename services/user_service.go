package services

import (
	"database/sql"
	"errors"
	"fmt"
	"net/mail"
	"nyumba/auth"
	"nyumba/models"
	"regexp"
	"strings"
	"time"
)

var (
	ErrEmailExists        = errors.New("email address is already registered")
	ErrInvalidInput       = errors.New("invalid input data")
	ErrInvalidCredentials = errors.New("invalid email or password")
	ErrUserNotFound       = errors.New("user not found")
)

var phoneRegex = regexp.MustCompile(`^\+?[0-9]{10,15}$`)

// ValidateUserRegistration checks input rules before database queries
func ValidateUserRegistration(name, email, phone, password, role string) error {
	if strings.TrimSpace(name) == "" {
		return errors.New("name is required")
	}

	if _, err := mail.ParseAddress(email); err != nil {
		return errors.New("invalid email address format")
	}

	cleanPhone := strings.ReplaceAll(phone, " ", "")
	if !phoneRegex.MatchString(cleanPhone) {
		return errors.New("invalid phone number format (expected e.g. +254700000000)")
	}

	if len(password) < 6 {
		return errors.New("password must be at least 6 characters long")
	}

	validRole := false
	for _, r := range []string{"tenant", "renter", "landlord", "admin"} {
		if strings.EqualFold(role, r) {
			validRole = true
			break
		}
	}
	if !validRole {
		return errors.New("role must be 'tenant', 'renter', 'landlord', or 'admin'")
	}

	return nil
}

// RegisterUser validates input, hashes password, and persists user in DB
func RegisterUser(db *sql.DB, name, email, phone, password, role string) (*models.User, error) {
	if err := ValidateUserRegistration(name, email, phone, password, role); err != nil {
		return nil, fmt.Errorf("%w: %s", ErrInvalidInput, err.Error())
	}

	cleanEmail := strings.ToLower(strings.TrimSpace(email))

	// Check existing email
	var exists bool
	err := db.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE email = $1)", cleanEmail).Scan(&exists)
	if err != nil {
		return nil, fmt.Errorf("failed to check existing user: %w", err)
	}
	if exists {
		return nil, ErrEmailExists
	}

	passwordHash, err := auth.HashPassword(password)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}

	roleLower := strings.ToLower(role)
	if roleLower == "renter" {
		roleLower = "tenant"
	}

	user := &models.User{
		Name:         strings.TrimSpace(name),
		Email:        cleanEmail,
		Phone:        strings.TrimSpace(phone),
		PasswordHash: passwordHash,
		Role:         roleLower,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	query := `
		INSERT INTO users (name, email, phone, password_hash, role, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id
	`
	err = db.QueryRow(query, user.Name, user.Email, user.Phone, user.PasswordHash, user.Role, user.CreatedAt, user.UpdatedAt).Scan(&user.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to insert user: %w", err)
	}

	return user, nil
}

// AuthenticateUser verifies user credentials against PostgreSQL DB
func AuthenticateUser(db *sql.DB, email, password string) (*models.User, error) {
	cleanEmail := strings.ToLower(strings.TrimSpace(email))

	var user models.User
	query := `
		SELECT id, name, email, phone, password_hash, role, created_at, updated_at
		FROM users
		WHERE email = $1
	`
	err := db.QueryRow(query, cleanEmail).Scan(
		&user.ID, &user.Name, &user.Email, &user.Phone, &user.PasswordHash, &user.Role, &user.CreatedAt, &user.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrInvalidCredentials
		}
		return nil, fmt.Errorf("database query error: %w", err)
	}

	if !auth.VerifyPassword(password, user.PasswordHash) {
		return nil, ErrInvalidCredentials
	}

	return &user, nil
}

// GetUserByID retrieves a user profile by ID
func GetUserByID(db *sql.DB, userID int) (*models.User, error) {
	var user models.User
	query := `
		SELECT id, name, email, phone, role, created_at, updated_at
		FROM users
		WHERE id = $1
	`
	err := db.QueryRow(query, userID).Scan(
		&user.ID, &user.Name, &user.Email, &user.Phone, &user.Role, &user.CreatedAt, &user.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	return &user, nil
}
