package services

import (
	"database/sql"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var (
	ErrFileTooLarge       = errors.New("file exceeds maximum allowed size of 5MB")
	ErrInvalidFileType    = errors.New("invalid file type, allowed formats: JPEG, PNG, WEBP")
	ErrImageNotFound      = errors.New("property image not found")
	ErrUnauthorizedUpload = errors.New("unauthorized to manage images for this property")
)

const MaxImageSizeBytes = 5 * 1024 * 1024 // 5 MB

// ValidateAndSaveImage validates image size, MIME header, extension and writes to uploads directory
func ValidateAndSaveImage(file io.Reader, filename string, fileSize int64, uploadsDir string) (string, error) {
	if fileSize > MaxImageSizeBytes {
		return "", ErrFileTooLarge
	}

	ext := strings.ToLower(filepath.Ext(filename))
	allowedExts := map[string]bool{".jpg": true, ".jpeg": true, ".png": true, ".webp": true}
	if !allowedExts[ext] {
		return "", ErrInvalidFileType
	}

	// Read first 512 bytes for MIME type verification
	head := make([]byte, 512)
	n, err := file.Read(head)
	if err != nil && err != io.EOF {
		return "", fmt.Errorf("failed to read file header: %w", err)
	}

	mimeType := http.DetectContentType(head[:n])
	allowedMimes := map[string]bool{
		"image/jpeg": true,
		"image/png":  true,
		"image/webp": true,
	}
	if !allowedMimes[mimeType] {
		return "", ErrInvalidFileType
	}

	if err := os.MkdirAll(uploadsDir, 0755); err != nil {
		return "", fmt.Errorf("failed to create upload directory: %w", err)
	}

	safeFilename := fmt.Sprintf("%d_img%s", time.Now().UnixNano(), ext)
	targetPath := filepath.Join(uploadsDir, safeFilename)

	out, err := os.Create(targetPath)
	if err != nil {
		return "", fmt.Errorf("failed to create target file: %w", err)
	}
	defer out.Close()

	// Write head bytes first, then remaining
	if _, err := out.Write(head[:n]); err != nil {
		return "", err
	}
	if _, err := io.Copy(out, file); err != nil {
		return "", err
	}

	return "/uploads/" + safeFilename, nil
}

// AddPropertyImageRecord inserts new image into property_images database table
func AddPropertyImageRecord(db *sql.DB, propertyID, userID int, userRole string, imageURL string) (int, error) {
	prop, err := GetPropertyByID(db, propertyID)
	if err != nil {
		return 0, err
	}

	if userRole != "admin" && (prop.LandlordID == nil || *prop.LandlordID != userID) {
		return 0, ErrUnauthorizedUpload
	}

	var imageID int
	query := "INSERT INTO property_images (property_id, image_url, display_order) VALUES ($1, $2, $3) RETURNING id"
	err = db.QueryRow(query, propertyID, imageURL, 0).Scan(&imageID)
	if err != nil {
		return 0, fmt.Errorf("failed to save image record: %w", err)
	}

	return imageID, nil
}

// DeletePropertyImageRecord deletes property image from DB and cleans up disk file
func DeletePropertyImageRecord(db *sql.DB, imageID, userID int, userRole string, uploadsDir string) error {
	var propertyID int
	var imageURL string

	err := db.QueryRow("SELECT property_id, image_url FROM property_images WHERE id = $1", imageID).Scan(&propertyID, &imageURL)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ErrImageNotFound
		}
		return err
	}

	prop, err := GetPropertyByID(db, propertyID)
	if err != nil {
		return err
	}

	if userRole != "admin" && (prop.LandlordID == nil || *prop.LandlordID != userID) {
		return ErrUnauthorizedUpload
	}

	_, err = db.Exec("DELETE FROM property_images WHERE id = $1", imageID)
	if err != nil {
		return err
	}

	// Clean up physical file if located in uploadsDir
	RemovePhysicalFile(imageURL, uploadsDir)

	return nil
}

// ReplacePropertyImageRecord updates property image record and cleans up old file
func ReplacePropertyImageRecord(db *sql.DB, imageID, userID int, userRole string, newImageURL, uploadsDir string) error {
	var propertyID int
	var oldImageURL string

	err := db.QueryRow("SELECT property_id, image_url FROM property_images WHERE id = $1", imageID).Scan(&propertyID, &oldImageURL)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ErrImageNotFound
		}
		return err
	}

	prop, err := GetPropertyByID(db, propertyID)
	if err != nil {
		return err
	}

	if userRole != "admin" && (prop.LandlordID == nil || *prop.LandlordID != userID) {
		return ErrUnauthorizedUpload
	}

	_, err = db.Exec("UPDATE property_images SET image_url = $1 WHERE id = $2", newImageURL, imageID)
	if err != nil {
		return err
	}

	// Remove old physical file
	RemovePhysicalFile(oldImageURL, uploadsDir)

	return nil
}

// RemovePhysicalFile deletes file from disk safely
func RemovePhysicalFile(publicURL, uploadsDir string) {
	if strings.HasPrefix(publicURL, "/uploads/") {
		filename := filepath.Base(publicURL)
		filePath := filepath.Join(uploadsDir, filename)
		os.Remove(filePath)
	}
}
