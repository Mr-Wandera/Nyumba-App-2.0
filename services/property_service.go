package services

import (
	"database/sql"
	"errors"
	"fmt"
	"nyumba/models"
	"time"
)

var (
	ErrPropertyNotFound     = errors.New("property not found")
	ErrUnauthorizedProperty = errors.New("unauthorized to modify this property")
)

// PropertyFilter defines search parameters
type PropertyFilter struct {
	Location       string
	CountyID       int
	SubCountyID    int
	TownID         int
	NeighborhoodID int
	MaxPrice       float64
	Bedrooms       int
	PublishedOnly  bool
}

// GetAllProperties retrieves properties from DB with optional filters
func GetAllProperties(db *sql.DB, filter PropertyFilter) ([]models.Property, error) {
	query := `
		SELECT p.id, p.landlord_id, p.building_name, p.county_id, p.sub_county_id, p.ward_id, p.town_id, p.neighborhood_id, p.location, p.price, p.bedrooms, p.bathrooms, p.description, p.is_paid, COALESCE(p.is_published, true), p.landlord_phone, p.created_at, p.updated_at
		FROM properties p
		WHERE 1=1
	`
	args := []interface{}{}
	argID := 1

	if filter.PublishedOnly {
		query += fmt.Sprintf(" AND COALESCE(p.is_published, true) = $%d", argID)
		args = append(args, true)
		argID++
	}

	if filter.CountyID > 0 {
		query += fmt.Sprintf(" AND p.county_id = $%d", argID)
		args = append(args, filter.CountyID)
		argID++
	}

	if filter.SubCountyID > 0 {
		query += fmt.Sprintf(" AND p.sub_county_id = $%d", argID)
		args = append(args, filter.SubCountyID)
		argID++
	}

	if filter.TownID > 0 {
		query += fmt.Sprintf(" AND p.town_id = $%d", argID)
		args = append(args, filter.TownID)
		argID++
	}

	if filter.NeighborhoodID > 0 {
		query += fmt.Sprintf(" AND p.neighborhood_id = $%d", argID)
		args = append(args, filter.NeighborhoodID)
		argID++
	}

	if filter.Location != "" {
		query += fmt.Sprintf(" AND LOWER(p.location) LIKE $%d", argID)
		args = append(args, "%"+strings.ToLower(filter.Location)+"%")
		argID++
	}

	if filter.MaxPrice > 0 {
		query += fmt.Sprintf(" AND p.price <= $%d", argID)
		args = append(args, filter.MaxPrice)
		argID++
	}

	if filter.Bedrooms > 0 {
		query += fmt.Sprintf(" AND p.bedrooms >= $%d", argID)
		args = append(args, filter.Bedrooms)
		argID++
	}

	query += " ORDER BY p.created_at DESC"

	rows, err := db.Query(query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to query properties: %w", err)
	}
	defer rows.Close()

	var properties []models.Property
	for rows.Next() {
		var p models.Property
		err := rows.Scan(
			&p.ID, &p.LandlordID, &p.BuildingName, &p.CountyID, &p.SubCountyID, &p.WardID, &p.TownID, &p.NeighborhoodID, &p.Location, &p.Price,
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

// ValidatePropertyInput validates property payload
func ValidatePropertyInput(buildingName, location string, price float64) error {
	if buildingName == "" {
		return errors.New("building name is required")
	}
	if location == "" {
		return errors.New("location is required")
	}
	if price <= 0 {
		return errors.New("price must be greater than zero")
	}
	return nil
}

// CreateProperty handles atomic creation of property and images
func CreateProperty(db *sql.DB, p *models.Property, imageURLs []string) error {
	if err := ValidatePropertyInput(p.BuildingName, p.Location, p.Price); err != nil {
		return err
	}

	if err := ValidateLocationHierarchy(db, p.CountyID, p.SubCountyID, p.WardID, p.TownID, p.NeighborhoodID); err != nil {
		return err
	}

	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("failed to start transaction: %w", err)
	}
	defer tx.Rollback()

	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()
	p.IsPublished = true

	query := `
		INSERT INTO properties (landlord_id, building_name, county_id, sub_county_id, ward_id, town_id, neighborhood_id, location, price, bedrooms, bathrooms, description, is_paid, is_published, landlord_phone, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17)
		RETURNING id
	`
	err = tx.QueryRow(
		query, p.LandlordID, p.BuildingName, p.CountyID, p.SubCountyID, p.WardID, p.TownID, p.NeighborhoodID, p.Location, p.Price,
		p.Bedrooms, p.Bathrooms, p.Description, p.IsPaid, p.IsPublished, p.LandlordPhone,
		p.CreatedAt, p.UpdatedAt,
	).Scan(&p.ID)

	if err != nil {
		return fmt.Errorf("failed to insert property: %w", err)
	}

	for idx, imgURL := range imageURLs {
		if imgURL != "" {
			_, err = tx.Exec(
				"INSERT INTO property_images (property_id, image_url, display_order) VALUES ($1, $2, $3)",
				p.ID, imgURL, idx,
			)
			if err != nil {
				return fmt.Errorf("failed to insert property image: %w", err)
			}
		}
	}

	return tx.Commit()
}

// GetPropertyByID fetches a single property
func GetPropertyByID(db *sql.DB, id int) (*models.Property, error) {
	var p models.Property
	query := `
		SELECT id, landlord_id, building_name, county_id, sub_county_id, ward_id, town_id, neighborhood_id, location, price, bedrooms, bathrooms, description, is_paid, COALESCE(is_published, true), landlord_phone, created_at, updated_at
		FROM properties
		WHERE id = $1
	`
	err := db.QueryRow(query, id).Scan(
		&p.ID, &p.LandlordID, &p.BuildingName, &p.CountyID, &p.SubCountyID, &p.WardID, &p.TownID, &p.NeighborhoodID, &p.Location, &p.Price,
		&p.Bedrooms, &p.Bathrooms, &p.Description, &p.IsPaid, &p.IsPublished, &p.LandlordPhone,
		&p.CreatedAt, &p.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrPropertyNotFound
		}
		return nil, err
	}

	p.ImageURLs, _ = GetPropertyImages(db, p.ID)
	return &p, nil
}

// UpdateProperty enforces ownership check and updates property fields
func UpdateProperty(db *sql.DB, updated *models.Property, userID int, userRole string) error {
	if err := ValidatePropertyInput(updated.BuildingName, updated.Location, updated.Price); err != nil {
		return err
	}

	if err := ValidateLocationHierarchy(db, updated.CountyID, updated.SubCountyID, updated.WardID, updated.TownID, updated.NeighborhoodID); err != nil {
		return err
	}

	existing, err := GetPropertyByID(db, updated.ID)
	if err != nil {
		return err
	}

	// Ownership check
	if userRole != "admin" && (existing.LandlordID == nil || *existing.LandlordID != userID) {
		return ErrUnauthorizedProperty
	}

	updated.UpdatedAt = time.Now()
	query := `
		UPDATE properties
		SET building_name = $1, county_id = $2, sub_county_id = $3, ward_id = $4, town_id = $5, neighborhood_id = $6, location = $7, price = $8, bedrooms = $9, bathrooms = $10, description = $11, landlord_phone = $12, updated_at = $13
		WHERE id = $14
	`
	_, err = db.Exec(
		query,
		updated.BuildingName, updated.CountyID, updated.SubCountyID, updated.WardID, updated.TownID, updated.NeighborhoodID, updated.Location, updated.Price,
		updated.Bedrooms, updated.Bathrooms, updated.Description,
		updated.LandlordPhone, updated.UpdatedAt, updated.ID,
	)
	return err
}

// TogglePublishProperty publishes or unpublishes a property with ownership check
func TogglePublishProperty(db *sql.DB, propertyID int, isPublished bool, userID int, userRole string) error {
	existing, err := GetPropertyByID(db, propertyID)
	if err != nil {
		return err
	}

	if userRole != "admin" && (existing.LandlordID == nil || *existing.LandlordID != userID) {
		return ErrUnauthorizedProperty
	}

	_, err = db.Exec("UPDATE properties SET is_published = $1, updated_at = $2 WHERE id = $3", isPublished, time.Now(), propertyID)
	return err
}

// DeleteProperty enforces ownership check before deleting
func DeleteProperty(db *sql.DB, propertyID, userID int, userRole string) error {
	prop, err := GetPropertyByID(db, propertyID)
	if err != nil {
		return err
	}

	// Enforce authorization: landlord owner or admin
	if userRole != "admin" && (prop.LandlordID == nil || *prop.LandlordID != userID) {
		return ErrUnauthorizedProperty
	}

	_, err = db.Exec("DELETE FROM properties WHERE id = $1", propertyID)
	return err
}

// GetPropertyImages helper
func GetPropertyImages(db *sql.DB, propertyID int) ([]string, error) {
	rows, err := db.Query("SELECT image_url FROM property_images WHERE property_id = $1 ORDER BY display_order ASC", propertyID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var urls []string
	for rows.Next() {
		var url string
		if err := rows.Scan(&url); err == nil {
			urls = append(urls, url)
		}
	}
	return urls, nil
}
