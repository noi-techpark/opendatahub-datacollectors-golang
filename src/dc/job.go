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
const precipitationMax string = "forecast-precipitation-max"
const precipitationMin string = "forecast-precipitation-min"

func Job() {
	forecast := Mapping(GetData())
	var stations []lib.Station
	var modelStations []lib.Station

	modelStation := lib.CreateStation(forecast.Info.Model, forecast.Info.Model, stationTypeModel, bzLat, bzLon, origin)
	modelStations = append(modelStations, modelStation)

	for _, mun := range forecast.Municipalities {
		// TODO add municipality mapping with data from Open Data Hub
		station := lib.CreateStation(mun.Code, mun.NameDe+"_"+mun.NameIt, stationTypeData, bzLat, bzLon, origin)
		station.ParentStation = modelStation.Id

		stations = append(stations, station)
	}

	lib.SyncStations(stationTypeModel, modelStations)
	lib.SyncStations(stationTypeData, stations)
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
