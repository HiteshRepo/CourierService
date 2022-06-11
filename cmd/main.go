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

	ceSvc := service.ProvideCostEstimationService()
	output := ceSvc.CalculateAllPackagesCost(inputPkgDetails)

	for _, pkg := range output.Packages {
		fmt.Printf("%s %s %s\n", pkg.Id, formatValue(pkg.Discount), formatValue(pkg.TotalCost))
	}
}

func formatValue(value float32) string {
	return strconv.FormatFloat(float64(value), 'f', -1, BITSIZE_32)
}
