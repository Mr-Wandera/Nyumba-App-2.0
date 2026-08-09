package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"nyumba/auth"
	"nyumba/db"
	"nyumba/middleware"
	"nyumba/models"
	"nyumba/services"
	"nyumba/templates"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

var Houses []models.House

// AddHouseHandler handles adding a property with auth check and file upload
func AddHouseHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Auth & Role check
	claims, err := middleware.ExtractClaims(r)
	if err != nil {
		http.Error(w, "Unauthorized: login required to list property", http.StatusUnauthorized)
		return
	}

	if !strings.EqualFold(claims.Role, "landlord") && !strings.EqualFold(claims.Role, "admin") {
		http.Error(w, "Forbidden: landlord or admin account required to list property", http.StatusForbidden)
		return
	}

	// 10MB upload limit
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		http.Error(w, "Payload too large or invalid multipart form", http.StatusBadRequest)
		return
	}

	buildingName := strings.TrimSpace(r.FormValue("building_name"))
	location := strings.TrimSpace(r.FormValue("location"))
	priceStr := r.FormValue("price")
	bedroomsStr := r.FormValue("bedrooms")
	bathroomsStr := r.FormValue("bathrooms")
	description := strings.TrimSpace(r.FormValue("description"))

	if buildingName == "" || location == "" || priceStr == "" {
		http.Error(w, "Missing required fields: building_name, location, price", http.StatusBadRequest)
		return
	}

	price, err := strconv.ParseFloat(priceStr, 64)
	if err != nil || price <= 0 {
		http.Error(w, "Invalid price amount", http.StatusBadRequest)
		return
	}

	bedrooms := 1
	if bedroomsStr != "" {
		if b, err := strconv.Atoi(bedroomsStr); err == nil && b >= 0 {
			bedrooms = b
		}
	}

	bathrooms := 1
	if bathroomsStr != "" {
		if b, err := strconv.Atoi(bathroomsStr); err == nil && b >= 0 {
			bathrooms = b
		}
	}

	// Handle Image Upload securely
	imagePath := "/uploads/default.jpg"
	file, header, err := r.FormFile("property_photo")
	if err == nil {
		defer file.Close()

		ext := strings.ToLower(filepath.Ext(header.Filename))
		allowedExts := map[string]bool{".jpg": true, ".jpeg": true, ".png": true, ".webp": true}
		if !allowedExts[ext] {
			http.Error(w, "Invalid image format. Allowed: .jpg, .jpeg, .png, .webp", http.StatusBadRequest)
			return
		}

		os.MkdirAll("./uploads", 0755)
		safeFilename := fmt.Sprintf("%d_%s%s", time.Now().UnixNano(), "prop", ext)
		dstPath := filepath.Join("./uploads", safeFilename)

		dst, err := os.Create(dstPath)
		if err != nil {
			http.Error(w, "Failed to save uploaded file", http.StatusInternalServerError)
			return
		}
		defer dst.Close()

		if _, err := io.Copy(dst, file); err != nil {
			http.Error(w, "Failed to copy file contents", http.StatusInternalServerError)
			return
		}
		imagePath = "/uploads/" + safeFilename
	}

	var countyIDPtr, subCountyIDPtr, wardIDPtr, townIDPtr, neighborhoodIDPtr *int
	if cID, err := strconv.Atoi(r.FormValue("county_id")); err == nil && cID > 0 {
		countyIDPtr = &cID
	}
	if scID, err := strconv.Atoi(r.FormValue("sub_county_id")); err == nil && scID > 0 {
		subCountyIDPtr = &scID
	}
	if wID, err := strconv.Atoi(r.FormValue("ward_id")); err == nil && wID > 0 {
		wardIDPtr = &wID
	}
	if tID, err := strconv.Atoi(r.FormValue("town_id")); err == nil && tID > 0 {
		townIDPtr = &tID
	}
	if nID, err := strconv.Atoi(r.FormValue("neighborhood_id")); err == nil && nID > 0 {
		neighborhoodIDPtr = &nID
	}

	database := db.GetDB()
	location = services.BuildLocationString(database, countyIDPtr, subCountyIDPtr, wardIDPtr, townIDPtr, neighborhoodIDPtr, location)

	landlordID := claims.UserID
	newProperty := models.Property{
		LandlordID:     &landlordID,
		BuildingName:   buildingName,
		CountyID:       countyIDPtr,
		SubCountyID:    subCountyIDPtr,
		WardID:         wardIDPtr,
		TownID:         townIDPtr,
		NeighborhoodID: neighborhoodIDPtr,
		Location:       location,
		Price:          price,
		Bedrooms:       bedrooms,
		Bathrooms:      bathrooms,
		Description:    description,
		IsPaid:         false,
		LandlordPhone:  claims.Email,
	}

	// Save to DB if initialized, else fall back to memory
	if database != nil {
		err := services.CreateProperty(database, &newProperty, []string{imagePath})
		if err != nil {
			http.Error(w, "Database error: "+err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		house := models.House{
			ID:            len(Houses) + 1,
			BuildingName:  buildingName,
			Location:      location,
			Price:         price,
			Bedrooms:      bedrooms,
			Bathrooms:     bathrooms,
			Description:   description,
			ImageURLs:     []string{imagePath},
			LandlordPhone: claims.Email,
		}
		Houses = append(Houses, house)
	}

	http.Redirect(w, r, "/explore", http.StatusSeeOther)
}

// DeleteHouseHandler handles deleting a property with ownership checks
func DeleteHouseHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost && r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	claims, err := middleware.ExtractClaims(r)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	propIDStr := r.URL.Query().Get("id")
	if propIDStr == "" {
		propIDStr = r.FormValue("id")
	}

	propID, err := strconv.Atoi(propIDStr)
	if err != nil || propID <= 0 {
		http.Error(w, "Invalid property ID", http.StatusBadRequest)
		return
	}

	database := db.GetDB()
	if database != nil {
		err := services.DeleteProperty(database, propID, claims.UserID, claims.Role)
		if err != nil {
			if errors.Is(err, services.ErrUnauthorizedProperty) {
				http.Error(w, "Forbidden: you do not own this property", http.StatusForbidden)
				return
			}
			if errors.Is(err, services.ErrPropertyNotFound) {
				http.Error(w, "Property not found", http.StatusNotFound)
				return
			}
			http.Error(w, "Failed to delete property: "+err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		newHouses := []models.House{}
		for _, h := range Houses {
			if h.ID != propID {
				newHouses = append(newHouses, h)
			}
		}
		Houses = newHouses
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status":"success","message":"property deleted successfully"}`))
}

// TogglePublishHandler handles publishing or unpublishing a property
func TogglePublishHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	claims, err := middleware.ExtractClaims(r)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	propIDStr := r.FormValue("id")
	if propIDStr == "" {
		propIDStr = r.URL.Query().Get("id")
	}
	propID, err := strconv.Atoi(propIDStr)
	if err != nil || propID <= 0 {
		http.Error(w, "Invalid property ID", http.StatusBadRequest)
		return
	}

	publishAction := r.URL.Path // e.g. /publish-house or /unpublish-house
	isPublished := true
	if strings.Contains(publishAction, "unpublish") {
		isPublished = false
	} else if r.FormValue("is_published") != "" {
		isPublished, _ = strconv.ParseBool(r.FormValue("is_published"))
	}

	database := db.GetDB()
	if database != nil {
		err := services.TogglePublishProperty(database, propID, isPublished, claims.UserID, claims.Role)
		if err != nil {
			if errors.Is(err, services.ErrUnauthorizedProperty) {
				http.Error(w, "Forbidden: you do not own this property", http.StatusForbidden)
				return
			}
			http.Error(w, "Failed to toggle publish status: "+err.Error(), http.StatusInternalServerError)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf(`{"status":"success","is_published":%t}`, isPublished)))
}

// UpdateHouseHandler handles updating an existing property listing
func UpdateHouseHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost && r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	claims, err := middleware.ExtractClaims(r)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	propIDStr := r.FormValue("id")
	propID, err := strconv.Atoi(propIDStr)
	if err != nil || propID <= 0 {
		http.Error(w, "Invalid property ID", http.StatusBadRequest)
		return
	}

	buildingName := strings.TrimSpace(r.FormValue("building_name"))
	location := strings.TrimSpace(r.FormValue("location"))
	priceStr := r.FormValue("price")
	bedroomsStr := r.FormValue("bedrooms")
	bathroomsStr := r.FormValue("bathrooms")
	description := strings.TrimSpace(r.FormValue("description"))

	price, err := strconv.ParseFloat(priceStr, 64)
	if err != nil || price <= 0 {
		http.Error(w, "Invalid price amount", http.StatusBadRequest)
		return
	}

	bedrooms, _ := strconv.Atoi(bedroomsStr)
	bathrooms, _ := strconv.Atoi(bathroomsStr)

	var countyIDPtr, subCountyIDPtr, wardIDPtr, townIDPtr, neighborhoodIDPtr *int
	if cID, err := strconv.Atoi(r.FormValue("county_id")); err == nil && cID > 0 {
		countyIDPtr = &cID
	}
	if scID, err := strconv.Atoi(r.FormValue("sub_county_id")); err == nil && scID > 0 {
		subCountyIDPtr = &scID
	}
	if wID, err := strconv.Atoi(r.FormValue("ward_id")); err == nil && wID > 0 {
		wardIDPtr = &wID
	}
	if tID, err := strconv.Atoi(r.FormValue("town_id")); err == nil && tID > 0 {
		townIDPtr = &tID
	}
	if nID, err := strconv.Atoi(r.FormValue("neighborhood_id")); err == nil && nID > 0 {
		neighborhoodIDPtr = &nID
	}

	database := db.GetDB()
	location = services.BuildLocationString(database, countyIDPtr, subCountyIDPtr, wardIDPtr, townIDPtr, neighborhoodIDPtr, location)

	updatedProp := models.Property{
		ID:             propID,
		BuildingName:   buildingName,
		CountyID:       countyIDPtr,
		SubCountyID:    subCountyIDPtr,
		WardID:         wardIDPtr,
		TownID:         townIDPtr,
		NeighborhoodID: neighborhoodIDPtr,
		Location:       location,
		Price:          price,
		Bedrooms:       bedrooms,
		Bathrooms:      bathrooms,
		Description:    description,
		LandlordPhone:  claims.Email,
	}

	if database != nil {
		err := services.UpdateProperty(database, &updatedProp, claims.UserID, claims.Role)
		if err != nil {
			if errors.Is(err, services.ErrUnauthorizedProperty) {
				http.Error(w, "Forbidden: you do not own this property", http.StatusForbidden)
				return
			}
			http.Error(w, "Failed to update property: "+err.Error(), http.StatusInternalServerError)
			return
		}
	}

	http.Redirect(w, r, "/explore", http.StatusSeeOther)
}

// GetHouses returns JSON property list with optional filters
func GetHouses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if database := db.GetDB(); database != nil {
		filter := services.PropertyFilter{
			Location:      r.URL.Query().Get("location"),
			PublishedOnly: true,
		}
		if countyStr := r.URL.Query().Get("county_id"); countyStr != "" {
			filter.CountyID, _ = strconv.Atoi(countyStr)
		}
		if subCountyStr := r.URL.Query().Get("sub_county_id"); subCountyStr != "" {
			filter.SubCountyID, _ = strconv.Atoi(subCountyStr)
		}
		if townStr := r.URL.Query().Get("town_id"); townStr != "" {
			filter.TownID, _ = strconv.Atoi(townStr)
		}
		if neighStr := r.URL.Query().Get("neighborhood_id"); neighStr != "" {
			filter.NeighborhoodID, _ = strconv.Atoi(neighStr)
		}
		if maxPriceStr := r.URL.Query().Get("max_price"); maxPriceStr != "" {
			filter.MaxPrice, _ = strconv.ParseFloat(maxPriceStr, 64)
		}
		if bedroomsStr := r.URL.Query().Get("bedrooms"); bedroomsStr != "" {
			filter.Bedrooms, _ = strconv.Atoi(bedroomsStr)
		}
		props, err := services.GetAllProperties(database, filter)
		if err != nil {
			http.Error(w, `{"error":"failed to fetch properties"}`, http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(props)
		return
	}

	json.NewEncoder(w).Encode(Houses)
}

// ExploreHandler renders exploration view
func ExploreHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	var houseList []models.House
	if database := db.GetDB(); database != nil {
		props, _ := services.GetAllProperties(database, services.PropertyFilter{})
		for _, p := range props {
			houseList = append(houseList, models.House{
				ID:            p.ID,
				BuildingName:  p.BuildingName,
				Location:      p.Location,
				Price:         p.Price,
				Bedrooms:      p.Bedrooms,
				Bathrooms:     p.Bathrooms,
				Description:   p.Description,
				ImageURLs:     p.ImageURLs,
				IsPaid:        p.IsPaid,
				LandlordPhone: p.LandlordPhone,
			})
		}
	} else {
		houseList = Houses
	}

	fmt.Fprint(w, templates.GetExploreHTML(houseList))
}

// HomePage renders landing view
func HomePage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	var houseList []models.House
	if database := db.GetDB(); database != nil {
		props, _ := services.GetAllProperties(database, services.PropertyFilter{})
		for _, p := range props {
			houseList = append(houseList, models.House{
				ID:            p.ID,
				BuildingName:  p.BuildingName,
				Location:      p.Location,
				Price:         p.Price,
				Bedrooms:      p.Bedrooms,
				Bathrooms:     p.Bathrooms,
				Description:   p.Description,
				ImageURLs:     p.ImageURLs,
				IsPaid:        p.IsPaid,
				LandlordPhone: p.LandlordPhone,
			})
		}
	} else {
		houseList = Houses
	}

	fmt.Fprint(w, templates.GetLandingHTML(houseList))
}

