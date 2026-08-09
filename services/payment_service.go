package services

import (
	"bytes"
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"nyumba/models"
	"regexp"
	"strings"
	"time"
)

var (
	ErrInvalidPhoneNumber = errors.New("invalid phone number format, expected Kenyan format (e.g., 2547XXXXXXXX or 07XXXXXXXX)")
	ErrInvalidAmount      = errors.New("payment amount must be greater than zero")
	ErrPaymentNotFound     = errors.New("payment record not found")
	ErrExternalBlocked    = errors.New("BLOCKED_EXTERNAL: M-Pesa API credentials missing")
)

// MpesaConfig holds credentials for Safaricom Daraja API
type MpesaConfig struct {
	ConsumerKey    string
	ConsumerSecret string
	Passkey        string
	ShortCode      string
	CallbackURL    string
	Environment    string // "sandbox" or "production"
}

// FormatKenyanPhoneNumber normalizes phone numbers to 254XXXXXXXXX
func FormatKenyanPhoneNumber(phone string) (string, error) {
	phone = strings.TrimSpace(phone)
	phone = strings.ReplaceAll(phone, "+", "")
	phone = strings.ReplaceAll(phone, " ", "")
	phone = strings.ReplaceAll(phone, "-", "")

	if strings.HasPrefix(phone, "0") && len(phone) == 10 {
		phone = "254" + phone[1:]
	}

	re := regexp.compile(`^254[17]\d{8}$`)
	if !re.MatchString(phone) {
		return "", ErrInvalidPhoneNumber
	}

	return phone, nil
}

