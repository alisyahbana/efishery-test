package data

type Product struct {
	ID           *string `json:"uuid"`
	Barang       *string `json:"komoditas"`
	AreaProvinsi *string `json:"area_provinsi"`
	AreaKota     *string `json:"area_kota"`
	Size         *string `json:"size"`
	Price        *string `json:"price"`
	TglParsed    *string `json:"tgl_parsed"`
	Timestamp    *string `json:"timestamp"`
	PriceUsd     string  `json:"price_usd"`
}

type ProductData struct {
	Provinsi string `json:"provinsi"`
	Profit   map[string]map[string]int
	Max      float64 `json:"max_profit"`
	Min      float64 `json:"min_profit"`
	Avg      float64 `json:"average_profit"`
	Median   float64 `json:"median_profit"`
}
