package model_test

import (
	"github.com/hiteshpattanayak-tw/courier_service/internal/app/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetOfferByCode(t *testing.T) {

	testcases := map[string]map[string]interface{}{
		"tc1": {
			"offerCode": "OFR001",
			"expected": model.Offer{
				Id: "OFR001",
				Distance: model.Distance{
					Min: 0,
					Max: 200,
				},
				Weight: model.Weight{
					Min: 70,
					Max: 200,
				},
				Discount: 0.1,
			},
		},

		"tc2": {
			"offerCode": "OFR002",
			"expected": model.Offer{
				Id: "OFR002",
				Distance: model.Distance{
					Min: 50,
					Max: 150,
				},
				Weight: model.Weight{
					Min: 100,
					Max: 250,
				},
				Discount: 0.07,
			},
		},

		"tc3": {
			"offerCode": "OFR003",
			"expected": model.Offer{
				Id: "OFR003",
				Distance: model.Distance{
					Min: 50,
					Max: 250,
				},
				Weight: model.Weight{
					Min: 10,
					Max: 150,
				},
				Discount: 0.05,
			},
		},
	}

	for _, tc := range testcases {
		offerCode := tc["offerCode"].(string)
		expected := tc["expected"].(model.Offer)
		actual := model.GetOfferByCode(offerCode)
		assert.Equal(t, expected, actual)
	}
}

func TestGetOfferByCode_ShouldReturnNillOfferForInvalidOfferCode(t *testing.T) {
	offerCode := "OFR008"
	assert.True(t, model.GetOfferByCode(offerCode).IsNilOffer())
}

func TestOffer_IsNilOffer(t *testing.T) {
	offer := model.Offer{}
	assert.True(t, offer.IsNilOffer())
}
