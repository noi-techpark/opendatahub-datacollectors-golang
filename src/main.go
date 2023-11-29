// SPDX-FileCopyrightText: (c) NOI Techpark <digital@noi.bz.it>

// SPDX-License-Identifier: AGPL-3.0-or-later

package main

import (
	"helloworld/lib"
)

func main() {

	// test data types
	var dataTypes []lib.DataType
	dataType := lib.CreateDataType("golang test", "test", "testing golang", "instantaneous", 600)
	dataTypes = append(dataTypes, dataType)
	lib.SyncDataTypes(dataTypes)

	// test stations
	var stations []lib.Station
	station := lib.CreateStation("golang test station")
	stations = append(stations, station)
	lib.SyncStations(stations)

	// test records
	var records []lib.Record
	for i := 1; i < 12; i++ {
		record := lib.CreateRecord(12.7 * float64(i))
		records = append(records, record)
	}

	lib.PushData(records)
}
