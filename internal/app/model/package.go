package model

import "strings"

var ValidOfferCodes = []string{"OFR001", "OFR002", "OFR003"}

type Package struct {
	Id           string `json:"id"`
	Weight       int    `json:"weight"`
	DistanceInKm int    `json:"distance_in_km"`
	OfferCode    string `json:"offer_code"`
}

func (p Package) IsOfferCodeValid() bool {
	for _,o := range ValidOfferCodes {
		if o == strings.TrimSpace(p.OfferCode) {
			return true
		}
	}

	return false
}
