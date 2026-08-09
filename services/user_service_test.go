package services

import (
	"testing"
)

func TestValidateUserRegistration(t *testing.T) {
	tests := []struct {
		name     string
		userName string
		email    string
		phone    string
		password string
		role     string
		wantErr  bool
	}{
		{
			name:     "Valid Landlord Registration",
			userName: "Abdul Wandera",
			email:    "abdul@example.com",
			phone:    "+254700000000",
			password: "securepassword",
			role:     "landlord",
			wantErr:  false,
		},
		{
			name:     "Valid Renter Alias",
			userName: "Jane Doe",
			email:    "jane@example.com",
			phone:    "0712345678",
			password: "password123",
			role:     "renter",
			wantErr:  false,
		},
		{
			name:     "Invalid Email Format",
			userName: "Bad Email",
			email:    "invalid-email-address",
			phone:    "+254700000000",
			password: "password123",
			role:     "tenant",
			wantErr:  true,
		},
		{
			name:     "Password Too Short",
			userName: "Short Pass",
			email:    "short@example.com",
			phone:    "+254700000000",
			password: "123",
			role:     "tenant",
			wantErr:  true,
		},
		{
			name:     "Invalid Role",
			userName: "Unknown Role",
			email:    "role@example.com",
			phone:    "+254700000000",
			password: "password123",
			role:     "hacker",
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateUserRegistration(tt.userName, tt.email, tt.phone, tt.password, tt.role)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateUserRegistration() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