// LoginHandler handles user authentication login
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprint(w, templates.GetAuthHTML("Login"))
		return
	}

	if r.Method == http.MethodPost {
		email := r.FormValue("email")
		password := r.FormValue("password")

		if email == "" || password == "" {
			http.Error(w, "Email and password are required", http.StatusBadRequest)
			return
		}

		database := db.GetDB()
		if database != nil {
			user, err := services.AuthenticateUser(database, email, password)
			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Fprintf(w, `<div style="color:red;padding:20px;">Invalid email or password. <a href="/login">Try again</a></div>`)
				return
			}

			// Generate 24h JWT token
			token, err := auth.GenerateToken(user.ID, user.Email, user.Role, 24*time.Hour)
			if err != nil {
				http.Error(w, "Token generation failed", http.StatusInternalServerError)
				return
			}

			// Set Secure HTTP-only Cookie
			http.SetCookie(w, &http.Cookie{
				Name:     "nyumba_token",
				Value:    token,
				Path:     "/",
				Expires:  time.Now().Add(24 * time.Hour),
				HttpOnly: true,
				SameSite: http.SameSiteLaxMode,
			})

			http.Redirect(w, r, "/explore", http.StatusSeeOther)
			return
		}

		// Memory fallback for demo mode
		token, _ := auth.GenerateToken(1, email, "tenant", 24*time.Hour)
		http.SetCookie(w, &http.Cookie{
			Name:     "nyumba_token",
			Value:    token,
			Path:     "/",
			Expires:  time.Now().Add(24 * time.Hour),
			HttpOnly: true,
			SameSite: http.SameSiteLaxMode,
		})
		http.Redirect(w, r, "/explore", http.StatusSeeOther)
	}
}

