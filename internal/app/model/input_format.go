package model

type InputFormat struct {
	BaseDeliveryCost int       `json:"base_delivery_cost"`
	NoOfPackages     int       `json:"no_of_packages"`
	Packages         []Package `json:"packages"`
}