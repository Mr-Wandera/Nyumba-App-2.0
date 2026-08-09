package handlers

import (
	"encoding/json"
	"net/http"
	"nyumba/db"
	"nyumba/services"
	"strconv"
)

// GetCountiesHandler handles GET /api/locations/counties
func GetCountiesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	database := db.GetDB()
	counties, err := services.GetCounties(database)
	if err != nil {
		http.Error(w, "Failed to fetch counties: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(counties)
}

// GetSubCountiesHandler handles GET /api/locations/sub-counties?county_id=X
func GetSubCountiesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	countyIDStr := r.URL.Query().Get("county_id")
	if countyIDStr == "" {
		http.Error(w, "county_id is required", http.StatusBadRequest)
		return
	}

	countyID, err := strconv.Atoi(countyIDStr)
	if err != nil {
		http.Error(w, "invalid county_id parameter", http.StatusBadRequest)
		return
	}

	database := db.GetDB()
	subs, err := services.GetSubCounties(database, countyID)
	if err != nil {
		http.Error(w, "Failed to fetch sub-counties: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(subs)
}

// GetWardsHandler handles GET /api/locations/wards?sub_county_id=X
func GetWardsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	subIDStr := r.URL.Query().Get("sub_county_id")
	if subIDStr == "" {
		http.Error(w, "sub_county_id is required", http.StatusBadRequest)
		return
	}

	subID, err := strconv.Atoi(subIDStr)
	if err != nil {
		http.Error(w, "invalid sub_county_id parameter", http.StatusBadRequest)
		return
	}

	database := db.GetDB()
	wards, err := services.GetWards(database, subID)
	if err != nil {
		http.Error(w, "Failed to fetch wards: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(wards)
}

// GetTownsHandler handles GET /api/locations/towns?county_id=X
func GetTownsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	countyIDStr := r.URL.Query().Get("county_id")
	if countyIDStr == "" {
		http.Error(w, "county_id is required", http.StatusBadRequest)
		return
	}

	countyID, err := strconv.Atoi(countyIDStr)
	if err != nil {
		http.Error(w, "invalid county_id parameter", http.StatusBadRequest)
		return
	}

	database := db.GetDB()
	towns, err := services.GetTowns(database, countyID)
	if err != nil {
		http.Error(w, "Failed to fetch towns: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(towns)
}

// GetNeighborhoodsHandler handles GET /api/locations/neighborhoods?sub_county_id=X&town_id=Y
func GetNeighborhoodsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	subID, _ := strconv.Atoi(r.URL.Query().Get("sub_county_id"))
	townID, _ := strconv.Atoi(r.URL.Query().Get("town_id"))

	database := db.GetDB()
	neighs, err := services.GetNeighborhoods(database, subID, townID)
	if err != nil {
		http.Error(w, "Failed to fetch neighborhoods: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(neighs)
}

// GetLocationHierarchyHandler handles GET /api/locations/hierarchy
func GetLocationHierarchyHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(services.KenyaLocationsData)
}
