# cities

Simple server for counting cities with prefix based on data from [here](http://api.geonames.org/searchJSON?featureClass=P&maxRows=50&orderby=population&username=hsample).

## Prerequisites

Install [Go 1.27.3](https://go.dev/dl/). Optionally install `make`.

## Build

Build with `make`:
```sh
make build
```

Or directly:
```sh
go build -o bin/ .
```

## Run

Simply:
```sh
make run
```

Or directly:
```sh
bin/cities
```

You can put some parameters based on the output
```sh
bin/cities -h
```
```
Usage of bin/cities:
  -port uint
        port for accessing app (default 8080)
  -timeout duration
        timeout for accessing the external resource (default 10s)
  -url string
        source of cities (default "http://api.geonames.org/searchJSON?featureClass=P&maxRows=50&orderby=population&username=hsample")
```