// SignupHandler handles user registration
func SignupHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprint(w, templates.GetAuthHTML("Sign Up"))
		return
	}

	if r.Method == http.MethodPost {
		name := r.FormValue("name")
		email := r.FormValue("email")
		phone := r.FormValue("phone")
		password := r.FormValue("password")
		role := r.FormValue("role")

		database := db.GetDB()
		if database != nil {
			user, err := services.RegisterUser(database, name, email, phone, password, role)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				fmt.Fprintf(w, `<div style="color:red;padding:20px;">Registration failed: %s. <a href="/signup">Back to sign up</a></div>`, err.Error())
				return
			}

			token, err := auth.GenerateToken(user.ID, user.Email, user.Role, 24*time.Hour)
			if err != nil {
				http.Error(w, "Token error", http.StatusInternalServerError)
				return
			}

			http.SetCookie(w, &http.Cookie{
				Name:     "nyumba_token",
				Value:    token,
				Path:     "/",
				Expires:  time.Now().Add(24 * time.Hour),
				HttpOnly: true,
				SameSite: http.SameSiteLaxMode,
			})

			http.Redirect(w, r, "/explore", http.StatusSeeOther)
			return
		}

		// Fallback
		token, _ := auth.GenerateToken(1, email, role, 24*time.Hour)
		http.SetCookie(w, &http.Cookie{
			Name:     "nyumba_token",
			Value:    token,
			Path:     "/",
			Expires:  time.Now().Add(24 * time.Hour),
			HttpOnly: true,
			SameSite: http.SameSiteLaxMode,
		})
		http.Redirect(w, r, "/explore", http.StatusSeeOther)
	}
}

