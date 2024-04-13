/*
Package gohttpmetrics knows how to measure http metrics in VictoriaMetrics format.
it comes with a middleware that can be used for the main Go net/http handler:

	package main

	import (
		"log"
		"net/http"

		victoriametrics "github.com/VictoriaMetrics/metrics"
		victoria "github.com/dottedmag/go-http-metrics/metrics/victoria"
		httpmiddleware "github.com/dottedmag/go-http-metrics/middleware"
		httpstdmiddleware "github.com/dottedmag/go-http-metrics/middleware/std"
	)

	func main() {
		// Create our middleware.
		mdlw := victoria.New(victoria.Config{})

		// Our handler.
		myHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("hello world!"))
		})
		h := httpstdmiddleware.Handler("", mdlw, myHandler)

		// Serve metrics.
		log.Printf("serving metrics at: %s", ":9090")
		go http.ListenAndServe(":9090", promhttp.Handler())

		// Serve our handler.
		log.Printf("listening at: %s", ":8080")
		if err := http.ListenAndServe(":8080", h); err != nil {
			log.Panicf("error while serving: %s", err)
		}
	}
*/
package gohttpmetrics

// blank imports help docs.
import (
	// Import metrics package.
	_ "github.com/dottedmag/go-http-metrics/metrics"
	// Import middleware package.
	_ "github.com/dottedmag/go-http-metrics/middleware"
)
