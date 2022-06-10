package model

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
