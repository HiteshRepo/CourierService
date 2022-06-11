package model

type PackageOutputFormat struct {
	Packages []PackageOutput
}

type PackageOutput struct {
	Id        string  `json:"id"`
	Discount  float32 `json:"discount"`
	TotalCost float32 `json:"total_cost"`
}
