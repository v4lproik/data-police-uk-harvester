# data-police-uk-harvester

A simple Go application for educational purpose that consumes data from the RESTful webservice of the UK data police.

## Installation

```
cd data-police-uk-harvester
glide install
```

## Run
### Cli
```
✘ v4lproik@v4lproik-laptop  ~/Programmation/src/github.com/v4lproik/data-police-uk-harvester   master ●  go run main.go cli 52.63535488432963 -1.1374282836914062     
12:24:48 ERROR main main.go:25 usage ./data-police-uk-harvester cli <latitude> <longitude> <date>
exit status 1
 ✘ v4lproik@v4lproik-laptop  ~/Programmation/src/github.com/v4lproik/data-police-uk-harvester   master ●  go run main.go cli 52.63535488432963 -1.1374282836914062 2015-9
12:24:59 INFO  connector opendata.go:49 Get data from https://data.police.uk/api/stops-street?lat=52.63535488432963&lng=-1.1374282836914062&date=2015-9
12:25:00 INFO  main main.go:61 response %!(EXTRA []connector.Stop=[{25-34 Black or Black British -......... 
```

### Web Interface
´´´
go run main.go web-server
´´´

1. Go to http://localhost:9090/
2. Click to the english town of your choice to display the police stops information of this area