package main

import (
	"flag"
	"log"
	"net"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const (
	appName = "Sonos exporter"
	defaultAddr = ":1915"
	metricsPath = "/metrics"
)

func main() {
	flagAddress := flag.String("address", defaultAddr, "Listen address")
	flag.Parse()

	prometheus.MustRegister(collectionErrors)
	prometheus.MustRegister(collector{})

	http.HandleFunc("/", indexHandler)
	http.Handle(metricsPath, promhttp.Handler())

	l, err := net.Listen("tcp", *flagAddress)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%s listening on http://%s", appName, l.Addr().String())
	log.Fatal(http.Serve(l, nil))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	responseHTML := `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>`+appName+`</title>
</head>
<body>
	<h1>`+appName+`</h1>
	<a href="`+metricsPath+`">`+metricsPath+`</a>
</body>
</html>`

	w.Write([]byte(responseHTML))
}
