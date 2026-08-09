package services

import (
	"database/sql"
	"errors"
	"fmt"
	"nyumba/models"
	"time"
)

var (
	ErrFavoriteAlreadyExists = errors.New("property is already in favorites")
	ErrFavoriteNotFound      = errors.New("favorite record not found")
	ErrInquiryNotFound       = errors.New("inquiry not found")
	ErrBookingNotFound       = errors.New("booking not found")
	ErrPropertyUnavailable   = errors.New("property is not available for booking")
	ErrDoubleBooking         = errors.New("property is already booked for this date")
	ErrUnauthorizedBooking   = errors.New("unauthorized to modify this booking")
)

// AddFavorite adds a property to user's saved favorites
func AddFavorite(db *sql.DB, userID, propertyID int) (*models.Favorite, error) {
	_, err := GetPropertyByID(db, propertyID)
	if err != nil {
		return nil, err
	}

	var exists bool
	err = db.QueryRow("SELECT EXISTS(SELECT 1 FROM favorites WHERE user_id = $1 AND property_id = $2)", userID, propertyID).Scan(&exists)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, ErrFavoriteAlreadyExists
	}

	fav := &models.Favorite{
		UserID:     userID,
		PropertyID: propertyID,
		CreatedAt:  time.Now(),
	}

	query := "INSERT INTO favorites (user_id, property_id, created_at) VALUES ($1, $2, $3) RETURNING id"
	err = db.QueryRow(query, fav.UserID, fav.PropertyID, fav.CreatedAt).Scan(&fav.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to save favorite: %w", err)
	}

	return fav, nil
}

// RemoveFavorite removes a saved property for a user
func RemoveFavorite(db *sql.DB, userID, propertyID int) error {
	res, err := db.Exec("DELETE FROM favorites WHERE user_id = $1 AND property_id = $2", userID, propertyID)
	if err != nil {
		return err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return ErrFavoriteNotFound
	}

	return nil
}

