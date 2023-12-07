package lib

import (
	"bufio"
	"bytes"
	"encoding/json"
	"log/slog"
	"net/http"
	"os"
	"time"
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
	Timestamp int64       `json:"timestamp"`
	Type      string      `json:"_t"`
}

const syncDataTypesPath string = "/syncDataTypes"
const syncStationsPath string = "/syncStations"
const pushRecordsPath string = "/pushRecords"
const getDateOfLastRecordPath string = "/getDateOfLastRecord"
const stationsPath string = "/stations"
const provenancePath string = "/provenance"

var provenanceUuid string

var baseUri string = os.Getenv("BASE_URI")

var prv string = os.Getenv("PROVENANCE_VERSION")
var prn string = os.Getenv("PROVENANCE_NAME")

func SyncDataTypes(stationType string, dataTypes []DataType) {
	pushProvenance()

	slog.Debug("Syncing data types...")

	url := baseUri + syncDataTypes + "?stationType=" + stationType + "&prn=" + prn + "&prv=" + prv

	postToWriter(dataTypes, url)

	slog.Debug("Syncing data types done.")
}

func SyncStations(stationType string, stations []Station) {
	pushProvenance()

	slog.Info("Syncing stations...")

	url := baseUri + syncStationsPath + "/" + stationType + "?prn=" + prn + "&prv=" + prv

	postToWriter(stations, url)

	slog.Info("Syncing stations done.")
}

func PushData(stationType string, dataMap DataMap) {
	pushProvenance()

	slog.Info("Pushing records...")

	url := baseUri + pushRecordsPath + "/" + stationType + "?prn=" + prn + "&prv=" + prv

	postToWriter(dataMap, url)

	slog.Info("Pushing records done.")
}

func CreateDataType(name string, unit string, description string, rtype string) DataType {
	// TODO add some checks
	return DataType{
		Name:        name,
		Unit:        unit,
		Description: description,
		Rtype:       rtype,
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
		slog.Error("error", err)
	}

	client := http.Client{}
	req, err := http.NewRequest("POST", fullUrl, bytes.NewBuffer(json))
	if err != nil {
		slog.Error("error", err)
	}

	req.Header = http.Header{
		"Content-Type":  {"application/json"},
		"Authorization": {"Bearer " + GetToken()},
	}

	res, err := client.Do(req)
	if err != nil {
		slog.Error("error", err)
	}

	slog.Info(res.Status)

	scanner := bufio.NewScanner(res.Body)
	for i := 0; scanner.Scan() && i < 5; i++ {
		return scanner.Text(), nil
	}

	err = scanner.Err()
	if err != nil {
		slog.Error("error", err)
	}
	return "", err
}

func pushProvenance() {
	if len(provenanceUuid) > 0 {
		return
	}

	slog.Info("Pushing provenance...")
	slog.Info("prv: " + prv + " prn: " + prn)

	var provenance = Provenance{
		DataCollector:        prn,
		DataCollectorVersion: prv,
		Lineage:              "go-lang-lineage",
	}

	url := baseUri + provenancePath + "?&prn=" + prn + "&prv=" + prv

	res, err := postToWriter(provenance, url)

	if err != nil {
		slog.Error("error", err)
	}

	provenanceUuid = res

	slog.Info("Pushing provenance done.")
}
