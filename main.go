package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/tidwall/gjson"
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
	//vypis všeho v json
	//fmt.Println(string(body))
	if err != nil {
		fmt.Println("No Response")
	}
	return string(body)
}

func getDelay(data string, trnNum string) (actDel int) {

	var number = gjson.Get(data, "#.number") // Načte čísla vlaků z JSON
	var delay = gjson.Get(data, "#.delay")   // Načte zpoždění z JSON
	// Projde celý string čísel vlaků a hledá string odpovídající trnNum
	// TODO: pokud nepouziju proměnnou item mam problem protoze ji deklaruji ale nepouzivam!
	for i, item := range number.Array() {
		// Prasárna ale nevím jak jinak -> ukládám string na pozici I do proměnné actNum
		actNum := number.Array()[i].Str
		if actNum == trnNum {
			//fmt.Println(actNum, item)
			// Prasárna 2 ukládám delay na pozici i do proměnné actDel
			actDel := delay.Array()[i].Int()
			fmt.Println(actDel, item)
			return int(actDel)

		}

	}
	return
}

// func getTrNum(w http.ResponseWriter, r *http.Request) {
// 	// Určuji typ imputu na json
// 	w.Header().Set("Content-Type", "aplication/json")
// 	var test []spoj
// 	params := mux.Vars(r)
// 	// Prohledání knížek cyklem FOR
// 	for _, item := range test {
// 		if item.Number == params["Number"] {
// 			json.NewEncoder(w).Encode(test)
// 			return
// 		}

// 	}
// 	json.NewEncoder(w).Encode(&spoj{})
// }

func main() {
	//router := mux.NewRouter()
	url := "https://brn-ybus-pubapi.sa.cz/restapi/routes/372825000/departures"
	//Vypis delay ->
	// POZOR v zdrojovem JSON se nachazi pole protože
	// v úvodu file je [] pokud by tam bylo {} jedná se o objekt
	// Proto musí být var result spoj s []!!!!!!
	var dataLoaded = loadData(url)
	var trnNum string = "1002"
	var delay int

	delay = getDelay(dataLoaded, trnNum)
	println("Vraceno", delay)

	// tmp := gjson.Get(dataLoaded, "#.number")
	// tmp1 := gjson.Get(dataLoaded, "#.delay")
	// tmp.ForEach(func(key, value gjson.Result) bool {
	// 	tmp1 := value.Str
	// 	if tmp1 == trnNum {

	// 	}

	// 	return true
	// })
	//fmt.Println(tmp)
	//fmt.Println(tmp1)
	//VYPIŠ DELAY
	// for _, del := range result {
	// 	fmt.Println(del.Delay)
	// }

	// //VYPIŠ TRAIN NUM
	// for _, trN := range result {
	// 	fmt.Println(trN.Number)
	// }

	//router.HandleFunc("/api/train/{Number}", getTrNum).Methods("GET")
	//log.Fatal(http.ListenAndServe(":8000", router))
}
