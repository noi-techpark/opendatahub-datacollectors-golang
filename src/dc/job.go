package dc

import "helloworld/lib"

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

func Job() {
	forecast := Mapping(GetData())
	var stations []lib.Station
	var modelStations []lib.Station

	modelStation := lib.CreateStation(forecast.Info.Model, forecast.Info.Model, STATION_TYPE_MODEL, BZ_LAT, BZ_LON, ORIGIN)
	modelStations = append(modelStations, modelStation)

	for _, mun := range forecast.Municipalities {
		// TODO add municipality mapping with data from Open Data Hub
		station := lib.CreateStation(mun.Code, mun.NameDe+"_"+mun.NameIt, STATION_TYPE_DATA, BZ_LAT, BZ_LON, ORIGIN)
		station.ParentStation = modelStation.Id

		stations = append(stations, station)
	}

	lib.SyncStations(STATION_TYPE_MODEL, modelStations)
	lib.SyncStations(STATION_TYPE_DATA, stations)
}

func DataTypesModel() {
	var dataTypes []lib.DataType

	dataTypes = append(dataTypes, lib.CreateDataType("forecast-air-temperature-max", "Celcius", "Forecast of max air temperature during a day", "Forecast"))
	dataTypes = append(dataTypes, lib.CreateDataType("forecast-air-temperature-max", "Celcius", "Forecast of max air temperature during a day", "Forecast"))
	dataTypes = append(dataTypes, lib.CreateDataType("forecast-air-temperature-max", "Celcius", "Forecast of max air temperature during a day", "Forecast"))
	dataTypes = append(dataTypes, lib.CreateDataType("forecast-air-temperature-max", "Celcius", "Forecast of max air temperature during a day", "Forecast"))

	lib.SyncDataTypes(STATION_TYPE_MODEL, dataTypes)
}
