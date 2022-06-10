package service_test

import (
	"github.com/hiteshpattanayak-tw/courier_service/internal/app/model"
	"github.com/hiteshpattanayak-tw/courier_service/internal/app/service"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCostEstimation_CalculateAllPackagesCost(t *testing.T) {
	ceSvc := service.ProvideCostEstimationService()

	packages := []model.Package{
		{
			Id:           "PKG1",
			Weight:       145,
			DistanceInKm: 105,
			OfferCode:    "OFR003",
		},
		{
			Id:           "PKG2",
			Weight:       190,
			DistanceInKm: 150,
			OfferCode:    "OFR001",
		},
		{
			Id:           "PKG3",
			Weight:       240,
			DistanceInKm: 55,
			OfferCode:    "OFR002",
		},
		{
			Id:           "PKG4",
			Weight:       240,
			DistanceInKm: 45,
			OfferCode:    "OFR002",
		},
	}
	input := model.InputFormat{
		BaseDeliveryCost: 100,
		NoOfPackages:     4,
		Packages:         packages,
	}

	expected := float32(9752)
	actual := ceSvc.CalculateAllPackagesCost(input)

	assert.Equal(t, expected, actual)
}
