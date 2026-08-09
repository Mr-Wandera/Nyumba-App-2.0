package models

import "time"

// User represents an authenticated account in Nyumba
type User struct {
	ID           int       `json:"id" db:"id"`
	Name         string    `json:"name" db:"name"`
	Email        string    `json:"email" db:"email"`
	Phone        string    `json:"phone" db:"phone"`
	PasswordHash string    `json:"-" db:"password_hash"`
	Role         string    `json:"role" db:"role"` // tenant, landlord, admin
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
}

// Property represents a house or apartment listing
type Property struct {
	ID             int       `json:"id" db:"id"`
	LandlordID     *int      `json:"landlord_id" db:"landlord_id"`
	BuildingName   string    `json:"building_name" db:"building_name"`
	CountyID       *int      `json:"county_id,omitempty" db:"county_id"`
	SubCountyID    *int      `json:"sub_county_id,omitempty" db:"sub_county_id"`
	WardID         *int      `json:"ward_id,omitempty" db:"ward_id"`
	TownID         *int      `json:"town_id,omitempty" db:"town_id"`
	NeighborhoodID *int      `json:"neighborhood_id,omitempty" db:"neighborhood_id"`
	Location       string    `json:"location" db:"location"`
	Price          float64   `json:"price" db:"price"`
	Bedrooms       int       `json:"bedrooms" db:"bedrooms"`
	Bathrooms      int       `json:"bathrooms" db:"bathrooms"`
	Description    string    `json:"description" db:"description"`
	IsPaid         bool      `json:"is_paid" db:"is_paid"`
	IsPublished    bool      `json:"is_published" db:"is_published"`
	LandlordPhone  string    `json:"landlord_phone" db:"landlord_phone"`
	ImageURLs      []string  `json:"image_urls" db:"-"`
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
	UpdatedAt      time.Time `json:"updated_at" db:"updated_at"`
}

// County represents an official Kenyan county
type County struct {
	ID   int    `json:"id" db:"id"`
	Code int    `json:"code" db:"code"`
	Name string `json:"name" db:"name"`
}

// SubCounty represents a sub-county / constituency within a county
type SubCounty struct {
	ID       int    `json:"id" db:"id"`
	CountyID int    `json:"county_id" db:"county_id"`
	Name     string `json:"name" db:"name"`
}

// Ward represents an electoral ward within a sub-county
type Ward struct {
	ID          int    `json:"id" db:"id"`
	SubCountyID int    `json:"sub_county_id" db:"sub_county_id"`
	Name        string `json:"name" db:"name"`
}

// Town represents a town or urban centre within a county
type Town struct {
	ID       int    `json:"id" db:"id"`
	CountyID int    `json:"county_id" db:"county_id"`
	Name     string `json:"name" db:"name"`
}

// Neighborhood represents an estate, locality, or neighborhood
type Neighborhood struct {
	ID          int    `json:"id" db:"id"`
	SubCountyID *int   `json:"sub_county_id,omitempty" db:"sub_county_id"`
	TownID      *int   `json:"town_id,omitempty" db:"town_id"`
	Name        string `json:"name" db:"name"`
}

// PropertyImage represents an image belonging to a property
type PropertyImage struct {
	ID           int       `json:"id" db:"id"`
	PropertyID   int       `json:"property_id" db:"property_id"`
	ImageURL     string    `json:"image_url" db:"image_url"`
	DisplayOrder int       `json:"display_order" db:"display_order"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
}

// Favorite represents a user's saved property listing
type Favorite struct {
	ID         int       `json:"id" db:"id"`
	UserID     int       `json:"user_id" db:"user_id"`
	PropertyID int       `json:"property_id" db:"property_id"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
}

// Inquiry represents a message sent to a landlord regarding a property
type Inquiry struct {
	ID         int       `json:"id" db:"id"`
	UserID     int       `json:"user_id" db:"user_id"`
	PropertyID int       `json:"property_id" db:"property_id"`
	Message    string    `json:"message" db:"message"`
	Status     string    `json:"status" db:"status"` // pending, read, replied
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
}

// Booking represents a scheduled property inspection booking
type Booking struct {
	ID             int       `json:"id" db:"id"`
	TenantID       int       `json:"tenant_id" db:"tenant_id"`
	PropertyID     int       `json:"property_id" db:"property_id"`
	InspectionDate time.Time `json:"inspection_date" db:"inspection_date"`
	Status         string    `json:"status" db:"status"` // pending, confirmed, cancelled
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
}

// Payment represents an M-Pesa transaction record
type Payment struct {
	ID                     int       `json:"id" db:"id"`
	PropertyID             int       `json:"property_id" db:"property_id"`
	UserID                 int       `json:"user_id" db:"user_id"`
	Amount                 float64   `json:"amount" db:"amount"`
	PhoneNumber            string    `json:"phone_number" db:"phone_number"`
	MpesaMerchantRequestID string    `json:"mpesa_merchant_request_id" db:"mpesa_merchant_request_id"`
	MpesaCheckoutRequestID string    `json:"mpesa_checkout_request_id" db:"mpesa_checkout_request_id"`
	MpesaReceiptNumber     string    `json:"mpesa_receipt_number" db:"mpesa_receipt_number"`
	Status                 string    `json:"status" db:"status"` // pending, completed, failed
	CreatedAt              time.Time `json:"created_at" db:"created_at"`
	UpdatedAt              time.Time `json:"updated_at" db:"updated_at"`
}

// Notification represents a user alert
type Notification struct {
	ID        int       `json:"id" db:"id"`
	UserID    int       `json:"user_id" db:"user_id"`
	Title     string    `json:"title" db:"title"`
	Message   string    `json:"message" db:"message"`
	IsRead    bool      `json:"is_read" db:"is_read"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}
