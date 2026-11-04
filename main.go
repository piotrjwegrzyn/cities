package main

import (
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"
	"wegrzyn/cities/api"
	"wegrzyn/cities/crawler"
)

func main() {
	port := flag.Uint("port", 8080, "port for accessing app")
	url := flag.String("url", "http://api.geonames.org/searchJSON?featureClass=P&maxRows=50&orderby=population&username=hsample", "source of cities")
	timeout := flag.Duration("timeout", 10*time.Second, "timeout for accessing the external resource")

	flag.Parse()

	m, err := api.New(&crawler.Crawler{
		URL:     *url,
		Timeout: *timeout,
	})
	if err != nil {
		slog.Error("failed to start server", slog.String("error", err.Error()))
		os.Exit(1)
	}

	server := &http.Server{
		Addr:    ":" + fmt.Sprint(*port),
		Handler: m,
	}

	slog.Info(fmt.Sprintf("serving on address: http://127.0.0.1:%d/", *port))

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		slog.Error("server failed with error", slog.String("error", err.Error()))
		os.Exit(1)
	}
}
