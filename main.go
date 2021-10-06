package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type delay struct {
	Delay string `json:"delay"`
}

const (
	depUrl = "https://brn-ybus-pubapi.sa.cz/restapi/routes/372825000/departures"
)

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

// MOJE IMPLEMENTACE S využitím JSON PARSERU GJSON
// func getDelay(data string, trnNum string) (actDel int) {

// 	var number = gjson.Get(data, "#.number") // Načte čísla vlaků z JSON
// 	var delay = gjson.Get(data, "#.delay")   // Načte zpoždění z JSON
// 	// Projde celý string čísel vlaků a hledá string odpovídající trnNum
// 	// TODO: pokud nepouziju proměnnou item mam problem protoze ji deklaruji ale nepouzivam!
// 	for i, item := range number.Array() {
// 		// Prasárna ale nevím jak jinak -> ukládám string na pozici I do proměnné actNum
// 		actNum := number.Array()[i].Str
// 		if actNum == trnNum {
// 			//fmt.Println(actNum, item)
// 			// Prasárna 2 ukládám delay na pozici i do proměnné actDel
// 			actDel := delay.Array()[i].Int()
// 			fmt.Println(actDel, item)
// 			return int(actDel)

// 		}

// 	}
// 	return
// }

func getDelayAlt(data string, trnNum string) (actDel int) {
	// Není nutné tahat z JSON všechna data a někam je ukládat
	//Takto si mohu vybrat jen ty které mě zajímají
	var res []struct {
		Number string `json:"number"`
		Delay  int    `json:"delay"`
	}
	err := json.Unmarshal([]byte(data), &res)
	if err != nil {
		panic(err)
	}
	for _, item := range res {
		if item.Number == trnNum {
			return item.Delay
		}
	}
	return
}

func getTrNum(w http.ResponseWriter, r *http.Request) {
	// Určuji typ outputu na JSON
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var dataLoaded = loadData(depUrl)
	// ukládám Vrácený INT (zpoždění) do var resource jako parametry
	//posílám do fce načtený JSON a paramer number z url kterým volám API
	resource := getDelayAlt(dataLoaded, params["number"])
	var d delay
	// Přetypovávám INT na string a vracím JSON
	d.Delay = strconv.Itoa(resource)
	json.NewEncoder(w).Encode(d)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/train/{number}", getTrNum).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
}
