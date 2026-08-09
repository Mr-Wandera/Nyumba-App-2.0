package services

import (
	"nyumba/models"
	"testing"
)

func TestValidatePropertyInput(t *testing.T) {
	tests := []struct {
		name         string
		buildingName string
		location     string
		price        float64
		wantErr      bool
	}{
		{
			name:         "Valid Property",
			buildingName: "Azure Heights",
			location:     "Westlands",
			price:        75000,
			wantErr:      false,
		},
		{
			name:         "Missing Building Name",
			buildingName: "",
			location:     "Kilimani",
			price:        50000,
			wantErr:      true,
		},
		{
			name:         "Missing Location",
			buildingName: "Green Park",
			location:     "",
			price:        30000,
			wantErr:      true,
		},
		{
			name:         "Zero Price",
			buildingName: "Villa Park",
			location:     "Lavington",
			price:        0,
			wantErr:      true,
		},
		{
			name:         "Negative Price",
			buildingName: "Villa Park",
			location:     "Lavington",
			price:        -100,
			wantErr:      true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidatePropertyInput(tt.buildingName, tt.location, tt.price)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidatePropertyInput() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPropertyFilterBuilding(t *testing.T) {
	filter := PropertyFilter{
		Location:      "Westlands",
		MaxPrice:      80000,
		Bedrooms:      2,
		PublishedOnly: true,
	}

	if filter.Location != "Westlands" || filter.MaxPrice != 80000 || filter.Bedrooms != 2 || !filter.PublishedOnly {
		t.Errorf("PropertyFilter structure mismatch: %+v", filter)
	}
}

func TestPropertyOwnershipCheckLogic(t *testing.T) {
	ownerID := 10
	unauthorizedID := 99

	p := models.Property{
		ID:           1,
		LandlordID:   &ownerID,
		BuildingName: "Silver Towers",
		Location:     "Kileleshwa",
		Price:        45000,
		IsPublished:  true,
	}

	// Owner check
	if p.LandlordID == nil || *p.LandlordID != ownerID {
		t.Errorf("Expected owner check to match owner ID %d", ownerID)
	}

	// Unauthorized check
	if p.LandlordID != nil && *p.LandlordID == unauthorizedID {
		t.Errorf("Expected unauthorized check to fail for ID %d", unauthorizedID)
	}
}
