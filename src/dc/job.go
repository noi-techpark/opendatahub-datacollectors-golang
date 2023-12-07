package dc

import (
	"fmt"
	"helloworld/lib"
	"log/slog"
	"time"
)

const stationTypeModel string = "WeatherForecastService"
const stationTypeData string = "WeatherForecast"
const origin string = "province-bolzano"

const bzLat float64 = 46.49067
const bzLon float64 = 11.33982

const airTemperatureMax string = "forecast-air-temperature-max"
const airTemperatureMin string = "forecast-air-temperature-min"
const airTemperature string = "forecast-air-temperature"
const windDirection string = "forecast-wind-direction"
const windSpeed string = "forecast-wind-speed"
const sunshineDuration string = "forecast-sunshine-duration"
const precipitationProbability string = "forecast-precipitation-probability"
const qualitativeForecast string = "qualitative-forecast"
const precipitationSum string = "forecast-precipitation-sum"
const precipitationMax string = "forecast-precipitation-max"
const precipitationMin string = "forecast-precipitation-min"

const period3 uint64 = 10800
const period12 uint64 = 43200
const period24 uint64 = 86400

func Job() {
	forecast := Mapping(GetData())
	var modelStations []lib.Station
	var modelDataMap lib.DataMap
	var modelRecords []lib.Record

	var stations []lib.Station
	var dataMap lib.DataMap

	modelTs := dateToTs(forecast.Info.CurrentModelRun)
	slog.Debug("model timestamp " + fmt.Sprint(modelTs))

	////////////////////
	// Model
	////////////////////
	modelStation := lib.CreateStation(forecast.Info.Model, forecast.Info.Model, stationTypeModel, bzLat, bzLon, origin)
	modelStations = append(modelStations, modelStation)

	modelRecords = append(modelRecords, lib.CreateRecord(modelTs, forecast.Info.AbsTempMax, period12))
	modelRecords = append(modelRecords, lib.CreateRecord(modelTs, forecast.Info.AbsTempMin, period12))
	modelRecords = append(modelRecords, lib.CreateRecord(modelTs, forecast.Info.AbsPrecMax, period12))
	modelRecords = append(modelRecords, lib.CreateRecord(modelTs, forecast.Info.AbsPrecMin, period12))

	lib.AddRecords(modelStation.Id, airTemperatureMin, modelRecords, &modelDataMap)

	////////////////////
	// Forecast Data
	////////////////////
	for _, mun := range forecast.Municipalities {
		// TODO add municipality mapping with data from Open Data Hub
		station := lib.CreateStation(mun.Code, mun.NameDe+"_"+mun.NameIt, stationTypeData, bzLat, bzLon, origin)
		station.ParentStation = modelStation.Id

		var records []lib.Record

		// temperature min 24 hours
		for _, value := range mun.TempMin24.Data {
			records = append(records, lib.CreateRecord(dateToTs(value.Date), value.Value, period24))
		}
		// temperature max 24 hours
		for _, value := range mun.TempMax24.Data {
			records = append(records, lib.CreateRecord(dateToTs(value.Date), value.Value, period24))
		}
		// temperature every 3 hours
		for _, value := range mun.Temp3.Data {
			records = append(records, lib.CreateRecord(dateToTs(value.Date), value.Value, period3))
		}
		// sunshine duration 24 hours
		for _, value := range mun.Ssd24.Data {
			records = append(records, lib.CreateRecord(dateToTs(value.Date), value.Value, period24))
		}
		// precipitation probability 3 hours
		for _, value := range mun.PrecProb3.Data {
			records = append(records, lib.CreateRecord(dateToTs(value.Date), value.Value, period3))
		}
		// probably precipitation 24 hours
		for _, value := range mun.PrecProb24.Data {
			records = append(records, lib.CreateRecord(dateToTs(value.Date), value.Value, period24))
		}
		// probably precipitation sum 3 hours
		for _, value := range mun.PrecSum3.Data {
			records = append(records, lib.CreateRecord(dateToTs(value.Date), value.Value, period3))
		}
		// probably precipitation sum 24 hours
		for _, value := range mun.PrecSum24.Data {
			records = append(records, lib.CreateRecord(dateToTs(value.Date), value.Value, period24))
		}
		// wind direction 3 hours
		for _, value := range mun.WindDir3.Data {
			records = append(records, lib.CreateRecord(dateToTs(value.Date), value.Value, period3))
		}
		// wind speed 3 hours
		for _, value := range mun.WindSpd3.Data {
			records = append(records, lib.CreateRecord(dateToTs(value.Date), value.Value, period3))
		}
		// weather status symbols 3 hours
		for _, value := range mun.Symbols3.Data {
			records = append(records, lib.CreateRecord(dateToTs(value.Date), MapQuantitative(value.Value.(string)), period3))
		}
		// weather status symbols 24 hours
		for _, value := range mun.Symbols24.Data {
			records = append(records, lib.CreateRecord(dateToTs(value.Date), MapQuantitative(value.Value.(string)), period24))
		}

		stations = append(stations, station)
		lib.AddRecords(station.Id, airTemperatureMin, records, &dataMap)
	}

	// sync stations
	lib.SyncStations(stationTypeModel, modelStations)
	lib.SyncStations(stationTypeData, stations)

	// push data
	lib.PushData(stationTypeData, dataMap)
	lib.PushData(stationTypeModel, modelDataMap)

}

