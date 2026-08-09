package services

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"nyumba/models"
	"strings"
)

var (
	ErrCountyNotFound       = errors.New("county not found")
	ErrSubCountyNotFound    = errors.New("sub-county not found")
	ErrInvalidLocationMatch = errors.New("invalid location hierarchy combination")
)

// SeedLocations populates PostgreSQL with Kenya's 47 counties and hierarchical data
func SeedLocations(db *sql.DB) error {
	if db == nil {
		return nil
	}

	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("failed to start transaction: %w", err)
	}
	defer tx.Rollback()

	for _, countySeed := range KenyaLocationsData {
		var countyID int
		err := tx.QueryRow(`
			INSERT INTO counties (code, name)
			VALUES ($1, $2)
			ON CONFLICT (code) DO UPDATE SET name = EXCLUDED.name
			RETURNING id
		`, countySeed.Code, countySeed.Name).Scan(&countyID)
		if err != nil {
			return fmt.Errorf("failed to seed county %s: %w", countySeed.Name, err)
		}

		// Seed Towns
		for _, townName := range countySeed.Towns {
			var townID int
			err := tx.QueryRow(`
				INSERT INTO towns (county_id, name)
				VALUES ($1, $2)
				ON CONFLICT (county_id, name) DO UPDATE SET name = EXCLUDED.name
				RETURNING id
			`, countyID, townName).Scan(&townID)
			if err != nil {
				log.Printf("town insert notice %s: %v", townName, err)
			}
		}

		// Seed SubCounties, Wards, Neighborhoods
		for _, subSeed := range countySeed.SubCounties {
			var subID int
			err := tx.QueryRow(`
				INSERT INTO sub_counties (county_id, name)
				VALUES ($1, $2)
				ON CONFLICT (county_id, name) DO UPDATE SET name = EXCLUDED.name
				RETURNING id
			`, countyID, subSeed.Name).Scan(&subID)
			if err != nil {
				return fmt.Errorf("failed to seed sub-county %s: %w", subSeed.Name, err)
			}

			for _, wardName := range subSeed.Wards {
				_, err := tx.Exec(`
					INSERT INTO wards (sub_county_id, name)
					VALUES ($1, $2)
					ON CONFLICT (sub_county_id, name) DO NOTHING
				`, subID, wardName)
				if err != nil {
					log.Printf("ward insert notice %s: %v", wardName, err)
				}
			}

			for _, neighName := range subSeed.Neighborhoods {
				_, err := tx.Exec(`
					INSERT INTO neighborhoods (sub_county_id, name)
					VALUES ($1, $2)
				`, subID, neighName)
				if err != nil {
					log.Printf("neighborhood insert notice %s: %v", neighName, err)
				}
			}
		}
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit location seed: %w", err)
	}

	log.Println("Successfully seeded all 47 Kenyan counties and administrative hierarchy into PostgreSQL")
	return nil
}

// GetCounties returns all 47 Kenyan counties
func GetCounties(db *sql.DB) ([]models.County, error) {
	if db == nil {
		// Fallback to memory dataset
		var counties []models.County
		for idx, c := range KenyaLocationsData {
			counties = append(counties, models.County{
				ID:   idx + 1,
				Code: c.Code,
				Name: c.Name,
			})
		}
		return counties, nil
	}

	rows, err := db.Query("SELECT id, code, name FROM counties ORDER BY code ASC")
	if err != nil {
		return nil, fmt.Errorf("failed to query counties: %w", err)
	}
	defer rows.Close()

	var counties []models.County
	for rows.Next() {
		var c models.County
		if err := rows.Scan(&c.ID, &c.Code, &c.Name); err != nil {
			return nil, err
		}
		counties = append(counties, c)
	}

	if len(counties) == 0 {
		// Fallback
		return GetCounties(nil)
	}

	return counties, nil
}

