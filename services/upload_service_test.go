package services

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"
)

func TestValidateAndSaveImage_ValidJPEG(t *testing.T) {
	// Create a minimal valid JPEG header (SOI marker 0xFFD8FF)
	validJPEG := []byte{0xFF, 0xD8, 0xFF, 0xE0, 0x00, 0x10, 0x4A, 0x46, 0x49, 0x46, 0x00}
	for i := 0; i < 500; i++ {
		validJPEG = append(validJPEG, 0x00)
	}

	reader := bytes.NewReader(validJPEG)
	tempDir := t.TempDir()

	savedPath, err := ValidateAndSaveImage(reader, "test.jpg", int64(len(validJPEG)), tempDir)
	if err != nil {
		t.Fatalf("Expected valid image save to succeed, got: %v", err)
	}

	if savedPath == "" {
		t.Error("Expected non-empty saved image path")
	}

	// Verify file created on disk
	filename := filepath.Base(savedPath)
	if _, err := os.Stat(filepath.Join(tempDir, filename)); os.IsNotExist(err) {
		t.Errorf("Expected saved file %s to exist on disk", filename)
	}
}

func TestValidateAndSaveImage_Oversized(t *testing.T) {
	reader := bytes.NewReader([]byte("fake data"))
	tempDir := t.TempDir()

	hugeSize := int64(6 * 1024 * 1024) // 6 MB
	_, err := ValidateAndSaveImage(reader, "huge.jpg", hugeSize, tempDir)
	if err != ErrFileTooLarge {
		t.Errorf("Expected ErrFileTooLarge, got: %v", err)
	}
}

func TestValidateAndSaveImage_InvalidExtension(t *testing.T) {
	reader := bytes.NewReader([]byte("executable content"))
	tempDir := t.TempDir()

	_, err := ValidateAndSaveImage(reader, "script.sh", 100, tempDir)
	if err != ErrInvalidFileType {
		t.Errorf("Expected ErrInvalidFileType, got: %v", err)
	}
}

func TestValidateAndSaveImage_DisguisedFile(t *testing.T) {
	// File with .jpg extension but plain text content
	plainText := []byte("<html><body>Malicious HTML Script</body></html>")
	reader := bytes.NewReader(plainText)
	tempDir := t.TempDir()

	_, err := ValidateAndSaveImage(reader, "malicious.jpg", int64(len(plainText)), tempDir)
	if err != ErrInvalidFileType {
		t.Errorf("Expected MIME rejection ErrInvalidFileType, got: %v", err)
	}
}

func TestRemovePhysicalFile(t *testing.T) {
	tempDir := t.TempDir()
	testFile := filepath.Join(tempDir, "sample.jpg")

	os.WriteFile(testFile, []byte("sample image content"), 0644)

	publicURL := "/uploads/sample.jpg"
	RemovePhysicalFile(publicURL, tempDir)

	if _, err := os.Stat(testFile); !os.IsNotExist(err) {
		t.Error("Expected test file to be removed from disk")
	}
}
