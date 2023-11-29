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

type DataMap struct {
	Name       string
	Data       []Record
	Branch     map[string]DataMap
	Provenance string
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

var provenancePushed bool = false

func SyncDataTypes(stationType string, dataTypes []DataType) {
	pushProvenance()

	log.Println("Syncing data types...")
	log.Println(dataTypes)

	url := baseUri + SYNC_DATA_TYPES + "?stationType=" + stationType + "&prn=" + prn + "&prv=" + prv

	postToWriter(dataTypes, url)

	log.Println("Syncing data types done.")
}

func SyncStations(stationType string, stations []Station) {
	pushProvenance()

	log.Println("Syncing stations...")
	log.Println(stations)

	url := baseUri + SYNC_STATIONS + "/" + stationType + "?prn=" + prn + "&prv=" + prv

	postToWriter(stations, url)

	log.Println("Syncing stations done.")
}

func PushData(stationType string, dataMap DataMap) {
	pushProvenance()

	log.Println("Pushing records...")
	log.Println(dataMap)

	url := baseUri + PUSH_RECORDS + "/" + stationType + "?prn=" + prn + "&prv=" + prv

	postToWriter(dataMap, url)

	log.Println("Pushing records done.")
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

func CreateDataMap(records []Record, name string, branch map[string]DataMap) DataMap {
	// TODO add some checks
	var dataMap = DataMap{
		Name:       name,
		Data:       records,
		Provenance: PROVENANCE,
		// Branch:     branch,
	}

	return dataMap
}

func postToWriter(data interface{}, fullUrl string) {
	json, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}

	client := http.Client{}
	req, err := http.NewRequest("POST", fullUrl, bytes.NewBuffer(json))
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

	log.Println(res.StatusCode)
}

func pushProvenance() {
	if provenancePushed {
		return
	}

	log.Println("Pushing provenance...")
	log.Println("prv: " + prv + " prn: " + prn)

	var provenance = Provenance{
		DataCollector:        prn,
		DataCollectorVersion: prv,
		Uuid:                 "suchUuuid12345678ByGolangDataCollecotr",
		Lineage:              "go-lang-lineage",
	}

	url := baseUri + PROVENANCE + "?&prn=" + prn + "&prv=" + prv

	postToWriter(provenance, url)

	log.Println("Pushing provenance done.")

	provenancePushed = true
}
