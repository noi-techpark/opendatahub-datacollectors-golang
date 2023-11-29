// SPDX-FileCopyrightText: (c) NOI Techpark <digital@noi.bz.it>

// SPDX-License-Identifier: AGPL-3.0-or-later

package main

import (
	"helloworld/lib"
)

const STATION_TYPE string = "GolangTest"
const ORIGIN string = "GolangUniverse"
const PERIOD int64 = 600

func main() {

	// first create provenance
	lib.PushProvenance()

	// test data types
	var dataTypes []lib.DataType
	dataType := lib.CreateDataType("golang-test", "kg", "Such description", "Instantaneous", PERIOD)
	dataTypes = append(dataTypes, dataType)
	lib.SyncDataTypes(STATION_TYPE, dataTypes)

	// // test stations
	var stations []lib.Station
	station := lib.CreateStation("golang-test-id", "golang-test-name", STATION_TYPE, 42.1, 11, ORIGIN)
	stations = append(stations, station)
	lib.SyncStations(STATION_TYPE, stations)

	// // test records
	var records []lib.Record
	for i := 1; i < 12; i++ {
		record := lib.CreateRecord(12.7*float64(i), PERIOD)
		records = append(records, record)
	}

	dataMap := lib.CreateDataMap(records, "golang-data-map", nil)

	lib.PushData(STATION_TYPE, dataMap)
}
