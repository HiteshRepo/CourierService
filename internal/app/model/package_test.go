package model_test

import (
	"github.com/hiteshpattanayak-tw/courier_service/internal/app/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPackage_IsOfferCodeValidShouldReturnTrueIfOfferCodeIsValid(t *testing.T) {
	pkg := model.Package{
		Id:           "PKG1",
		Weight:       1,
		DistanceInKm: 1,
		OfferCode:    "OFR001",
	}

	assert.True(t, pkg.IsOfferCodeValid())
}

func TestPackage_IsOfferCodeValidShouldReturnFalseIfOfferCodeIsInvalid(t *testing.T) {
	pkg := model.Package{
		Id:           "PKG1",
		Weight:       1,
		DistanceInKm: 1,
		OfferCode:    "OFR008",
	}

	assert.False(t, pkg.IsOfferCodeValid())
}
