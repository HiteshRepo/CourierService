package model

type Offer struct {
	Id       string
	Distance Distance
	Weight   Weight
	Discount float32
}

type Distance struct {
	Min int
	Max int
}

type Weight struct {
	Min int
	Max int
}

func (o Offer) IsNilOffer() bool {
	return o.Id == ""
}

func GetOfferByCode(code string) Offer {
	switch code {
	case "OFR001":
		return Offer{
			Id:       code,
			Distance: Distance{Min: 0, Max: 200},
			Weight:   Weight{Min: 70, Max: 200},
			Discount: 0.1,
		}
	case "OFR002":
		return Offer{
			Id:       code,
			Distance: Distance{Min: 50, Max: 150},
			Weight:   Weight{Min: 100, Max: 250},
			Discount: 0.07,
		}
	case "OFR003":
		return Offer{
			Id:       code,
			Distance: Distance{Min: 50, Max: 250},
			Weight:   Weight{Min: 10, Max: 150},
			Discount: 0.05,
		}
	default:
		return Offer{}
	}
}
