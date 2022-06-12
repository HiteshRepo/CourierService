package main

import (
	"bufio"
	"fmt"
	"github.com/hiteshpattanayak-tw/courier_service/internal/app/model"
	"github.com/hiteshpattanayak-tw/courier_service/internal/app/service"
	"log"
	"os"
	"strconv"
	"strings"
)

const BITSIZE_32 = 32

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter base delivery cost and number of packages in the format: <base_delivery_cost> <no_of_packages>")
	input, _ := reader.ReadString('\n')

	parts := strings.Split(strings.TrimSpace(input), " ")
	if len(parts) != 2 {
		log.Fatal("invalid input, format should be: <base_delivery_cost> <no_of_packages>")
	}

	baseDeliveryCost, err := strconv.Atoi(parts[0])
	if err != nil {
		log.Fatal("invalid entry for base_delivery_cost")
	}

	noOfPackages, err := strconv.Atoi(parts[1])
	if err != nil {
		log.Fatal("invalid entry for no_of_packages")
	}

	packages := make([]model.Package, 0)
	for i := 0; i < noOfPackages; i++ {
		fmt.Println("Enter package info in the format: <pkg_id> <weight_in_kg> <distance_in_km> <offer_code>")
		input, _ = reader.ReadString('\n')

		parts = strings.Split(strings.TrimSpace(input), " ")
		if len(parts) != 4 {
			log.Fatal("invalid input, format should be: <pkg_id> <weight_in_kg> <distance_in_km> <offer_code>")
		}
		pkgId := parts[0]

		weight, err := strconv.Atoi(parts[1])
		if err != nil {
			log.Fatal("invalid entry for weight")
		}

		distance, err := strconv.Atoi(parts[2])
		if err != nil {
			log.Fatal("invalid entry for distance")
		}

		offerCode := parts[3]

		pkg := model.Package{
			Id:           pkgId,
			Weight:       weight,
			DistanceInKm: distance,
			OfferCode:    offerCode,
		}

		packages = append(packages, pkg)
	}

	inputPkgDetails := model.PackageInputFormat{
		BaseDeliveryCost: baseDeliveryCost,
		NoOfPackages:     noOfPackages,
		Packages:         packages,
	}

	fmt.Println("Enter vehicle details in the format: <no_of_vehicles> <max_speed_limit> <max_carry_weight>")
	input, _ = reader.ReadString('\n')

	parts = strings.Split(strings.TrimSpace(input), " ")
	if len(parts) != 3 {
		log.Fatal("invalid input, format should be: <no_of_vehicles> <max_speed_limit> <max_carry_weight>")
	}

	noOfVehicles, err := strconv.Atoi(parts[0])
	if err != nil {
		log.Fatal("invalid entry for no_of_vehicles")
	}

	speedLimit, err := strconv.Atoi(parts[1])
	if err != nil {
		log.Fatal("invalid entry for max_speed_limit")
	}

	weightLimit, err := strconv.Atoi(parts[2])
	if err != nil {
		log.Fatal("invalid entry for max_carry_weight")
	}

	vehicles := make([]model.Vehicle, noOfVehicles)

	for i := 0; i < noOfVehicles; i++ {
		v := model.Vehicle{
			Id:                fmt.Sprintf("VH%d", i),
			MaxWeightLimit:    float32(weightLimit),
			MaxSpeedLimit:     float32(speedLimit),
			Shipments:         make([]model.Shipment, 0),
			NextAvailableTime: 0,
		}

		vehicles[i] = v
	}

	ceSvc := service.ProvideCostEstimationService()
	output := ceSvc.CalculateAllPackagesCost(inputPkgDetails)

	psSvc := service.ProvidePackageSelectionService()
	deSvc := service.ProvideDeliveryEstimationService(psSvc)
	pkgDeliveryTimings := deSvc.FetchDeliveryEstimations(packages, vehicles)

	for i,p := range output.Packages {
		output.Packages[i].DeliveryTimeInHrs = pkgDeliveryTimings[p.Id]
	}

	for _, pkg := range output.Packages {
		fmt.Printf("%s %s %s %f\n", pkg.Id, formatValue(pkg.Discount), formatValue(pkg.TotalCost), formatValue(pkg.DeliveryTimeInHrs))
	}
}

func formatValue(value float32) string {
	return strconv.FormatFloat(float64(value), 'f', -1, BITSIZE_32)
}

// 100 5
// PKG1 50 30 OFR001
// PKG2 75 125 OFR008
// PKG3 175 100 OFR003
// PKG4 110 60 OFR002
// PKG5 155 95 0
// 2 70 200
