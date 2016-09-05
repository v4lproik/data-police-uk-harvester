package main

import (
	"log"
	"os"

	"github.com/v4lproik/data-police-uk-harvester/connector"
	"fmt"
	"strconv"
	"strings"
	"net/http"
)

func init() {
	log.SetOutput(os.Stdout)
}

func CliDisplay(argsWithoutProg []string){

	if len(argsWithoutProg) < 4 {
		fmt.Fprintf(os.Stderr, "usage ./data-police-uk-harvester cli <latitude> <longitude> <date>\n")
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
		fmt.Println("Date parameter: should have the following syntax <year>-<month>")
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


	log.Println("response ", data)
}

func WebServerDisplay(){
	http.Handle("/",  http.FileServer(http.Dir("./public")))
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func UsageDisplay(){
	fmt.Fprintf(os.Stderr, "usage ./data-police-uk-harvester <web-server|cli>\n")
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
