package service

import "github.com/hiteshpattanayak-tw/courier_service/internal/app/model"

type CostEstimation struct {}

func ProvideCostEstimationService() CostEstimation {
	return CostEstimation{}
}


func (ce CostEstimation) CalculateAllPackagesCost(input model.InputFormat) float32 {
	packageTracker := make(map[string]bool)
	totalCost := float32(0)
	for _,pkg := range input.Packages {
		var cost float32
		if offerApplied,ok := packageTracker[pkg.Id]; ok && offerApplied {
			cost = pkg.GetCost(input.BaseDeliveryCost, 0)
		} else {
			validity, discount := pkg.IsOfferValid()
			packageTracker[pkg.Id] = validity
			cost = pkg.GetCost(input.BaseDeliveryCost, discount)
		}
		totalCost += cost
	}

	return totalCost
 }