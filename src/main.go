// SPDX-FileCopyrightText: (c) NOI Techpark <digital@noi.bz.it>

// SPDX-License-Identifier: AGPL-3.0-or-later

package main

import (
	"log"
)

type DataType struct {
	Name        string
	Unit        string
	Description string
	Rtype       string
	Period      int64
	Metadata    map[string]string
}

type Station struct {
	Id            string
	Name          string
	StationType   string
	Latitude      float64
	Longitude     float64
	Origin        string
	ParentStation string
	Metadata      map[string]string
}

type Record struct {
	Value     interface{}
	Period    int64
	CreatedOn uint32
	Timestamp uint32
}

func main() {

	// test data types
	var dataTypes []DataType
	dataType := createDataType("golang test", "test", "testing golang", "instantaneous", 600)
	dataTypes = append(dataTypes, dataType)
	syncDataTypes(dataTypes)

	// test stations
	var stations []Station
	station := createStation("golang test station")
	stations = append(stations, station)
	syncStations(stations)

	// test records
	var records []Record
	record := createRecord(12.7)
	records = append(records, record)
	pushData(records)
}

func syncDataTypes(dataTypes []DataType) {
	log.Println("Syncing data types...")
	log.Println(dataTypes)
	log.Println("Syncing data types done.")

}

func syncStations(stations []Station) {
	log.Println("Syncing stations...")
	log.Println(stations)
	log.Println("Syncing stations done.")
}

func pushData(records []Record) {
	log.Println("Syncing records...")
	log.Println(records)
	log.Println("Syncing records done.")
}

func createDataType(name string, unit string, description string, rtype string, period int64) DataType {
	var dataType DataType
	dataType.Name = name
	dataType.Unit = unit
	dataType.Description = description
	dataType.Rtype = rtype
	dataType.Period = period
	return dataType
}

func createStation(name string) Station {
	var station Station
	station.Name = name
	return station
}

func createRecord(value interface{}) Record {
	var record Record
	record.Value = value
	return record
}