// FavoriteHandler handles adding/removing/getting user favorites
func FavoriteHandler(w http.ResponseWriter, r *http.Request) {
	claims, err := middleware.ExtractClaims(r)
	if err != nil {
		http.Error(w, `{"error":"Unauthorized"}`, http.StatusUnauthorized)
		return
	}

	database := db.GetDB()
	if database == nil {
		http.Error(w, `{"error":"Database required"}`, http.StatusServiceUnavailable)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if r.Method == http.MethodGet {
		favs, err := services.GetUserFavorites(database, claims.UserID)
		if err != nil {
			http.Error(w, `{"error":"failed to fetch favorites"}`, http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(favs)
		return
	}

	if r.Method == http.MethodPost {
		propIDStr := r.FormValue("property_id")
		propID, err := strconv.Atoi(propIDStr)
		if err != nil || propID <= 0 {
			http.Error(w, `{"error":"Invalid property ID"}`, http.StatusBadRequest)
			return
		}

		fav, err := services.AddFavorite(database, claims.UserID, propID)
		if err != nil {
			if errors.Is(err, services.ErrFavoriteAlreadyExists) {
				http.Error(w, `{"error":"Already in favorites"}`, http.StatusConflict)
				return
			}
			http.Error(w, `{"error":"Failed to add favorite"}`, http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(fav)
		return
	}

	if r.Method == http.MethodDelete {
		propIDStr := r.URL.Query().Get("property_id")
		propID, err := strconv.Atoi(propIDStr)
		if err != nil || propID <= 0 {
			http.Error(w, `{"error":"Invalid property ID"}`, http.StatusBadRequest)
			return
		}

		err = services.RemoveFavorite(database, claims.UserID, propID)
		if err != nil {
			if errors.Is(err, services.ErrFavoriteNotFound) {
				http.Error(w, `{"error":"Favorite not found"}`, http.StatusNotFound)
				return
			}
			http.Error(w, `{"error":"Failed to remove favorite"}`, http.StatusInternalServerError)
			return
		}
		w.Write([]byte(`{"status":"success","message":"Favorite removed"}`))
		return
	}

	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}

// InquiryHandler handles creating and fetching inquiries
func InquiryHandler(w http.ResponseWriter, r *http.Request) {
	claims, err := middleware.ExtractClaims(r)
	if err != nil {
		http.Error(w, `{"error":"Unauthorized"}`, http.StatusUnauthorized)
		return
	}

	database := db.GetDB()
	if database == nil {
		http.Error(w, `{"error":"Database required"}`, http.StatusServiceUnavailable)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if r.Method == http.MethodPost {
		propIDStr := r.FormValue("property_id")
		message := strings.TrimSpace(r.FormValue("message"))
		propID, err := strconv.Atoi(propIDStr)
		if err != nil || propID <= 0 || message == "" {
			http.Error(w, `{"error":"Invalid property_id or empty message"}`, http.StatusBadRequest)
			return
		}

		inq, err := services.CreateInquiry(database, claims.UserID, propID, message)
		if err != nil {
			http.Error(w, fmt.Sprintf(`{"error":"%s"}`, err.Error()), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(inq)
		return
	}

	if r.Method == http.MethodGet {
		propIDStr := r.URL.Query().Get("property_id")
		propID, err := strconv.Atoi(propIDStr)
		if err != nil || propID <= 0 {
			http.Error(w, `{"error":"Invalid property_id"}`, http.StatusBadRequest)
			return
		}

		inquiries, err := services.GetPropertyInquiries(database, propID, claims.UserID, claims.Role)
		if err != nil {
			if errors.Is(err, services.ErrUnauthorizedProperty) {
				http.Error(w, `{"error":"Forbidden"}`, http.StatusForbidden)
				return
			}
			http.Error(w, `{"error":"Failed to fetch inquiries"}`, http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(inquiries)
		return
	}

	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}

// BookingHandler handles booking inspection appointments
func BookingHandler(w http.ResponseWriter, r *http.Request) {
	claims, err := middleware.ExtractClaims(r)
	if err != nil {
		http.Error(w, `{"error":"Unauthorized"}`, http.StatusUnauthorized)
		return
	}

	database := db.GetDB()
	if database == nil {
		http.Error(w, `{"error":"Database required"}`, http.StatusServiceUnavailable)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if r.Method == http.MethodGet {
		bookings, err := services.GetUserBookings(database, claims.UserID)
		if err != nil {
			http.Error(w, `{"error":"Failed to fetch bookings"}`, http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(bookings)
		return
	}

	if r.Method == http.MethodPost {
		propIDStr := r.FormValue("property_id")
		dateStr := r.FormValue("inspection_date")
		propID, err := strconv.Atoi(propIDStr)
		if err != nil || propID <= 0 || dateStr == "" {
			http.Error(w, `{"error":"Invalid property_id or inspection_date"}`, http.StatusBadRequest)
			return
		}

		inspectionDate, err := time.Parse("2006-01-02", dateStr)
		if err != nil {
			http.Error(w, `{"error":"Invalid date format, use YYYY-MM-DD"}`, http.StatusBadRequest)
			return
		}

		booking, err := services.CreateBooking(database, claims.UserID, propID, inspectionDate)
		if err != nil {
			if errors.Is(err, services.ErrDoubleBooking) {
				http.Error(w, `{"error":"Property is already booked for this date"}`, http.StatusConflict)
				return
			}
			if errors.Is(err, services.ErrPropertyUnavailable) {
				http.Error(w, `{"error":"Property is not available for booking"}`, http.StatusUnprocessableEntity)
				return
			}
			if errors.Is(err, services.ErrPropertyNotFound) {
				http.Error(w, `{"error":"Property not found"}`, http.StatusNotFound)
				return
			}
			http.Error(w, fmt.Sprintf(`{"error":"%s"}`, err.Error()), http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(booking)
		return
	}

	if r.Method == http.MethodDelete {
		bookingIDStr := r.URL.Query().Get("id")
		bookingID, err := strconv.Atoi(bookingIDStr)
		if err != nil || bookingID <= 0 {
			http.Error(w, `{"error":"Invalid booking ID"}`, http.StatusBadRequest)
			return
		}

		err = services.CancelBooking(database, bookingID, claims.UserID, claims.Role)
		if err != nil {
			if errors.Is(err, services.ErrUnauthorizedBooking) {
				http.Error(w, `{"error":"Forbidden: unauthorized to cancel this booking"}`, http.StatusForbidden)
				return
			}
			if errors.Is(err, services.ErrBookingNotFound) {
				http.Error(w, `{"error":"Booking not found"}`, http.StatusNotFound)
				return
			}
			http.Error(w, fmt.Sprintf(`{"error":"%s"}`, err.Error()), http.StatusInternalServerError)
			return
		}

		w.Write([]byte(`{"status":"success","message":"Booking cancelled"}`))
		return
	}

	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}

// PaymentInitiateHandler initiates an M-Pesa STK push for property rent/listing payment
func PaymentInitiateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	claims, err := middleware.ExtractClaims(r)
	if err != nil {
		http.Error(w, `{"error":"Unauthorized"}`, http.StatusUnauthorized)
		return
	}

	database := db.GetDB()
	if database == nil {
		http.Error(w, `{"error":"Database required"}`, http.StatusServiceUnavailable)
		return
	}

	propIDStr := r.FormValue("property_id")
	amountStr := r.FormValue("amount")
	phone := r.FormValue("phone")

	propID, err := strconv.Atoi(propIDStr)
	if err != nil || propID <= 0 {
		http.Error(w, `{"error":"Invalid property_id"}`, http.StatusBadRequest)
		return
	}

	amount, err := strconv.ParseFloat(amountStr, 64)
	if err != nil || amount <= 0 {
		http.Error(w, `{"error":"Invalid amount"}`, http.StatusBadRequest)
		return
	}

	config := services.MpesaConfig{
		ConsumerKey:    os.Getenv("MPESA_CONSUMER_KEY"),
		ConsumerSecret: os.Getenv("MPESA_CONSUMER_SECRET"),
		Passkey:        os.Getenv("MPESA_PASSKEY"),
		ShortCode:      os.Getenv("MPESA_SHORTCODE"),
		CallbackURL:    os.Getenv("MPESA_CALLBACK_URL"),
		Environment:    os.Getenv("MPESA_ENV"),
	}

	payment, err := services.InitiateSTKPush(database, config, claims.UserID, propID, amount, phone)
	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		if errors.Is(err, services.ErrExternalBlocked) {
			w.WriteHeader(http.StatusAccepted)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"status":     "BLOCKED_EXTERNAL",
				"message":    "Live M-Pesa API credentials missing. Payment saved locally in pending state.",
				"payment":    payment,
				"checkout_id": payment.MpesaCheckoutRequestID,
			})
			return
		}
		http.Error(w, fmt.Sprintf(`{"error":"%s"}`, err.Error()), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":     "initiated",
		"payment":    payment,
		"checkout_id": payment.MpesaCheckoutRequestID,
	})
}

// PaymentCallbackHandler processes Safaricom M-Pesa STK Push callbacks
func PaymentCallbackHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	database := db.GetDB()
	if database == nil {
		http.Error(w, `{"error":"Database required"}`, http.StatusServiceUnavailable)
		return
	}

	var payload struct {
		Body struct {
			StkCallback struct {
				MerchantRequestID string `json:"MerchantRequestID"`
				CheckoutRequestID string `json:"CheckoutRequestID"`
				ResultCode        int    `json:"ResultCode"`
				ResultDesc        string `json:"ResultDesc"`
				CallbackMetadata  struct {
					Item []struct {
						Name  string      `json:"Name"`
						Value interface{} `json:"Value"`
					} `json:"Item"`
				} `json:"CallbackMetadata"`
			} `json:"stkCallback"`
		} `json:"Body"`
		// Direct webhook payload format support
		CheckoutRequestID string `json:"checkout_request_id"`
		ResultCode        *int   `json:"result_code"`
		ResultDesc        string `json:"result_desc"`
		MpesaReceipt      string `json:"receipt_number"`
	}

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, `{"error":"Invalid payload"}`, http.StatusBadRequest)
		return
	}

	checkoutID := payload.Body.StkCallback.CheckoutRequestID
	resultCode := payload.Body.StkCallback.ResultCode
	resultDesc := payload.Body.StkCallback.ResultDesc
	receiptNumber := ""

	// Handle direct JSON payload format
	if checkoutID == "" && payload.CheckoutRequestID != "" {
		checkoutID = payload.CheckoutRequestID
		if payload.ResultCode != nil {
			resultCode = *payload.ResultCode
		}
		resultDesc = payload.ResultDesc
		receiptNumber = payload.MpesaReceipt
	} else {
		for _, item := range payload.Body.StkCallback.CallbackMetadata.Item {
			if item.Name == "MpesaReceiptNumber" {
				receiptNumber = fmt.Sprintf("%v", item.Value)
			}
		}
	}

	if checkoutID == "" {
		http.Error(w, `{"error":"Missing CheckoutRequestID"}`, http.StatusBadRequest)
		return
	}

	payment, err := services.ProcessMpesaCallback(database, checkoutID, resultCode, resultDesc, receiptNumber)
	if err != nil {
		if errors.Is(err, services.ErrPaymentNotFound) {
			http.Error(w, `{"error":"Payment record not found"}`, http.StatusNotFound)
			return
		}
		http.Error(w, fmt.Sprintf(`{"error":"%s"}`, err.Error()), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"ResultCode": 0,
		"ResultDesc": "Callback processed successfully",
		"payment":    payment,
	})
}

// PaymentStatusHandler queries status of a payment by checkout ID
func PaymentStatusHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	database := db.GetDB()
	if database == nil {
		http.Error(w, `{"error":"Database required"}`, http.StatusServiceUnavailable)
		return
	}

	checkoutID := r.URL.Query().Get("checkout_id")
	if checkoutID == "" {
		http.Error(w, `{"error":"Missing checkout_id parameter"}`, http.StatusBadRequest)
		return
	}

	payment, err := services.GetPaymentByCheckoutID(database, checkoutID)
	if err != nil {
		if errors.Is(err, services.ErrPaymentNotFound) {
			http.Error(w, `{"error":"Payment not found"}`, http.StatusNotFound)
			return
		}
		http.Error(w, fmt.Sprintf(`{"error":"%s"}`, err.Error()), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(payment)
}

// UploadImageHandler handles uploading a new image for a property listing
func UploadImageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	claims, err := middleware.ExtractClaims(r)
	if err != nil {
		http.Error(w, `{"error":"Unauthorized"}`, http.StatusUnauthorized)
		return
	}

	database := db.GetDB()
	if database == nil {
		http.Error(w, `{"error":"Database required"}`, http.StatusServiceUnavailable)
		return
	}

	propIDStr := r.FormValue("property_id")
	propID, err := strconv.Atoi(propIDStr)
	if err != nil || propID <= 0 {
		http.Error(w, `{"error":"Invalid property_id"}`, http.StatusBadRequest)
		return
	}

	file, header, err := r.FormFile("image")
	if err != nil {
		http.Error(w, `{"error":"Missing image file"}`, http.StatusBadRequest)
		return
	}
	defer file.Close()

	publicURL, err := services.ValidateAndSaveImage(file, header.Filename, header.Size, "./uploads")
	if err != nil {
		http.Error(w, fmt.Sprintf(`{"error":"%s"}`, err.Error()), http.StatusBadRequest)
		return
	}

	imageID, err := services.AddPropertyImageRecord(database, propID, claims.UserID, claims.Role, publicURL)
	if err != nil {
		services.RemovePhysicalFile(publicURL, "./uploads")
		if errors.Is(err, services.ErrUnauthorizedUpload) {
			http.Error(w, `{"error":"Forbidden: you do not own this property"}`, http.StatusForbidden)
			return
		}
		http.Error(w, fmt.Sprintf(`{"error":"%s"}`, err.Error()), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":    "success",
		"image_id":  imageID,
		"image_url": publicURL,
	})
}

// ReplaceImageHandler handles replacing an existing image
func ReplaceImageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost && r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	claims, err := middleware.ExtractClaims(r)
	if err != nil {
		http.Error(w, `{"error":"Unauthorized"}`, http.StatusUnauthorized)
		return
	}

	database := db.GetDB()
	if database == nil {
		http.Error(w, `{"error":"Database required"}`, http.StatusServiceUnavailable)
		return
	}

	imageIDStr := r.FormValue("image_id")
	imageID, err := strconv.Atoi(imageIDStr)
	if err != nil || imageID <= 0 {
		http.Error(w, `{"error":"Invalid image_id"}`, http.StatusBadRequest)
		return
	}

	file, header, err := r.FormFile("image")
	if err != nil {
		http.Error(w, `{"error":"Missing replacement image file"}`, http.StatusBadRequest)
		return
	}
	defer file.Close()

	newURL, err := services.ValidateAndSaveImage(file, header.Filename, header.Size, "./uploads")
	if err != nil {
		http.Error(w, fmt.Sprintf(`{"error":"%s"}`, err.Error()), http.StatusBadRequest)
		return
	}

	err = services.ReplacePropertyImageRecord(database, imageID, claims.UserID, claims.Role, newURL, "./uploads")
	if err != nil {
		services.RemovePhysicalFile(newURL, "./uploads")
		if errors.Is(err, services.ErrUnauthorizedUpload) {
			http.Error(w, `{"error":"Forbidden"}`, http.StatusForbidden)
			return
		}
		if errors.Is(err, services.ErrImageNotFound) {
			http.Error(w, `{"error":"Image not found"}`, http.StatusNotFound)
			return
		}
		http.Error(w, fmt.Sprintf(`{"error":"%s"}`, err.Error()), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":    "success",
		"image_id":  imageID,
		"image_url": newURL,
	})
}

// DeleteImageHandler handles deleting an image and its physical file
func DeleteImageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost && r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	claims, err := middleware.ExtractClaims(r)
	if err != nil {
		http.Error(w, `{"error":"Unauthorized"}`, http.StatusUnauthorized)
		return
	}

	database := db.GetDB()
	if database == nil {
		http.Error(w, `{"error":"Database required"}`, http.StatusServiceUnavailable)
		return
	}

	imageIDStr := r.URL.Query().Get("id")
	if imageIDStr == "" {
		imageIDStr = r.FormValue("image_id")
	}
	imageID, err := strconv.Atoi(imageIDStr)
	if err != nil || imageID <= 0 {
		http.Error(w, `{"error":"Invalid image_id"}`, http.StatusBadRequest)
		return
	}

	err = services.DeletePropertyImageRecord(database, imageID, claims.UserID, claims.Role, "./uploads")
	if err != nil {
		if errors.Is(err, services.ErrUnauthorizedUpload) {
			http.Error(w, `{"error":"Forbidden"}`, http.StatusForbidden)
			return
		}
		if errors.Is(err, services.ErrImageNotFound) {
			http.Error(w, `{"error":"Image not found"}`, http.StatusNotFound)
			return
		}
		http.Error(w, fmt.Sprintf(`{"error":"%s"}`, err.Error()), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"status":"success","message":"Image deleted successfully"}`))
}

// LogoutHandler logs out user by clearing cookie
func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:     "nyumba_token",
		Value:    "",
		Path:     "/",
		MaxAge:   -1,
		Expires:  time.Now().Add(-1 * time.Hour),
		HttpOnly: true,
	})
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

// SeedHouses initializes seed data in memory or DB
func SeedHouses() {
	if len(Houses) > 0 {
		return
	}
	Houses = append(Houses, models.House{
		ID:            1,
		BuildingName:  "Sunset Heights",
		Location:      "Section 9, Thika",
		Price:         15000,
		Deposit:       15000,
		Type:          "Apartment",
		ImageURLs:     []string{"/uploads/default.jpg"},
		IsPaid:        false,
		Bedrooms:      2,
		Bathrooms:     1,
		LandlordPhone: "+254712345678",
		Description:   "Quiet neighborhood, great views.",
	})
}
