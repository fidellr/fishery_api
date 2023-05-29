package model

type Commodity struct {
	UUID         string `json:"uuid"`
	Komoditas    string `json:"komoditas" validate:"required"`
	AreaProvinsi string `json:"area_provinsi"`
	AreaKota     string `json:"area_kota"`
	Size         string `json:"size" validate:"required"`
	Price        string `json:"price" validate:"required"`
	TglParsed    string `json:"tgl_parsed"`
	Timestamp    string `json:"timestamp"`
	USD          string `json:"price-usd,omitempty"`
}

type SizeOption struct {
	Size string `json:"size" validate:"required"`
}

type AreaOption struct {
	Province string `json:"province" validate:"required"`
	City     string `json:"city" validate:"required"`
}
