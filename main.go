package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type spoj struct {
	BusConnectionID    int64  `json:"busConnectionId"`
	Label              string `json:"label"`
	Number             string `json:"number"`
	Delay              int    `json:"delay"`
	VehicleCategory    string `json:"vehicleCategory"`
	FreeSeatsCount     int    `json:"freeSeatsCount"`
	ConnectionStations []struct {
		StationID        int64       `json:"stationId"`
		Arrival          interface{} `json:"arrival"`
		Departure        time.Time   `json:"departure"`
		Platform         interface{} `json:"platform"`
		DepartingStation bool        `json:"departingStation"`
	} `json:"connectionStations"`
}

func main() {
	url := "https://brn-ybus-pubapi.sa.cz/restapi/routes/372825000/departures"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("No Response")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	//vypis všeho v json
	//fmt.Println(string(body))

	//Vypis delay
	var result spoj
	if err := json.Unmarshal(body, &result); err != nil { // Parse []byte to the go struct pointer
		fmt.Println("Can not unmarshal JSON")
	}
	for _, rec := range result.Label {
		fmt.Println(rec)
	}
}
