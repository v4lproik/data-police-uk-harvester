package main

import (
	"log"
	"os"

	"github.com/v4lproik/data-police-uk-harvester/connector"
	"fmt"
	"strconv"
	"strings"
)

func init() {
	log.SetOutput(os.Stdout)
}

func main() {

	argsWithoutProg := os.Args[1:]

	if len(argsWithoutProg) < 3 {
		fmt.Fprintf(os.Stderr, "usage ./data-police-uk-harvester <latitude> <longitude> <date>\n")
		os.Exit(1)
	}

	lat, err := strconv.ParseFloat(argsWithoutProg[0], 64)
	if err != nil {
		panic(err.Error())
	}

	lgt, err := strconv.ParseFloat(argsWithoutProg[1], 64)
	if err != nil {
		panic(err.Error())
	}

	splitDate := strings.Split(argsWithoutProg[2], "-")
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
