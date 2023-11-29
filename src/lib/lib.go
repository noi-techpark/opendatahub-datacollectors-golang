package lib

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"helloworld/auth"
)

type Provenance struct {
	Uuid                 string `json:"uuid"`
	Lineage              string `json:"lineage"`
	DataCollector        string `json:"dataCollector"`
	DataCollectorVersion string `json:"dataCollectorVersion"`
}

type DataType struct {
	Name        string            `json:"name"`
	Unit        string            `json:"unit"`
	Description string            `json:"description"`
	Rtype       string            `json:"rType"`
	Period      int64             `json:"period"`
	Metadata    map[string]string `json:"metadata"`
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
	CreatedOn int64
	Timestamp int64
}

const SYNC_DATA_TYPES string = "/syncDataTypes"
const SYNC_STATIONS string = "/syncStations"
const PUSH_RECORDS string = "/pushRecords"
const GET_DATE_OF_LAST_RECORD string = "/getDateOfLastRecord"
const STATIONS string = "/stations"
const PROVENANCE string = "/provenance"

var baseUri string = os.Getenv("BASE_URI")

var prv string = os.Getenv("PROVENANCE_VERSION")
var prn string = os.Getenv("PROVENANCE_NAME")

func PushProvenance() {
	log.Println("Pushing provenance...")
	log.Println("prv: " + prv + " prn: " + prn)

	var provenance = Provenance{
		DataCollector:        prn,
		DataCollectorVersion: prv,
		Uuid:                 "suchUuuid12345678ByGolangDataCollecotr",
		Lineage:              "go-lang-lineage",
	}

	// prepare data
	data, err := json.Marshal(provenance)
	if err != nil {
		log.Fatal(err)
	}

	// URL
	fullUrl := baseUri + PROVENANCE + "?&prn=" + prn + "&prv=" + prv
	log.Println("fullUri = " + fullUrl)
	// fullUrl = url.QueryEscape(fullUrl)

	// http client
	client := http.Client{}
	req, err := http.NewRequest("POST", fullUrl, bytes.NewBuffer(data))
	if err != nil {
		log.Fatal(err)
	}

	req.Header = http.Header{
		"Content-Type":  {"application/json"},
		"Authorization": {"Bearer " + auth.GetToken()},
	}

	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(res)

	log.Println("Pushing provenance done.")
}

func SyncDataTypes(stationType string, dataTypes []DataType) {
	log.Println("Syncing data types...")
	log.Println(dataTypes)

	// prepare data
	data, err := json.Marshal(dataTypes)
	if err != nil {
		log.Fatal(err)
	}

	// URL
	fullUrl := baseUri + SYNC_DATA_TYPES + "?stationType=" + stationType + "&prn=" + prn + "&prv=" + prv
	log.Println("fullUri = " + fullUrl)
	// fullUrl = url.QueryEscape(fullUrl)

	// http client
	client := http.Client{}
	req, err := http.NewRequest("POST", fullUrl, bytes.NewBuffer(data))
	if err != nil {
		log.Fatal(err)
	}

	req.Header = http.Header{
		"Content-Type":  {"application/json"},
		"Authorization": {"Bearer " + auth.GetToken()},
	}

	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(res)

	log.Println("Syncing data types done.")
}

func SyncStations(stationType string, stations []Station) {
	log.Println("Syncing stations...")
	log.Println(stations)
	// prepare data
	data, err := json.Marshal(stations)
	if err != nil {
		log.Fatal(err)
	}

	// URL
	fullUrl := baseUri + SYNC_STATIONS + "/" + stationType + "?prn=" + prn + "&prv=" + prv
	log.Println("fullUri = " + fullUrl)
	// fullUrl = url.QueryEscape(fullUrl)

	// http client
	client := http.Client{}
	req, err := http.NewRequest("POST", fullUrl, bytes.NewBuffer(data))
	if err != nil {
		log.Fatal(err)
	}

	req.Header = http.Header{
		"Content-Type":  {"application/json"},
		"Authorization": {"Bearer " + auth.GetToken()},
	}

	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(res)
	log.Println("Syncing stations done.")
}

func PushData(records []Record) {
	log.Println("Syncing records...")
	log.Println(records)
	log.Println("Syncing records done.")
}

func CreateDataType(name string, unit string, description string, rtype string, period int64) DataType {
	// TODO add some checks
	return DataType{
		Name:        name,
		Unit:        unit,
		Description: description,
		Rtype:       rtype,
		Period:      period,
	}
}

func CreateStation(id string, name string, stationType string, lat float64, lon float64, origin string) Station {
	// TODO add some checks
	var station = Station{
		Name:        name,
		StationType: stationType,
		Latitude:    lat,
		Longitude:   lon,
		Origin:      origin,
		Id:          id,
		// Metadata:    metaData,
		// ParentStation:   parentStation,
	}
	return station
}

func CreateRecord(value interface{}, period int64) Record {
	// TODO add some checks
	var record = Record{
		Value:     value,
		Timestamp: time.Now().Unix(),
		CreatedOn: time.Now().Unix(),
		Period:    period,
	}
	return record
}
