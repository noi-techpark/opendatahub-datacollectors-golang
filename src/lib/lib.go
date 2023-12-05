package lib

import (
	"bufio"
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"helloworld/auth"
)

type Provenance struct {
	Lineage              string `json:"lineage"`
	DataCollector        string `json:"dataCollector"`
	DataCollectorVersion string `json:"dataCollectorVersion"`
}

type DataType struct {
	Name        string            `json:"name"`
	Unit        string            `json:"unit"`
	Description string            `json:"description"`
	Rtype       string            `json:"rType"`
	Period      uint32            `json:"period"`
	Metadata    map[string]string `json:"metadata"`
}

type Station struct {
	Id            string            `json:"id"`
	Name          string            `json:"name"`
	StationType   string            `json:"stationType"`
	Latitude      float64           `json:"latitude"`
	Longitude     float64           `json:"longitude"`
	Origin        string            `json:"origin"`
	ParentStation string            `json:"parentStation"`
	Metadata      map[string]string `json:"metadata"`
}

type DataMap struct {
	Name       string             `json:"name"`
	Data       []Record           `json:"data"`
	Branch     map[string]DataMap `json:"branch"`
	Provenance string             `json:"provenance"`
}

type Record struct {
	Value     interface{} `json:"value"`
	Period    uint32      `json:"period"`
	CreatedOn int64       `json:"createdOn"`
	Timestamp int64       `json:"timestamp"`
	Type      string      `json:"_t"`
}

const SYNC_DATA_TYPES string = "/syncDataTypes"
const SYNC_STATIONS string = "/syncStations"
const PUSH_RECORDS string = "/pushRecords"
const GET_DATE_OF_LAST_RECORD string = "/getDateOfLastRecord"
const STATIONS string = "/stations"
const PROVENANCE string = "/provenance"

var provenanceUuid string

var baseUri string = os.Getenv("BASE_URI")

var prv string = os.Getenv("PROVENANCE_VERSION")
var prn string = os.Getenv("PROVENANCE_NAME")

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

func CreateDataType(name string, unit string, description string, rtype string, period uint32) DataType {
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

func CreateRecord(value interface{}, period uint32) Record {
	// TODO add some checks
	var record = Record{
		Value:     value,
		Timestamp: time.Now().Unix(),
		CreatedOn: time.Now().Unix(),
		Period:    period,
		Type:      "it.bz.idm.bdp.dto.SimpleRecordDto",
	}
	return record
}

func createDataMap() DataMap {
	var dataMap = DataMap{
		Name:       "(default)",
		Provenance: provenanceUuid,
		Branch:     make(map[string]DataMap),
	}
	return dataMap
}

func AddRecords(stationCode string, datatType string, records []Record, dataMap *DataMap) {

	if dataMap.Name == "" {
		*dataMap = createDataMap()
	}

	dataMap.Branch[stationCode] = DataMap{
		Name:   "(default)",
		Branch: make(map[string]DataMap),
	}

	dataMap.Branch[stationCode].Branch[datatType] = DataMap{
		Name: "(default)",
		Data: records,
	}
}

func postToWriter(data interface{}, fullUrl string) (string, error) {
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

	scanner := bufio.NewScanner(res.Body)
	for i := 0; scanner.Scan() && i < 5; i++ {
		return scanner.Text(), nil
	}

	err = scanner.Err()
	if err != nil {
		log.Fatal(err)
	}
	return "", err
}

func pushProvenance() {
	if len(provenanceUuid) > 0 {
		return
	}

	log.Println("Pushing provenance...")
	log.Println("prv: " + prv + " prn: " + prn)

	var provenance = Provenance{
		DataCollector:        prn,
		DataCollectorVersion: prv,
		Lineage:              "go-lang-lineage",
	}

	url := baseUri + PROVENANCE + "?&prn=" + prn + "&prv=" + prv

	res, err := postToWriter(provenance, url)

	if err != nil {
		log.Fatal(err)
	}

	provenanceUuid = res

	log.Println("Pushing provenance done. UUID: ", provenanceUuid)
}
