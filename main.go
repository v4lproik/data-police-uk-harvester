package main

import (
	"log"
	"os"
	"github.com/v4lproik/data-police-uk-harvester/connector"
	"fmt"
	"strconv"
	"strings"
	"net/http"
	"time"
	"github.com/juju/loggo"
)

var logger = loggo.GetLogger("main")

func init()  {
	// Start displaying from info
	loggo.ConfigureLoggers("info")
}

func CliDisplay(argsWithoutProg []string){

	if len(argsWithoutProg) < 4 {
		logger.Errorf("usage ./data-police-uk-harvester cli <latitude> <longitude> <date>\n")
		os.Exit(1)
	}

	lat, err := strconv.ParseFloat(argsWithoutProg[1], 64)
	if err != nil {
		panic(err.Error())
	}

	lgt, err := strconv.ParseFloat(argsWithoutProg[2], 64)
	if err != nil {
		panic(err.Error())
	}

	splitDate := strings.Split(argsWithoutProg[3], "-")
	if len(splitDate) != 2 {
		logger.Errorf("Date parameter: should have the following syntax <year>-<month>")
		os.Exit(1)
	}

	year, err := strconv.ParseInt(splitDate[0], 10, 64)
	if err != nil {
		panic(err.Error())
	}

	month, err := strconv.ParseInt(splitDate[1], 10, 64)
	if err != nil {
		panic(err.Error())
	}

	data, err := connector.GetData(lat, lgt, year, month)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	logger.Infof("response ", data)
}

func getData(w http.ResponseWriter, r *http.Request) {
	longitude := r.URL.Query().Get("longitude")
	latitude := r.URL.Query().Get("latitude")

	lat, err := strconv.ParseFloat(latitude, 64)
	if err != nil {
		panic(err.Error())
	}

	lgt, err := strconv.ParseFloat(longitude, 64)
	if err != nil {
		panic(err.Error())
	}

	now := time.Now()
	data, err := connector.GetData(lat, lgt, int64(now.Year()-1), int64(now.Month()))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	logger.Infof("response ", data)
	fmt.Fprint(w, data)
}

func WebServerDisplay(){
	//static html
	http.Handle("/",  http.FileServer(http.Dir("./public")))
	//ajax call
	http.HandleFunc("/getData", getData)

	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func UsageDisplay(){
	logger.Errorf("usage ./data-police-uk-harvester <web-server|cli>\n")
	os.Exit(1)
}

func main() {

	argsWithoutProg := os.Args[1:]

	if len(argsWithoutProg) < 1 {
		UsageDisplay()
	}

	if argsWithoutProg[0] == "web-server" {
		WebServerDisplay()
	}else{
		if argsWithoutProg[0] == "cli" {
			CliDisplay(argsWithoutProg)
		}else {
			UsageDisplay()
		}
	}

}
