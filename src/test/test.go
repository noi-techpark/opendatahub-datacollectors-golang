package test

import (
	"helloworld/lib"
	"log/slog"
)

const stationType string = "GolangTest"
const stationCode string = "golang-test-id"
const dataTypeId string = "data-type-golang-test"
const origin string = "GolangUniverse"
const period uint32 = 600

func TestJob() {
	slog.Info("Cron job started...")
	// test data types
	var dataTypes []lib.DataType
	dataType := lib.CreateDataType(dataTypeId, "kg", "Such description", "Instantaneous")
	dataTypes = append(dataTypes, dataType)
	lib.SyncDataTypes(stationType, dataTypes)

	// // test stations
	var stations []lib.Station
	station := lib.CreateStation(stationCode, "golang-test-name", stationType, 42.1, 11, origin)
	stations = append(stations, station)
	lib.SyncStations(stationType, stations)

	// // test records
	var records []lib.Record
	for i := 1; i < 12; i++ {
		record := lib.CreateRecord(12.7*float64(i), period)
		records = append(records, record)
	}

	var dataMap lib.DataMap

	lib.AddRecords(stationCode, dataTypeId, records, &dataMap)

	lib.PushData(stationType, dataMap)

	slog.Info("Cron job done.")
}