// GetUserFavorites retrieves all properties favorited by user
func GetUserFavorites(db *sql.DB, userID int) ([]models.Property, error) {
	query := `
		SELECT p.id, p.landlord_id, p.building_name, p.location, p.price, p.bedrooms, p.bathrooms, p.description, p.is_paid, COALESCE(p.is_published, true), p.landlord_phone, p.created_at, p.updated_at
		FROM properties p
		INNER JOIN favorites f ON p.id = f.property_id
		WHERE f.user_id = $1
		ORDER BY f.created_at DESC
	`
	rows, err := db.Query(query, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch favorites: %w", err)
	}
	defer rows.Close()

	var properties []models.Property
	for rows.Next() {
		var p models.Property
		err := rows.Scan(
			&p.ID, &p.LandlordID, &p.BuildingName, &p.Location, &p.Price,
			&p.Bedrooms, &p.Bathrooms, &p.Description, &p.IsPaid, &p.IsPublished, &p.LandlordPhone,
			&p.CreatedAt, &p.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		p.ImageURLs, _ = GetPropertyImages(db, p.ID)
		properties = append(properties, p)
	}

	return properties, nil
}

// CreateInquiry creates a contact inquiry for a property
func CreateInquiry(db *sql.DB, userID, propertyID int, message string) (*models.Inquiry, error) {
	if message == "" {
		return nil, errors.New("inquiry message cannot be empty")
	}

	_, err := GetPropertyByID(db, propertyID)
	if err != nil {
		return nil, err
	}

	inq := &models.Inquiry{
		UserID:     userID,
		PropertyID: propertyID,
		Message:    message,
		Status:     "pending",
		CreatedAt:  time.Now(),
	}

	query := "INSERT INTO inquiries (user_id, property_id, message, status, created_at) VALUES ($1, $2, $3, $4, $5) RETURNING id"
	err = db.QueryRow(query, inq.UserID, inq.PropertyID, inq.Message, inq.Status, inq.CreatedAt).Scan(&inq.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to save inquiry: %w", err)
	}

	return inq, nil
}

// GetPropertyInquiries fetches inquiries for a landlord's property
func GetPropertyInquiries(db *sql.DB, propertyID, userID int, userRole string) ([]models.Inquiry, error) {
	prop, err := GetPropertyByID(db, propertyID)
	if err != nil {
		return nil, err
	}

	if userRole != "admin" && (prop.LandlordID == nil || *prop.LandlordID != userID) {
		return nil, ErrUnauthorizedProperty
	}

	query := "SELECT id, user_id, property_id, message, status, created_at FROM inquiries WHERE property_id = $1 ORDER BY created_at DESC"
	rows, err := db.Query(query, propertyID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var inquiries []models.Inquiry
	for rows.Next() {
		var inq models.Inquiry
		if err := rows.Scan(&inq.ID, &inq.UserID, &inq.PropertyID, &inq.Message, &inq.Status, &inq.CreatedAt); err == nil {
			inquiries = append(inquiries, inq)
		}
	}

	return inquiries, nil
}

// CreateBooking creates an inspection booking with atomic ROW-LOCKING transaction to prevent double bookings
func CreateBooking(db *sql.DB, tenantID, propertyID int, inspectionDate time.Time) (*models.Booking, error) {
	if inspectionDate.Before(time.Now()) {
		return nil, errors.New("inspection date must be in the future")
	}

	tx, err := db.Begin()
	if err != nil {
		return nil, fmt.Errorf("failed to start transaction: %w", err)
	}
	defer tx.Rollback()

	// Row lock on property to ensure strict atomic booking sequence
	var isPublished bool
	var landlordID *int
	err = tx.QueryRow("SELECT COALESCE(is_published, true), landlord_id FROM properties WHERE id = $1 FOR UPDATE", propertyID).Scan(&isPublished, &landlordID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrPropertyNotFound
		}
		return nil, err
	}

	if !isPublished {
		return nil, ErrPropertyUnavailable
	}

	// Check if date slot already booked (excluding cancelled)
	var alreadyBooked bool
	err = tx.QueryRow(`
		SELECT EXISTS(
			SELECT 1 FROM bookings 
			WHERE property_id = $1 
			  AND DATE(inspection_date) = DATE($2) 
			  AND status != 'cancelled'
		)
	`, propertyID, inspectionDate).Scan(&alreadyBooked)

	if err != nil {
		return nil, err
	}
	if alreadyBooked {
		return nil, ErrDoubleBooking
	}

	booking := &models.Booking{
		TenantID:       tenantID,
		PropertyID:     propertyID,
		InspectionDate: inspectionDate,
		Status:         "pending",
		CreatedAt:      time.Now(),
	}

	query := "INSERT INTO bookings (tenant_id, property_id, inspection_date, status, created_at) VALUES ($1, $2, $3, $4, $5) RETURNING id"
	err = tx.QueryRow(query, booking.TenantID, booking.PropertyID, booking.InspectionDate, booking.Status, booking.CreatedAt).Scan(&booking.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to create booking: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return nil, fmt.Errorf("failed to commit booking transaction: %w", err)
	}

	return booking, nil
}

// CancelBooking allows a tenant or landlord/admin to cancel a booking
func CancelBooking(db *sql.DB, bookingID, userID int, userRole string) error {
	var tenantID, propertyID int
	var status string
	err := db.QueryRow("SELECT tenant_id, property_id, status FROM bookings WHERE id = $1", bookingID).Scan(&tenantID, &propertyID, &status)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ErrBookingNotFound
		}
		return err
	}

	if status == "cancelled" {
		return errors.New("booking is already cancelled")
	}

	// Ownership / authorization check
	if userRole != "admin" && tenantID != userID {
		// Check if user is property landlord
		prop, err := GetPropertyByID(db, propertyID)
		if err != nil || prop.LandlordID == nil || *prop.LandlordID != userID {
			return ErrUnauthorizedBooking
		}
	}

	_, err = db.Exec("UPDATE bookings SET status = 'cancelled' WHERE id = $1", bookingID)
	return err
}

// GetUserBookings fetches bookings made by a tenant
func GetUserBookings(db *sql.DB, tenantID int) ([]models.Booking, error) {
	query := "SELECT id, tenant_id, property_id, inspection_date, status, created_at FROM bookings WHERE tenant_id = $1 ORDER BY inspection_date ASC"
	rows, err := db.Query(query, tenantID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var bookings []models.Booking
	for rows.Next() {
		var b models.Booking
		if err := rows.Scan(&b.ID, &b.TenantID, &b.PropertyID, &b.InspectionDate, &b.Status, &b.CreatedAt); err == nil {
			bookings = append(bookings, b)
		}
	}

	return bookings, nil
}
