// SPDX-FileCopyrightText: (c) NOI Techpark <digital@noi.bz.it>

// SPDX-License-Identifier: AGPL-3.0-or-later

package main

import (
	"helloworld/lib"
	"log/slog"
	"os"
)

const STATION_TYPE string = "GolangTest"
const STATION_CODE string = "golang-test-id"

const DATA_TYPE string = "data-type-golang-test"
const ORIGIN string = "GolangUniverse"
const PERIOD uint32 = 600

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	// test data types
	var dataTypes []lib.DataType
	dataType := lib.CreateDataType(DATA_TYPE, "kg", "Such description", "Instantaneous", PERIOD)
	dataTypes = append(dataTypes, dataType)
	lib.SyncDataTypes(STATION_TYPE, dataTypes)

	// // test stations
	var stations []lib.Station
	station := lib.CreateStation(STATION_CODE, "golang-test-name", STATION_TYPE, 42.1, 11, ORIGIN)
	stations = append(stations, station)
	lib.SyncStations(STATION_TYPE, stations)

	// // test records
	var records []lib.Record
	for i := 1; i < 12; i++ {
		record := lib.CreateRecord(12.7*float64(i), PERIOD)
		records = append(records, record)
	}

	var dataMap lib.DataMap

	lib.AddRecords(STATION_CODE, DATA_TYPE, records, &dataMap)

	lib.PushData(STATION_TYPE, dataMap)
}
