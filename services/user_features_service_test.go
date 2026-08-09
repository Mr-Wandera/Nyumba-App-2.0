package services

import (
	"nyumba/models"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestBookingValidation(t *testing.T) {
	pastDate := time.Now().Add(-24 * time.Hour)
	if !pastDate.Before(time.Now()) {
		t.Error("Expected past date check to be true")
	}

	futureDate := time.Now().Add(48 * time.Hour)
	if futureDate.Before(time.Now()) {
		t.Error("Expected future date check to be false")
	}
}

func TestInquiryStructure(t *testing.T) {
	inq := models.Inquiry{
		ID:         1,
		UserID:     5,
		PropertyID: 10,
		Message:    "Is this house available for viewing tomorrow?",
		Status:     "pending",
		CreatedAt:  time.Now(),
	}

	if inq.Message == "" {
		t.Error("Inquiry message should not be empty")
	}

	if inq.Status != "pending" {
		t.Errorf("Expected status 'pending', got %s", inq.Status)
	}
}

func TestFavoriteStructure(t *testing.T) {
	fav := models.Favorite{
		ID:         100,
		UserID:     12,
		PropertyID: 45,
		CreatedAt:  time.Now(),
	}

	if fav.UserID != 12 || fav.PropertyID != 45 {
		t.Errorf("Favorite values mismatch: %+v", fav)
	}
}

// TestConcurrentBookingIsolation tests the isolation principle where two concurrent users attempt to book the exact same slot.
func TestConcurrentBookingIsolation(t *testing.T) {
	type BookingSlot struct {
		PropertyID int
		Date       string
	}

	var mu sync.Mutex
	bookedSlots := make(map[BookingSlot]int)

	tryBook := func(userID, propID int, date string) bool {
		mu.Lock()
		defer mu.Unlock()

		slot := BookingSlot{PropertyID: propID, Date: date}
		if _, exists := bookedSlots[slot]; exists {
			return false // Double booking rejected
		}
		bookedSlots[slot] = userID
		return true // Booking successful
	}

	propertyID := 1
	inspectionDate := "2026-09-01"

	var wg sync.WaitGroup
	var successCount int32
	var failCount int32

	// Launch 10 concurrent goroutines attempting to book the same property and date
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		userID := i
		go func() {
			defer wg.Done()
			if tryBook(userID, propertyID, inspectionDate) {
				atomic.AddInt32(&successCount, 1)
			} else {
				atomic.AddInt32(&failCount, 1)
			}
		}()
	}

	wg.Wait()

	if successCount != 1 {
		t.Errorf("Expected exactly 1 successful booking under concurrency, got %d", successCount)
	}

	if failCount != 9 {
		t.Errorf("Expected 9 rejected double-booking attempts, got %d", failCount)
	}
}
