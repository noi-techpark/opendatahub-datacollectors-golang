package dc

type LocationOdh struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

type MunicipalityOdh struct {
	Id   string  `json:"id"`
	Name string  `json:"name"`
	Lat  float64 `json:"latitude"`
	Lon  float64 `json:"longitude"`
}
