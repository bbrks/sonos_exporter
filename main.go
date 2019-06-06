package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var	flagAddress = flag.String("address", "localhost:1915", "Listen address")

func init() {
	prometheus.MustRegister(collectionErrors)
	prometheus.MustRegister(collector{})
}

func main() {
	flag.Parse()

	log.Printf("Sonos exporter listening on %s", *flagAddress)
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(*flagAddress, nil))
}
