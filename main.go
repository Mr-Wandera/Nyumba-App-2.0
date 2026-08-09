package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"nyumba/db"
	"nyumba/handlers"
	"nyumba/middleware"
	"nyumba/services"
	"nyumba/templates"
)

func main() {
	// Initialize PostgreSQL connection if DATABASE_URL is configured
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL != "" {
		database, err := db.InitDB(dbURL)
		if err != nil {
			log.Printf("Warning: Failed to connect to PostgreSQL: %v. Running with fallback.", err)
		} else {
			defer database.Close()
			runMigrations(database)
		}
	} else {
		log.Println("DATABASE_URL not set. Running with local storage fallback.")
		handlers.SeedHouses()
	}

	// Static uploads routing
	os.MkdirAll("./uploads", 0755)
	fs := http.FileServer(http.Dir("./uploads"))
	http.Handle("/uploads/", http.StripPrefix("/uploads/", fs))

	// Public Routes
	http.HandleFunc("/", handlers.HomePage)
	http.HandleFunc("/explore", handlers.ExploreHandler)
	http.HandleFunc("/houses", handlers.GetHouses)
	http.HandleFunc("/login", handlers.LoginHandler)
	http.HandleFunc("/signup", handlers.SignupHandler)
	http.HandleFunc("/logout", handlers.LogoutHandler)
	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprint(w, templates.GetStaticHTML("About Us", "Nyumba is Kenya's premier sanctuary discovery platform. We connect verified renters with real landlords directly."))
	})
	http.HandleFunc("/contact", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprint(w, templates.GetStaticHTML("Contact Support", "Need help? Reach out to our 24/7 support team at support@nyumba.co.ke or call +254 700 000 000."))
	})

	// Location Discovery API Routes
	http.HandleFunc("/api/locations/counties", handlers.GetCountiesHandler)
	http.HandleFunc("/api/locations/sub-counties", handlers.GetSubCountiesHandler)
	http.HandleFunc("/api/locations/wards", handlers.GetWardsHandler)
	http.HandleFunc("/api/locations/towns", handlers.GetTownsHandler)
	http.HandleFunc("/api/locations/neighborhoods", handlers.GetNeighborhoodsHandler)
	http.HandleFunc("/api/locations/hierarchy", handlers.GetLocationHierarchyHandler)

	// Protected Landlord Routes
	http.HandleFunc("/landlord", middleware.RoleRequired([]string{"landlord", "admin"}, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprint(w, templates.GetLandlordHTML())
	}))
	http.HandleFunc("/add-house", handlers.AddHouseHandler)
	http.HandleFunc("/update-house", handlers.UpdateHouseHandler)
	http.HandleFunc("/publish-house", handlers.TogglePublishHandler)
	http.HandleFunc("/unpublish-house", handlers.TogglePublishHandler)
	http.HandleFunc("/delete-house", handlers.DeleteHouseHandler)

	// User Features Routes
	http.HandleFunc("/favorites", handlers.FavoriteHandler)
	http.HandleFunc("/inquiries", handlers.InquiryHandler)
	http.HandleFunc("/bookings", handlers.BookingHandler)

	// Payment Routes
	http.HandleFunc("/api/payments/initiate", handlers.PaymentInitiateHandler)
	http.HandleFunc("/api/payments/callback", handlers.PaymentCallbackHandler)
	http.HandleFunc("/api/payments/status", handlers.PaymentStatusHandler)

	// Image Management & Upload Static File Routes
	http.HandleFunc("/api/properties/images/upload", handlers.UploadImageHandler)
	http.HandleFunc("/api/properties/images/replace", handlers.ReplaceImageHandler)
	http.HandleFunc("/api/properties/images/delete", handlers.DeleteImageHandler)
	os.MkdirAll("./uploads", 0755)
	http.Handle("/uploads/", http.StripPrefix("/uploads/", http.FileServer(http.Dir("./uploads"))))

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	fmt.Printf("Nyumba Go Server listening on port %s...\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func runMigrations(database *sql.DB) {
	// Auto-create initial schema tables if not exist
	query := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		email VARCHAR(255) UNIQUE NOT NULL,
		phone VARCHAR(50) NOT NULL,
		password_hash VARCHAR(255) NOT NULL,
		role VARCHAR(20) NOT NULL DEFAULT 'tenant',
		created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
	);
	CREATE TABLE IF NOT EXISTS properties (
		id SERIAL PRIMARY KEY,
		landlord_id INT REFERENCES users(id) ON DELETE SET NULL,
		building_name VARCHAR(255) NOT NULL,
		location VARCHAR(255) NOT NULL,
		price NUMERIC(12, 2) NOT NULL CHECK (price >= 0),
		bedrooms INT NOT NULL DEFAULT 1 CHECK (bedrooms >= 0),
		bathrooms INT NOT NULL DEFAULT 1 CHECK (bathrooms >= 0),
		description TEXT,
		is_paid BOOLEAN NOT NULL DEFAULT FALSE,
		is_published BOOLEAN NOT NULL DEFAULT TRUE,
		landlord_phone VARCHAR(50) NOT NULL,
		created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
	);
	ALTER TABLE properties ADD COLUMN IF NOT EXISTS is_published BOOLEAN DEFAULT TRUE;
	CREATE TABLE IF NOT EXISTS property_images (
		id SERIAL PRIMARY KEY,
		property_id INT NOT NULL REFERENCES properties(id) ON DELETE CASCADE,
		image_url TEXT NOT NULL,
		display_order INT DEFAULT 0,
		created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
	);
	CREATE TABLE IF NOT EXISTS favorites (
		id SERIAL PRIMARY KEY,
		user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
		property_id INT NOT NULL REFERENCES properties(id) ON DELETE CASCADE,
		created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
		UNIQUE(user_id, property_id)
	);
	CREATE TABLE IF NOT EXISTS inquiries (
		id SERIAL PRIMARY KEY,
		user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
		property_id INT NOT NULL REFERENCES properties(id) ON DELETE CASCADE,
		message TEXT NOT NULL,
		status VARCHAR(20) NOT NULL DEFAULT 'pending',
		created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
	);
	CREATE TABLE IF NOT EXISTS bookings (
		id SERIAL PRIMARY KEY,
		tenant_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
		property_id INT NOT NULL REFERENCES properties(id) ON DELETE CASCADE,
		inspection_date TIMESTAMP WITH TIME ZONE NOT NULL,
		status VARCHAR(20) NOT NULL DEFAULT 'pending',
		created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
	);
	CREATE TABLE IF NOT EXISTS payments (
		id SERIAL PRIMARY KEY,
		property_id INT NOT NULL REFERENCES properties(id) ON DELETE CASCADE,
		user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
		amount NUMERIC(12, 2) NOT NULL CHECK (amount > 0),
		phone_number VARCHAR(50) NOT NULL,
		mpesa_merchant_request_id VARCHAR(100),
		mpesa_checkout_request_id VARCHAR(100) UNIQUE,
		mpesa_receipt_number VARCHAR(100),
		status VARCHAR(20) NOT NULL DEFAULT 'pending',
		created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
	);
	CREATE TABLE IF NOT EXISTS notifications (
		id SERIAL PRIMARY KEY,
		user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
		title VARCHAR(255) NOT NULL,
		message TEXT NOT NULL,
		is_read BOOLEAN NOT NULL DEFAULT FALSE,
		created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
	);

	-- Location System Schema
	CREATE TABLE IF NOT EXISTS counties (
		id SERIAL PRIMARY KEY,
		code INT UNIQUE NOT NULL,
		name VARCHAR(100) UNIQUE NOT NULL,
		created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
	);
	CREATE TABLE IF NOT EXISTS sub_counties (
		id SERIAL PRIMARY KEY,
		county_id INT NOT NULL REFERENCES counties(id) ON DELETE CASCADE,
		name VARCHAR(100) NOT NULL,
		created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
		UNIQUE(county_id, name)
	);
	CREATE TABLE IF NOT EXISTS wards (
		id SERIAL PRIMARY KEY,
		sub_county_id INT NOT NULL REFERENCES sub_counties(id) ON DELETE CASCADE,
		name VARCHAR(100) NOT NULL,
		created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
		UNIQUE(sub_county_id, name)
	);
	CREATE TABLE IF NOT EXISTS towns (
		id SERIAL PRIMARY KEY,
		county_id INT NOT NULL REFERENCES counties(id) ON DELETE CASCADE,
		name VARCHAR(100) NOT NULL,
		created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
		UNIQUE(county_id, name)
	);
	CREATE TABLE IF NOT EXISTS neighborhoods (
		id SERIAL PRIMARY KEY,
		sub_county_id INT REFERENCES sub_counties(id) ON DELETE CASCADE,
		town_id INT REFERENCES towns(id) ON DELETE CASCADE,
		name VARCHAR(100) NOT NULL,
		created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
	);

	ALTER TABLE properties ADD COLUMN IF NOT EXISTS county_id INT REFERENCES counties(id) ON DELETE SET NULL;
	ALTER TABLE properties ADD COLUMN IF NOT EXISTS sub_county_id INT REFERENCES sub_counties(id) ON DELETE SET NULL;
	ALTER TABLE properties ADD COLUMN IF NOT EXISTS ward_id INT REFERENCES wards(id) ON DELETE SET NULL;
	ALTER TABLE properties ADD COLUMN IF NOT EXISTS town_id INT REFERENCES towns(id) ON DELETE SET NULL;
	ALTER TABLE properties ADD COLUMN IF NOT EXISTS neighborhood_id INT REFERENCES neighborhoods(id) ON DELETE SET NULL;
	`
	_, err := database.Exec(query)
	if err != nil {
		log.Printf("Migration execution warning: %v", err)
	} else {
		log.Println("Database schema check/migration completed.")
	}

	if err := services.SeedLocations(database); err != nil {
		log.Printf("Location seeding notice: %v", err)
	}
}
