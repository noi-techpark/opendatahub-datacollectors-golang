package lib

import (
	"log"
	"os"
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

var baseUri string = os.Getenv("BASE_URI")
var authorizationUri string = os.Getenv("OAUTH_AUTH_URI")
var tokenUri string = os.Getenv("OAUTH_TOKEN_URI")
var clientId string = os.Getenv("OAUTH_CLIENT_ID")
var clientName string = os.Getenv("OAUTH_CLIENT_NAME")
var clientSecret string = os.Getenv("OAUTH_CLIENT_SECRET")
var scope string = os.Getenv("OAUTH_CLIENT_SCOPE")

func SyncDataTypes(dataTypes []DataType) {
	log.Println("Syncing data types...")
	log.Println(dataTypes)

	log.Println("Syncing data types done.")

}

func SyncStations(stations []Station) {
	log.Println("Syncing stations...")
	log.Println(stations)

	log.Println("Syncing stations done.")
}

func PushData(records []Record) {
	log.Println("Syncing records...")
	log.Println(records)
	log.Println("Syncing records done.")
}

func CreateDataType(name string, unit string, description string, rtype string, period int64) DataType {
	var dataType DataType
	dataType.Name = name
	dataType.Unit = unit
	dataType.Description = description
	dataType.Rtype = rtype
	dataType.Period = period
	return dataType
}

func CreateStation(name string) Station {
	var station Station
	station.Name = name
	return station
}

func CreateRecord(value interface{}) Record {
	var record Record
	record.Value = value
	return record
}
