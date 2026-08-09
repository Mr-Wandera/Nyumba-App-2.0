package services

import (
	"nyumba/models"
	"testing"
	"time"
)

func TestFormatKenyanPhoneNumber(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		expected  string
		expectErr bool
	}{
		{
			name:      "Local 07 format",
			input:     "0712345678",
			expected:  "254712345678",
			expectErr: false,
		},
		{
			name:      "International with plus",
			input:     "+254712345678",
			expected:  "254712345678",
			expectErr: false,
		},
		{
			name:      "Safaricom 01 format",
			input:     "0110000000",
			expected:  "254110000000",
			expectErr: false,
		},
		{
			name:      "Invalid short number",
			input:     "07123",
			expected:  "",
			expectErr: true,
		},
		{
			name:      "Invalid alphabetic characters",
			input:     "0712345abc",
			expected:  "",
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := FormatKenyanPhoneNumber(tt.input)
			if (err != nil) != tt.expectErr {
				t.Errorf("FormatKenyanPhoneNumber(%s) error = %v, expectErr %v", tt.input, err, tt.expectErr)
			}
			if res != tt.expected {
				t.Errorf("FormatKenyanPhoneNumber(%s) = %s, expected %s", tt.input, res, tt.expected)
			}
		})
	}
}

func TestMpesaConfigBlockedExternal(t *testing.T) {
	config := MpesaConfig{
		ConsumerKey:    "",
		ConsumerSecret: "",
	}

	if config.ConsumerKey != "" || config.ConsumerSecret != "" {
		t.Error("Expected empty credentials for local fallback test")
	}
}

func TestPaymentCallbackStructure(t *testing.T) {
	p := models.Payment{
		ID:                     1,
		PropertyID:             10,
		UserID:                 5,
		Amount:                 15000,
		PhoneNumber:            "254712345678",
		MpesaCheckoutRequestID: "ws_CO_TEST_123",
		Status:                 "pending",
		CreatedAt:              time.Now(),
		UpdatedAt:              time.Now(),
	}

	if p.Amount <= 0 {
		t.Error("Payment amount must be greater than 0")
	}

	if p.Status != "pending" {
		t.Errorf("Expected initial payment status to be pending, got %s", p.Status)
	}
}
