package main

import (
    "fmt"
    "time"
    "math/rand"
    "net/http"

    "github.com/prometheus/client_golang/prometheus"
    "github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
    // Create a new gauge metric to track the temperature
    temperatureGauge := prometheus.NewGauge(prometheus.GaugeOpts{
        Name: "random_temperature_celsius",
        Help: "A random temperature value in Celsius.",
    })

    // Register the metric with Prometheus
    prometheus.MustRegister(temperatureGauge)

    // Set up a ticker to update the temperature value every second
    ticker := time.NewTicker(1 * time.Second)
    go func() {
        for range ticker.C {
            // Update the temperature with a random value between -10 and 40
            temperatureGauge.Set(rand.Float64()*50 - 10)
        }
    }()

    fmt.Println("Prometheus metrics available at http://localhost:8080/metrics")

    // Expose the Prometheus metrics endpoint
    http.Handle("/metrics", promhttp.Handler())
    http.ListenAndServe(":8080", nil)
}
