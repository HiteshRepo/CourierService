package model

const (
	WeightMultiplier = 10
	DistanceMultiplier = 5
)

type Package struct {
	Id           string `json:"id"`
	Weight       int    `json:"weight"`
	DistanceInKm int    `json:"distance_in_km"`
	OfferCode    string `json:"offer_code"`
}

func (p Package) IsOfferValid() bool {
	validOfferByCode := GetOfferByCode(p.OfferCode)

	if validOfferByCode.IsNilOffer() {
		return false
	}

	if !(p.DistanceInKm > validOfferByCode.Distance.Min && p.DistanceInKm < validOfferByCode.Distance.Max) {
		return false
	}

	if !(p.Weight > validOfferByCode.Weight.Min && p.Weight < validOfferByCode.Weight.Max) {
		return false
	}

	return true
}

func (p Package) GetCost(baseDeliveryCost int, discount float32) float32 {
	costWithoutDiscount := baseDeliveryCost + (p.Weight * WeightMultiplier) + (p.DistanceInKm * DistanceMultiplier)
	return float32(costWithoutDiscount) * (1 - discount)
}