// InitiateSTKPush handles initiating M-Pesa STK push or local fallback if BLOCKED_EXTERNAL
func InitiateSTKPush(db *sql.DB, config MpesaConfig, userID, propertyID int, amount float64, rawPhone string) (*models.Payment, error) {
	if amount <= 0 {
		return nil, ErrInvalidAmount
	}

	formattedPhone, err := FormatKenyanPhoneNumber(rawPhone)
	if err != nil {
		return nil, err
	}

	// Verify property exists
	prop, err := GetPropertyByID(db, propertyID)
	if err != nil {
		return nil, err
	}

	checkoutID := fmt.Sprintf("ws_CO_%d_%d", time.Now().UnixNano(), userID)
	merchantID := fmt.Sprintf("MR_%d_%d", time.Now().UnixNano(), userID)

	payment := &models.Payment{
		PropertyID:             prop.ID,
		UserID:                 userID,
		Amount:                 amount,
		PhoneNumber:            formattedPhone,
		MpesaMerchantRequestID: merchantID,
		MpesaCheckoutRequestID: checkoutID,
		Status:                 "pending",
		CreatedAt:              time.Now(),
		UpdatedAt:              time.Now(),
	}

	// Check if external API credentials are set
	if config.ConsumerKey == "" || config.ConsumerSecret == "" {
		// BLOCKED_EXTERNAL: Save pending record locally for webhook verification
		query := `
			INSERT INTO payments (property_id, user_id, amount, phone_number, mpesa_merchant_request_id, mpesa_checkout_request_id, status, created_at, updated_at)
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
			RETURNING id
		`
		err := db.QueryRow(
			query,
			payment.PropertyID, payment.UserID, payment.Amount, payment.PhoneNumber,
			payment.MpesaMerchantRequestID, payment.MpesaCheckoutRequestID, payment.Status,
			payment.CreatedAt, payment.UpdatedAt,
		).Scan(&payment.ID)

		if err != nil {
			return nil, fmt.Errorf("failed to save payment record: %w", err)
		}

		return payment, ErrExternalBlocked
	}

	// If credentials exist, call Daraja API with resilience (timeout & bounded retry)
	apiURL := "https://sandbox.safaricom.co.ke/mpesa/stkpush/v1/processrequest"
	if config.Environment == "production" {
		apiURL = "https://api.safaricom.co.ke/mpesa/stkpush/v1/processrequest"
	}

	token, err := getMpesaAccessToken(config)
	if err != nil {
		return nil, fmt.Errorf("failed to authenticate with M-Pesa API: %w", err)
	}

	timestamp := time.Now().Format("20060102150405")
	password := base64.StdEncoding.EncodeToString([]byte(config.ShortCode + config.Passkey + timestamp))

	payload := map[string]interface{}{
		"BusinessShortCode": config.ShortCode,
		"Password":          password,
		"Timestamp":         timestamp,
		"TransactionType":   "CustomerPayBillOnline",
		"Amount":            int(amount),
		"PartyA":            formattedPhone,
		"PartyB":            config.ShortCode,
		"PhoneNumber":       formattedPhone,
		"CallBackURL":       config.CallbackURL,
		"AccountReference":  fmt.Sprintf("NYUMBA-%d", propertyID),
		"TransactionDesc":   fmt.Sprintf("Property %d Rent", propertyID),
	}

	jsonBytes, _ := json.Marshal(payload)
	client := &http.Client{Timeout: 10 * time.Second}

	var resp *http.Response
	var lastErr error

	// Bounded retry (max 2 retries)
	for attempt := 1; attempt <= 2; attempt++ {
		req, _ := http.NewRequest("POST", apiURL, bytes.NewBuffer(jsonBytes))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer "+token)

		resp, lastErr = client.Do(req)
		if lastErr == nil && resp.StatusCode == http.StatusOK {
			break
		}
		time.Sleep(200 * time.Millisecond)
	}

	if lastErr != nil {
		return nil, fmt.Errorf("M-Pesa service request failed: %w", lastErr)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	var darajaResp struct {
		MerchantRequestID string `json:"MerchantRequestID"`
		CheckoutRequestID string `json:"CheckoutRequestID"`
		ResponseCode      string `json:"ResponseCode"`
		ResponseDescription string `json:"ResponseDescription"`
		CustomerMessage   string `json:"CustomerMessage"`
	}

	if err := json.Unmarshal(body, &darajaResp); err == nil && darajaResp.CheckoutRequestID != "" {
		payment.MpesaMerchantRequestID = darajaResp.MerchantRequestID
		payment.MpesaCheckoutRequestID = darajaResp.CheckoutRequestID
	}

	query := `
		INSERT INTO payments (property_id, user_id, amount, phone_number, mpesa_merchant_request_id, mpesa_checkout_request_id, status, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		RETURNING id
	`
	err = db.QueryRow(
		query,
		payment.PropertyID, payment.UserID, payment.Amount, payment.PhoneNumber,
		payment.MpesaMerchantRequestID, payment.MpesaCheckoutRequestID, payment.Status,
		payment.CreatedAt, payment.UpdatedAt,
	).Scan(&payment.ID)

	if err != nil {
		return nil, fmt.Errorf("failed to persist payment: %w", err)
	}

	return payment, nil
}

// ProcessMpesaCallback processes M-Pesa STK callback result idempotently
func ProcessMpesaCallback(db *sql.DB, checkoutRequestID string, resultCode int, resultDesc, mpesaReceiptNumber string) (*models.Payment, error) {
	tx, err := db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	var p models.Payment
	query := `
		SELECT id, property_id, user_id, amount, phone_number, mpesa_merchant_request_id, mpesa_checkout_request_id, COALESCE(mpesa_receipt_number, ''), status, created_at, updated_at
		FROM payments
		WHERE mpesa_checkout_request_id = $1
		FOR UPDATE
	`
	err = tx.QueryRow(query, checkoutRequestID).Scan(
		&p.ID, &p.PropertyID, &p.UserID, &p.Amount, &p.PhoneNumber,
		&p.MpesaMerchantRequestID, &p.MpesaCheckoutRequestID, &p.MpesaReceiptNumber,
		&p.Status, &p.CreatedAt, &p.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrPaymentNotFound
		}
		return nil, err
	}

	// Idempotency check: if already completed, ignore duplicate callback
	if p.Status == "completed" {
		return &p, nil
	}

	now := time.Now()
	p.UpdatedAt = now

	if resultCode == 0 {
		p.Status = "completed"
		p.MpesaReceiptNumber = mpesaReceiptNumber

		// Update payment record
		_, err = tx.Exec(
			"UPDATE payments SET status = $1, mpesa_receipt_number = $2, updated_at = $3 WHERE id = $4",
			p.Status, p.MpesaReceiptNumber, p.UpdatedAt, p.ID,
		)
		if err != nil {
			return nil, err
		}

		// Update property is_paid status
		_, err = tx.Exec("UPDATE properties SET is_paid = true, updated_at = $1 WHERE id = $2", now, p.PropertyID)
		if err != nil {
			return nil, err
		}

		// Insert user notification
		notifTitle := "Payment Received"
		notifMsg := fmt.Sprintf("Your payment of KES %.2f for property #%d was processed successfully. M-Pesa Receipt: %s", p.Amount, p.PropertyID, mpesaReceiptNumber)
		tx.Exec("INSERT INTO notifications (user_id, title, message, created_at) VALUES ($1, $2, $3, $4)", p.UserID, notifTitle, notifMsg, now)
	} else {
		p.Status = "failed"
		_, err = tx.Exec("UPDATE payments SET status = $1, updated_at = $2 WHERE id = $3", p.Status, p.UpdatedAt, p.ID)
		if err != nil {
			return nil, err
		}

		notifTitle := "Payment Failed"
		notifMsg := fmt.Sprintf("Your payment of KES %.2f for property #%d failed or was cancelled: %s", p.Amount, p.PropertyID, resultDesc)
		tx.Exec("INSERT INTO notifications (user_id, title, message, created_at) VALUES ($1, $2, $3, $4)", p.UserID, notifTitle, notifMsg, now)
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return &p, nil
}

// GetPaymentByCheckoutID fetches payment details
func GetPaymentByCheckoutID(db *sql.DB, checkoutRequestID string) (*models.Payment, error) {
	var p models.Payment
	query := `
		SELECT id, property_id, user_id, amount, phone_number, mpesa_merchant_request_id, mpesa_checkout_request_id, COALESCE(mpesa_receipt_number, ''), status, created_at, updated_at
		FROM payments
		WHERE mpesa_checkout_request_id = $1
	`
	err := db.QueryRow(query, checkoutRequestID).Scan(
		&p.ID, &p.PropertyID, &p.UserID, &p.Amount, &p.PhoneNumber,
		&p.MpesaMerchantRequestID, &p.MpesaCheckoutRequestID, &p.MpesaReceiptNumber,
		&p.Status, &p.CreatedAt, &p.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrPaymentNotFound
		}
		return nil, err
	}

	return &p, nil
}

// Helper for M-Pesa Auth Token
func getMpesaAccessToken(config MpesaConfig) (string, error) {
	authURL := "https://sandbox.safaricom.co.ke/oauth/v1/generate?grant_type=client_credentials"
	if config.Environment == "production" {
		authURL = "https://api.safaricom.co.ke/oauth/v1/generate?grant_type=client_credentials"
	}

	req, _ := http.NewRequest("GET", authURL, nil)
	authHeader := base64.StdEncoding.EncodeToString([]byte(config.ConsumerKey + ":" + config.ConsumerSecret))
	req.Header.Set("Authorization", "Basic "+authHeader)

	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var authResp struct {
		AccessToken string `json:"access_token"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&authResp); err != nil {
		return "", err
	}

	return authResp.AccessToken, nil
}