// GetSubCounties returns sub-counties for a given county
func GetSubCounties(db *sql.DB, countyID int) ([]models.SubCounty, error) {
	if db == nil {
		var subs []models.SubCounty
		if countyID > 0 && countyID <= len(KenyaLocationsData) {
			c := KenyaLocationsData[countyID-1]
			for sIdx, s := range c.SubCounties {
				subs = append(subs, models.SubCounty{
					ID:       (countyID * 100) + sIdx + 1,
					CountyID: countyID,
					Name:     s.Name,
				})
			}
		}
		return subs, nil
	}

	rows, err := db.Query("SELECT id, county_id, name FROM sub_counties WHERE county_id = $1 ORDER BY name ASC", countyID)
	if err != nil {
		return nil, fmt.Errorf("failed to query sub-counties: %w", err)
	}
	defer rows.Close()

	var subs []models.SubCounty
	for rows.Next() {
		var s models.SubCounty
		if err := rows.Scan(&s.ID, &s.CountyID, &s.Name); err != nil {
			return nil, err
		}
		subs = append(subs, s)
	}
	return subs, nil
}

// GetWards returns wards for a sub-county
func GetWards(db *sql.DB, subCountyID int) ([]models.Ward, error) {
	if db == nil {
		return []models.Ward{}, nil
	}

	rows, err := db.Query("SELECT id, sub_county_id, name FROM wards WHERE sub_county_id = $1 ORDER BY name ASC", subCountyID)
	if err != nil {
		return nil, fmt.Errorf("failed to query wards: %w", err)
	}
	defer rows.Close()

	var wards []models.Ward
	for rows.Next() {
		var w models.Ward
		if err := rows.Scan(&w.ID, &w.SubCountyID, &w.Name); err != nil {
			return nil, err
		}
		wards = append(wards, w)
	}
	return wards, nil
}

// GetTowns returns towns for a county
func GetTowns(db *sql.DB, countyID int) ([]models.Town, error) {
	if db == nil {
		var towns []models.Town
		if countyID > 0 && countyID <= len(KenyaLocationsData) {
			c := KenyaLocationsData[countyID-1]
			for tIdx, tName := range c.Towns {
				towns = append(towns, models.Town{
					ID:       (countyID * 100) + tIdx + 1,
					CountyID: countyID,
					Name:     tName,
				})
			}
		}
		return towns, nil
	}

	rows, err := db.Query("SELECT id, county_id, name FROM towns WHERE county_id = $1 ORDER BY name ASC", countyID)
	if err != nil {
		return nil, fmt.Errorf("failed to query towns: %w", err)
	}
	defer rows.Close()

	var towns []models.Town
	for rows.Next() {
		var t models.Town
		if err := rows.Scan(&t.ID, &t.CountyID, &t.Name); err != nil {
			return nil, err
		}
		towns = append(towns, t)
	}
	return towns, nil
}

// GetNeighborhoods returns neighborhoods for a sub-county or town
func GetNeighborhoods(db *sql.DB, subCountyID, townID int) ([]models.Neighborhood, error) {
	if db == nil {
		return []models.Neighborhood{}, nil
	}

	query := "SELECT id, sub_county_id, town_id, name FROM neighborhoods WHERE 1=1"
	args := []interface{}{}
	argIdx := 1

	if subCountyID > 0 {
		query += fmt.Sprintf(" AND sub_county_id = $%d", argIdx)
		args = append(args, subCountyID)
		argIdx++
	}
	if townID > 0 {
		query += fmt.Sprintf(" AND town_id = $%d", argIdx)
		args = append(args, townID)
		argIdx++
	}

	query += " ORDER BY name ASC"

	rows, err := db.Query(query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to query neighborhoods: %w", err)
	}
	defer rows.Close()

	var neighborhoods []models.Neighborhood
	for rows.Next() {
		var n models.Neighborhood
		if err := rows.Scan(&n.ID, &n.SubCountyID, &n.TownID, &n.Name); err != nil {
			return nil, err
		}
		neighborhoods = append(neighborhoods, n)
	}
	return neighborhoods, nil
}

