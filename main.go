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

// Načte data z URL -> příjmá URL string a vrací string body (data v body proměnné)
func loadData(url string) (content string) {
	resp, err := http.Get(url)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("No Response")
	}
	return string(body)
}

func main() {
	url := "https://brn-ybus-pubapi.sa.cz/restapi/routes/372825000/departures"

	//vypis všeho v json
	//fmt.Println(string(body))

	//Vypis delay ->
	// POZOR v zdrojovem JSON se nachazi pole protože
	// v úvodu file je [] pokud by tam bylo {} jedná se o objekt
	// Proto musí být var result spoj s []!!!!!!
	var result []spoj
	if err := json.Unmarshal([]byte(loadData(url)), &result); err != nil { // Parse []byte to the go struct pointer
		fmt.Println("Can not unmarshal JSON")
	}
	for _, rec := range result {
		fmt.Println(rec.Delay)
	}
}
