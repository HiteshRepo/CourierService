package model

const (
	WeightMultiplier   = 10
	DistanceMultiplier = 5
)

type Package struct {
	Id           string `json:"id"`
	Weight       int    `json:"weight"`
	DistanceInKm int    `json:"distance_in_km"`
	OfferCode    string `json:"offer_code"`
}

func (p Package) IsOfferValid() (bool, float32) {
	validOfferByCode := GetOfferByCode(p.OfferCode)

	if validOfferByCode.IsNilOffer() {
		return false, 0
	}

	if !(p.DistanceInKm >= validOfferByCode.Distance.Min && p.DistanceInKm <= validOfferByCode.Distance.Max) {
		return false, 0
	}

	if !(p.Weight >= validOfferByCode.Weight.Min && p.Weight <= validOfferByCode.Weight.Max) {
		return false, 0
	}

	return true, validOfferByCode.Discount
}

func (p Package) GetCost(baseDeliveryCost int, discountPercent float32) (float32, float32) {
	costWithoutDiscount := baseDeliveryCost + (p.Weight * WeightMultiplier) + (p.DistanceInKm * DistanceMultiplier)
	discount := float32(costWithoutDiscount) * discountPercent
	return float32(costWithoutDiscount) - discount, discount
}