// ValidateLocationHierarchy verifies geographic relationships
func ValidateLocationHierarchy(db *sql.DB, countyID, subCountyID, wardID, townID, neighborhoodID *int) error {
	if db == nil {
		return nil
	}

	if countyID != nil && *countyID > 0 {
		var exists bool
		err := db.QueryRow("SELECT EXISTS(SELECT 1 FROM counties WHERE id = $1)", *countyID).Scan(&exists)
		if err != nil || !exists {
			return ErrCountyNotFound
		}
	}

	if subCountyID != nil && *subCountyID > 0 {
		if countyID == nil || *countyID <= 0 {
			return errors.New("sub-county provided without valid county")
		}
		var parentCountyID int
		err := db.QueryRow("SELECT county_id FROM sub_counties WHERE id = $1", *subCountyID).Scan(&parentCountyID)
		if err != nil {
			return ErrSubCountyNotFound
		}
		if parentCountyID != *countyID {
			return fmt.Errorf("%w: sub-county %d does not belong to county %d", ErrInvalidLocationMatch, *subCountyID, *countyID)
		}
	}

	if wardID != nil && *wardID > 0 {
		if subCountyID == nil || *subCountyID <= 0 {
			return errors.New("ward provided without valid sub-county")
		}
		var parentSubCountyID int
		err := db.QueryRow("SELECT sub_county_id FROM wards WHERE id = $1", *wardID).Scan(&parentSubCountyID)
		if err != nil {
			return errors.New("ward not found")
		}
		if parentSubCountyID != *subCountyID {
			return fmt.Errorf("%w: ward %d does not belong to sub-county %d", ErrInvalidLocationMatch, *wardID, *subCountyID)
		}
	}

	if townID != nil && *townID > 0 {
		if countyID == nil || *countyID <= 0 {
			return errors.New("town provided without valid county")
		}
		var parentCountyID int
		err := db.QueryRow("SELECT county_id FROM towns WHERE id = $1", *townID).Scan(&parentCountyID)
		if err != nil {
			return errors.New("town not found")
		}
		if parentCountyID != *countyID {
			return fmt.Errorf("%w: town %d does not belong to county %d", ErrInvalidLocationMatch, *townID, *countyID)
		}
	}

	return nil
}

// BuildLocationString constructs a human-readable string from hierarchy IDs or names
func BuildLocationString(db *sql.DB, countyID, subCountyID, wardID, townID, neighborhoodID *int, defaultLoc string) string {
	if db == nil {
		if defaultLoc != "" {
			return defaultLoc
		}
		return "Kenya"
	}

	parts := []string{}

	if neighborhoodID != nil && *neighborhoodID > 0 {
		var name string
		if db.QueryRow("SELECT name FROM neighborhoods WHERE id = $1", *neighborhoodID).Scan(&name) == nil {
			parts = append(parts, name)
		}
	}

	if wardID != nil && *wardID > 0 {
		var name string
		if db.QueryRow("SELECT name FROM wards WHERE id = $1", *wardID).Scan(&name) == nil {
			parts = append(parts, name)
		}
	}

	if townID != nil && *townID > 0 {
		var name string
		if db.QueryRow("SELECT name FROM towns WHERE id = $1", *townID).Scan(&name) == nil {
			parts = append(parts, name)
		}
	}

	if subCountyID != nil && *subCountyID > 0 {
		var name string
		if db.QueryRow("SELECT name FROM sub_counties WHERE id = $1", *subCountyID).Scan(&name) == nil {
			parts = append(parts, name)
		}
	}

	if countyID != nil && *countyID > 0 {
		var name string
		if db.QueryRow("SELECT name FROM counties WHERE id = $1", *countyID).Scan(&name) == nil {
			parts = append(parts, name)
		}
	}

	if len(parts) > 0 {
		return strings.Join(parts, ", ")
	}

	return defaultLoc
}
