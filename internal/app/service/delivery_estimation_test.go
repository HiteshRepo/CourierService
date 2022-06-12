package service_test

import (
	"github.com/hiteshpattanayak-tw/courier_service/internal/app/model"
	"github.com/hiteshpattanayak-tw/courier_service/internal/app/service"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDeliveryEstimationService_UpdateDeliveryEstimations(t *testing.T) {
	vehicles := []model.Vehicle{
		{
			Id:                "VH1",
			MaxSpeedLimit:     float32(70),
			MaxWeightLimit:    float32(200),
			Shipments:         make([]model.Shipment, 0),
			NextAvailableTime: float32(0),
		},
		{
			Id:                "VH2",
			MaxSpeedLimit:     float32(70),
			MaxWeightLimit:    float32(200),
			Shipments:         make([]model.Shipment, 0),
			NextAvailableTime: float32(0),
		},
	}

	packages := []model.Package{
		{
			Id:           "PKG1",
			Weight:       50,
			DistanceInKm: 30,
			OfferCode:    "OFR001",
		},
		{
			Id:           "PKG2",
			Weight:       75,
			DistanceInKm: 125,
			OfferCode:    "OFR008",
		},
		{
			Id:           "PKG3",
			Weight:       175,
			DistanceInKm: 100,
			OfferCode:    "OFR003",
		},
		{
			Id:           "PKG4",
			Weight:       110,
			DistanceInKm: 60,
			OfferCode:    "OFR002",
		},
		{
			Id:           "PKG5",
			Weight:       155,
			DistanceInKm: 95,
			OfferCode:    "",
		},
	}

	pkgSortSvc := service.ProvidePackageSelectionService()
	deSvc := service.ProvideDeliveryEstimationService(pkgSortSvc)

	expected := map[string]float32{
		"PKG1": float32(3.98),
		"PKG2": float32(1.78),
		"PKG3": float32(1.42),
		"PKG4": float32(0.85),
		"PKG5": float32(4.19),
	}

	actual := deSvc.UpdateDeliveryEstimations(packages, vehicles)
	assert.Equal(t, expected, actual)
}
