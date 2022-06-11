package service_test

import (
	"github.com/hiteshpattanayak-tw/courier_service/internal/app/model"
	"github.com/hiteshpattanayak-tw/courier_service/internal/app/service"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSortPackages_ShouldSortInDescendingOrderByWeight(t *testing.T) {
	pkgs := []model.Package{
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
	}

	expected := []model.Package{
		{
			Id:           "PKG3",
			Weight:       240,
			DistanceInKm: 55,
			OfferCode:    "OFR002",
		},
		{
			Id:           "PKG2",
			Weight:       190,
			DistanceInKm: 150,
			OfferCode:    "OFR001",
		},
		{
			Id:           "PKG1",
			Weight:       145,
			DistanceInKm: 105,
			OfferCode:    "OFR003",
		},
	}

	pkgSortingSvc := service.ProvidePackageSortingService()
	pkgSortingSvc.SortPackages(pkgs)

	assert.Equal(t, expected, pkgs)
}

func TestSortPackages_ShouldSortInDescendingOrderByWeightAndIncreasingOrderByDistanceIfWeightMatch(t *testing.T) {
	pkgs := []model.Package{
		{
			Id:           "PKG1",
			Weight:       145,
			DistanceInKm: 105,
			OfferCode:    "OFR003",
		},
		{
			Id:           "PKG2",
			Weight:       145,
			DistanceInKm: 150,
			OfferCode:    "OFR001",
		},
		{
			Id:           "PKG3",
			Weight:       240,
			DistanceInKm: 55,
			OfferCode:    "OFR002",
		},
	}

	expected := []model.Package{
		{
			Id:           "PKG3",
			Weight:       240,
			DistanceInKm: 55,
			OfferCode:    "OFR002",
		},
		{
			Id:           "PKG1",
			Weight:       145,
			DistanceInKm: 105,
			OfferCode:    "OFR003",
		},
		{
			Id:           "PKG2",
			Weight:       145,
			DistanceInKm: 150,
			OfferCode:    "OFR001",
		},
	}

	pkgSortingSvc := service.ProvidePackageSortingService()
	pkgSortingSvc.SortPackages(pkgs)

	assert.Equal(t, expected, pkgs)
}
