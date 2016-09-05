package connector

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"fmt"
	"strconv"
)

const DOMAIN  = "https://data.police.uk/api"
const API = "/stops-street"

type Stop struct {
	AgeRange string `json:"age_range"`
	SelfDefinedEthnicity string `json:"self_defined_ethnicity"`
	DateTime string `json:"datetime"`
	RemovalOfMoreThanOuterClothing bool `json:"removal_of_more_than_outer_clothing"`
	Operation string `json:"operation"`
	OfficerDefinedEthnicity string `json:"officer_defined_ethnicity"`
	ObjectOfSearch string `json:"object_of_search"`
	InvolvedPerson bool `json:"involved_person"`
	Gender string `json:"gender"`
	Legislation string `json:"legislation"`
	Outcome json.RawMessage `json:"outcome"`
	Type string `json:"type"`
	OperationName string `json:"operation_name"`
	Location Location `json:"location"`
}

type Location struct {
	Latitude string `json:"latitude"`
	Longitude string `json:"longitude"`
	Street Street `json:"street"`
}

type Street struct {
	Id int64 `json:"id"`
	Name string `json:"name"`
}


func GetData(lat float64, lng float64, year int64, month int64) ([]Stop, error) {

	url := DOMAIN + API + "?lat=" + strconv.FormatFloat(lat, 'f', -1, 64) + "&lng=" + strconv.FormatFloat(lng, 'f', -1, 64) + "&date=" + strconv.FormatInt(year, 10) + "-" + strconv.FormatInt(month, 10)

	res, err := http.Get(url)
	if err != nil {
		panic(err.Error())
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err.Error())
	}

	s, err := getStops([]byte(body))

	return s, err
}

func getStops(body []byte) ([]Stop, error) {
	items := make([]Stop, 10)

	err := json.Unmarshal(body, &items)
	if(err != nil){
		fmt.Println("whoops:", err)
	}
	return items, err
}
