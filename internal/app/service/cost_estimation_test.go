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
		{
			Id:           "PKG1",
			Weight:       145,
			DistanceInKm: 105,
			OfferCode:    "OFR002",
		},
	}
	input := model.PackageInputFormat{
		BaseDeliveryCost: 100,
		NoOfPackages:     4,
		Packages:         packages,
	}

	expected := model.PackageOutputFormat{Packages: []model.PackageOutput{
		{
			Id:        "PKG1",
			Discount:  103.75,
			TotalCost: 1971.25,
		},
		{
			Id:        "PKG2",
			Discount:  275,
			TotalCost: 2475,
		},
		{
			Id:        "PKG3",
			Discount:  194.25,
			TotalCost: 2580.75,
		},
		{
			Id:        "PKG4",
			Discount:  0,
			TotalCost: 2725,
		},
	}}
	actual := ceSvc.CalculateAllPackagesCost(input)

	assert.Equal(t, expected, actual)
}
