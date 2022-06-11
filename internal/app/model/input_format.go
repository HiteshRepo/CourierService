package model

type PackageInputFormat struct {
	BaseDeliveryCost int       `json:"base_delivery_cost"`
	NoOfPackages     int       `json:"no_of_packages"`
	Packages         []Package `json:"packages"`
}

type VehicleInputFormat struct {
	NumberOfVehicles int       `json:"number_of_vehicles"`
	MaxSpeed         int       `json:"max_speed"`
	MaxCarryWeight   []Package `json:"max_carriable_weight"`
}
