package main

import (
	"log"
	"net/http"
	"os"

	"github.com/SantiagoBedoya/restfull-webservices/chapter9/encryptService/helpers"
	kitlog "github.com/go-kit/kit/log"
	kitprometeus "github.com/go-kit/kit/metrics/prometheus"
	httptrasport "github.com/go-kit/kit/transport/http"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	logger := kitlog.NewLogfmtLogger(os.Stderr)
	fieldKeys := []string{"method", "error"}
	requestCount := kitprometeus.NewCounterFrom(stdprometheus.CounterOpts{
		Namespace: "encrpytion",
		Subsystem: "my_service",
		Name:      "request_count",
		Help:      "number of request received",
	}, fieldKeys)
	requestLatency := kitprometeus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace: "encryption",
		Subsystem: "my_service",
		Name:      "request_latency_microseconds",
		Help:      "total duration of requests in microseconds",
	}, fieldKeys)

	var svc helpers.EncryptService

	svc = helpers.EncryptServiceInstance{}
	svc = helpers.LogginMiddleware{Logger: logger, Next: svc}
	svc = helpers.InstrumentingMiddleware{RequestCount: requestCount, RequestLatency: requestLatency, Next: svc}

	encryptHandler := httptrasport.NewServer(helpers.MakeEncryptEndpoint(svc), helpers.DecodeEncryptRequest, helpers.EncodeResponse)
	decryptHandler := httptrasport.NewServer(helpers.MakeDecryptEndpoint(svc), helpers.DecodeDecryptRequest, helpers.EncodeResponse)

	http.Handle("/encrypt", encryptHandler)
	http.Handle("/decrypt", decryptHandler)
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":8080", nil))
}
