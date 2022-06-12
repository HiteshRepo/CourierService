package service

import (
	"github.com/hiteshpattanayak-tw/courier_service/internal/app/model"
	"strconv"
	"strings"
)

const BITSIZE_32 = 32

type DeliveryEstimationService struct {
	pkgSortingSvc PackageSelectionService
}

func ProvideDeliveryEstimationService(pkgSortingSvc PackageSelectionService) DeliveryEstimationService {
	return DeliveryEstimationService{pkgSortingSvc: pkgSortingSvc}
}

func (des DeliveryEstimationService) FetchDeliveryEstimations(pkgs []model.Package, vehicles []model.Vehicle) map[string]float32 {
	outputPkgsMap := make(map[string]float32)

	copyPkgs := make([]model.Package, len(pkgs))
	copy(copyPkgs, pkgs)
	currVehicle := 0
	currTime := float32(0)

	for len(copyPkgs) > 0 {
		selectedPkgs := des.pkgSortingSvc.SelectPackages(copyPkgs, vehicles[currVehicle].MaxWeightLimit)
		if selectedPkgs == nil {
			break
		}

		for _, sp := range selectedPkgs {
			outputPkgsMap[sp.Id] = currTime + des.formatValue(float32(sp.DistanceInKm)/vehicles[currVehicle].MaxSpeedLimit)
		}

		shipment := model.Shipment{Packages: selectedPkgs, Time: currTime}
		vehicles[currVehicle].Shipments = append(vehicles[currVehicle].Shipments, shipment)
		vehicles[currVehicle].NextAvailableTime = currTime + (des.formatValue(des.getHighestShipmentDeliveryTime(vehicles[currVehicle].MaxSpeedLimit, shipment)) * 2)

		copyPkgs = des.getRemainingPkgs(copyPkgs, selectedPkgs)
		currVehicle += 1

		if currVehicle == len(vehicles) {
			minVehicle := 0
			minTime := vehicles[0].NextAvailableTime
			if vehicles[0].NextAvailableTime > vehicles[1].NextAvailableTime {
				minVehicle = 1
				minTime = vehicles[1].NextAvailableTime
			}

			for i := 1; i < len(vehicles); i++ {
				if vehicles[minVehicle].NextAvailableTime > vehicles[i].NextAvailableTime {
					minVehicle = i
					minTime = vehicles[i].NextAvailableTime
				}
			}

			currTime = minTime
			currVehicle = minVehicle
		}
	}

	return outputPkgsMap
}

func (des DeliveryEstimationService) getHighestShipmentDeliveryTime(maxSpeedLimit float32, shipment model.Shipment) float32 {
	max := float32(0)
	for _, p := range shipment.Packages {
		currShipmentTime := float32(p.DistanceInKm) / maxSpeedLimit
		max = des.max(max, currShipmentTime)
	}

	return max
}

func (des DeliveryEstimationService) max(a, b float32) float32 {
	if a > b {
		return a
	}
	return b
}

func (des DeliveryEstimationService) getLowestNextDeliveryVehicle(vehicles []model.Vehicle) int {
	lowest := 0

	for i, v := range vehicles {
		if v.NextAvailableTime < vehicles[lowest].NextAvailableTime {
			lowest = i
		}
	}

	return lowest
}

func (des DeliveryEstimationService) getOutputPkgMap(outputPkgs []model.PackageOutput) map[string]model.PackageOutput {
	outputPkgMap := make(map[string]model.PackageOutput)

	for _, p := range outputPkgs {
		outputPkgMap[p.Id] = p
	}

	return outputPkgMap
}

func (des DeliveryEstimationService) getRemainingPkgs(pkgs []model.Package, currPkgs []model.Package) []model.Package {
	remainingPkgs := make([]model.Package, 0)
	for _, p := range pkgs {
		if !des.isPackagePresent(p, currPkgs) {
			remainingPkgs = append(remainingPkgs, p)
		}
	}

	return remainingPkgs
}

func (des DeliveryEstimationService) isPackagePresent(pkg model.Package, currPkgs []model.Package) bool {
	for _, p := range currPkgs {
		if p.Id == pkg.Id {
			return true
		}
	}

	return false
}

func (des DeliveryEstimationService) formatValue(value float32) float32 {
	strFloat := strconv.FormatFloat(float64(value), 'f', -1, BITSIZE_32)
	strFloatParts := strings.Split(strFloat, ".")
	strFloatParts[1] = strFloatParts[1][0:2]
	strFloat = strings.Join(strFloatParts, ".")
	res, _ := strconv.ParseFloat(strFloat, BITSIZE_32)
	return float32(res)
}
