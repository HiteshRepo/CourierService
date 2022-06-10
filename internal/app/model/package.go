package model

type Package struct {
	Id           string `json:"id"`
	Weight       int    `json:"weight"`
	DistanceInKm int    `json:"distance_in_km"`
	OfferCode    string `json:"offer_code"`
}
