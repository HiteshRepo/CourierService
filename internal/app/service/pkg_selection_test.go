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

	pkgSortingSvc := service.ProvidePackageSelectionService()
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

	pkgSortingSvc := service.ProvidePackageSelectionService()
	pkgSortingSvc.SortPackages(pkgs)

	assert.Equal(t, expected, pkgs)
}

func TestPackageSorting_SelectPackages(t *testing.T) {
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

	limit := float32(200)

	expected := []model.Package{
		{
			Id:           "PKG4",
			Weight:       110,
			DistanceInKm: 60,
			OfferCode:    "OFR002",
		},
		{
			Id:           "PKG2",
			Weight:       75,
			DistanceInKm: 125,
			OfferCode:    "OFR008",
		},
	}

	pkgSvc := service.ProvidePackageSelectionService()
	actual := pkgSvc.SelectPackages(packages, limit)

	assert.ElementsMatch(t, expected, actual)
}
