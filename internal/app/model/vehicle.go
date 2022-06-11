package model

type Vehicle struct {
	Id                string
	MaxWeightLimit    float32
	MaxSpeedLimit     float32
	Shipments         []Shipment
	NextAvailableTime float32
}
