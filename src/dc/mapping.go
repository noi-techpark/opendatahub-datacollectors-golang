package dc

import (
	"encoding/json"
	"fmt"
	"log/slog"
)

type Forecast struct {
	Info           Info           `json:"info"`
	Municipalities []Municipality `json:"municipalities"`
}

type Info struct {
	Model            string `json:"model"`
	CurrentModelRun  string `json:"currentModelRun"`
	NextModelRun     string `json:"nextModelRun"`
	FileName         string `json:"fileName"`
	FileCreationDate string `json:"fileCreationDate"`
	AbsTempMin       int    `json:"absTempMin"`
	AbsTempMax       int    `json:"absTempMax"`
	AbsPrecMin       int    `json:"absPrecMin"`
	AbsPrecMax       int    `json:"absPrecMax"`
}

type Municipality struct {
	Code       string   `json:"code"`
	NameDe     string   `json:"nameDe"`
	NameIt     string   `json:"nameIt"`
	NameEn     string   `json:"nameEn"`
	NameRm     string   `json:"nameRm"`
	TempMin24  ValueSet `json:"tempMin24"`
	TempMax24  ValueSet `json:"tempMax24"`
	Temp3      ValueSet `json:"temp3"`
	Ssd24      ValueSet `json:"ssd24"`
	PrecProb3  ValueSet `json:"precProb3"`
	PrecProb24 ValueSet `json:"precProb24"`
	PrecSum3   ValueSet `json:"precSum3"`
	PrecSum24  ValueSet `json:"precSum24"`
	Symbols3   ValueSet `json:"symbols3"`
	Symbols24  ValueSet `json:"symbols24"`
	WindDir3   ValueSet `json:"windDir3"`
	WindSpd3   ValueSet `json:"windSpd3"`
}

type ValueSet struct {
	NameDe string  `json:"nameDe"`
	NameIt string  `json:"nameIt"`
	NameEn string  `json:"nameEn"`
	NameRm string  `json:"nameRm"`
	Unit   string  `json:"unit"`
	Data   []Value `json:"data"`
}

type Value struct {
	Date  string      `json:"date"`
	Value interface{} `json:"value"`
}

func Mapping(data []byte) {
	var forecast Forecast

	err := json.Unmarshal(data, &forecast)
	if err != nil {
		slog.Error("error", err)
	}
	fmt.Println(forecast)
}
