package data

type Product struct {
	ID         *string `json:"uuid"`
	Barang     *string `json:"komoditas"`
	Provinsi   *string `json:"area_provinsi"`
	Kota       *string `json:"area_kota"`
	Size       *string `json:"size"`
	Price      *string `json:"price"`
	Tanggal    *string `json:"tgl_parsed"`
	Timestampz *string `json:"timestamp"`
	USDPrice   string  `json:"price_usd"`
}
