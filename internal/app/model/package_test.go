package model_test

import (
	"fmt"
	"github.com/hiteshpattanayak-tw/courier_service/internal/app/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPackage_IsOfferValidShouldReturnTrueIfPackageWeightAndDistanceIsWithinRangeByOfferCode(t *testing.T) {
	testcases := map[string]model.Package{
		"pkg1": {
			Id:           "PKG1",
			Weight:       145,
			DistanceInKm: 105,
			OfferCode:    "OFR003",
		},
		"pkg2": {
			Id:           "PKG2",
			Weight:       190,
			DistanceInKm: 150,
			OfferCode:    "OFR001",
		},
		"pkg3": {
			Id:           "PKG3",
			Weight:       240,
			DistanceInKm: 55,
			OfferCode:    "OFR002",
		},
	}

	for _, pkg := range testcases {
		validity, _ := pkg.IsOfferValid()
		assert.True(t, validity, fmt.Sprintf("offer: weight = %d, distance = %d", pkg.Weight, pkg.DistanceInKm))
	}
}

func TestPackage_IsOfferCodeValidShouldReturnFalseIfPackageWeightAndDistanceIsNotWithinRangeByOfferCode(t *testing.T) {
	testcases := map[string]model.Package{
		"pkg1": {
			Id:           "PKG1",
			Weight:       251,
			DistanceInKm: 45,
			OfferCode:    "OFR003",
		},
		"pkg2": {
			Id:           "PKG2",
			Weight:       240,
			DistanceInKm: 201,
			OfferCode:    "OFR001",
		},
		"pkg3": {
			Id:           "PKG3",
			Weight:       240,
			DistanceInKm: 45,
			OfferCode:    "OFR002",
		},
		"pkg4": {
			Id:           "PKG4",
			Weight:       251,
			DistanceInKm: 51,
			OfferCode:    "OFR002",
		},
		"pkg5": {
			Id:           "PKG5",
			Weight:       68,
			DistanceInKm: 55,
			OfferCode:    "OFR001",
		},
		"pkg6": {
			Id:           "PKG5",
			Weight:       9,
			DistanceInKm: 55,
			OfferCode:    "OFR003",
		},
	}

	for _, pkg := range testcases {
		validity, _ := pkg.IsOfferValid()
		assert.False(t, validity, fmt.Sprintf("offer: weight = %d, distance = %d", pkg.Weight, pkg.DistanceInKm))
	}
}

func TestPackage_GetCost(t *testing.T) {
	baseDeliveryPrice := 100
	testcases := map[string]map[string]interface{}{
		"tc1": {
			"pkg": model.Package{
				Id:           "PKG1",
				Weight:       145,
				DistanceInKm: 105,
				OfferCode:    "OFR003",
			},
			"expectedCost": float32(1971.25),
		},
		"tc2": {
			"pkg": model.Package{
				Id:           "PKG2",
				Weight:       190,
				DistanceInKm: 150,
				OfferCode:    "OFR001",
			},
			"expectedCost": float32(2475),
		},
		"tc3": {
			"pkg": model.Package{
				Id:           "PKG3",
				Weight:       240,
				DistanceInKm: 55,
				OfferCode:    "OFR002",
			},
			"expectedCost": float32(2580.75),
		},
		"tc4": {
			"pkg": model.Package{
				Id:           "PKG4",
				Weight:       240,
				DistanceInKm: 45,
				OfferCode:    "OFR002",
			},
			"expectedCost": float32(2725),
		},
	}

	for _, tc := range testcases {
		pkg := tc["pkg"].(model.Package)
		expectedCost := tc["expectedCost"].(float32)
		actualCost := pkg.GetCost(baseDeliveryPrice, getDiscountByOffer(pkg))
		assert.Equal(t, expectedCost, actualCost)
	}
}

func getDiscountByOffer(pkg model.Package) float32 {
	_, discount := pkg.IsOfferValid()
	return discount
}
