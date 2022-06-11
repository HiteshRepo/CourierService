package service

import "github.com/hiteshpattanayak-tw/courier_service/internal/app/model"

type CostEstimation struct {}

func ProvideCostEstimationService() CostEstimation {
	return CostEstimation{}
}


func (ce CostEstimation) CalculateAllPackagesCost(input model.InputFormat) model.OutputFormat {
	output := model.OutputFormat{}
	packageTracker := make(map[string]bool)

	for _,pkg := range input.Packages {
		if offerApplied,ok := packageTracker[pkg.Id]; ok || offerApplied {
			continue
		}

		validity, discountPercent := pkg.IsOfferValid()
		packageTracker[pkg.Id] = validity
		cost, discount := pkg.GetCost(input.BaseDeliveryCost, discountPercent)


		pkgOut := model.PackageOutput{
			Id:        pkg.Id,
			Discount:  discount,
			TotalCost: cost,
		}

		output.Packages = append(output.Packages, pkgOut)
	}

	return output
 }