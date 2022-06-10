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
		assert.True(t, pkg.IsOfferValid(), fmt.Sprintf("offer: weight = %d, distance = %d", pkg.Weight, pkg.DistanceInKm))
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
		assert.False(t, pkg.IsOfferValid(), fmt.Sprintf("offer: weight = %d, distance = %d", pkg.Weight, pkg.DistanceInKm))
	}
}