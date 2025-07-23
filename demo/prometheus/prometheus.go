package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
	"time"
)

func main() {
	gauge := prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "xin_test",
		Help: "this is xin test",
	})

	prometheus.MustRegister(gauge)

	var i int
	go func() {
		for {
			i++
			if i%2 == 0 {
				gauge.Inc()
			}
			time.Sleep(time.Second)
		}
	}()

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":1234", nil)
}