func dateToTs(date string) int64 {
	time, err := time.Parse(time.RFC3339Nano, date)
	if err != nil {
		slog.Error("error", err)
	}
	return time.UnixMilli()
}

func DataTypesModel() {
	var dataTypes []lib.DataType

	dataTypes = append(dataTypes, lib.CreateDataType(airTemperatureMax, "Celcius", "Forecast of max air temperature", "Forecast"))
	dataTypes = append(dataTypes, lib.CreateDataType(airTemperatureMin, "Celcius", "Forecast of min air temperature", "Forecast"))
	dataTypes = append(dataTypes, lib.CreateDataType(precipitationMax, "mm", "Forecast of max precipitation", "Forecast"))
	dataTypes = append(dataTypes, lib.CreateDataType(precipitationMin, "mm", "Forecast of min precipitation", "Forecast"))

	lib.SyncDataTypes(stationTypeModel, dataTypes)
}

func DataTypes() {
	var dataTypes []lib.DataType

	dataTypes = append(dataTypes, lib.CreateDataType(airTemperatureMax, "Celcius", "Forecast of max air temperature during a day", "Forecast"))
	dataTypes = append(dataTypes, lib.CreateDataType(airTemperatureMin, "Celcius", "Forecast of min air temperature during a day", "Forecast"))
	dataTypes = append(dataTypes, lib.CreateDataType(airTemperature, "Celcius", "Forecast of air temperature at a specific timestamp", "Forecast"))
	dataTypes = append(dataTypes, lib.CreateDataType(windDirection, "\\u00b0", "Forecast of wind direction", "Forecast"))
	dataTypes = append(dataTypes, lib.CreateDataType(windSpeed, "m/s", "Forecast of wind speed", "Forecast"))
	dataTypes = append(dataTypes, lib.CreateDataType(sunshineDuration, "h", "Forecast of sun shine duration", "Forecast"))
	dataTypes = append(dataTypes, lib.CreateDataType(precipitationProbability, "%", "Forecast of precipitation probability", "Forecast"))
	dataTypes = append(dataTypes, lib.CreateDataType(qualitativeForecast, "", "Forecast of overall weather condition. Example: sunny", "Forecast"))
	dataTypes = append(dataTypes, lib.CreateDataType(precipitationSum, "mm", "Forecast of cumulated precipitation", "Forecast"))

	lib.SyncDataTypes(stationTypeModel, dataTypes)
}
