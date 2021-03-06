package app

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/moonliightz/telegram-webhookinfo-exporter/internal/config"
	"github.com/moonliightz/telegram-webhookinfo-exporter/internal/model"
	"github.com/moonliightz/telegram-webhookinfo-exporter/internal/telegram"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Run starts the application
func Run() {
	configPath := flag.String("config", "./config.yml", "path to config.yml")
	flag.Parse()

	config, err := config.Load(*configPath)
	if err != nil {
		log.Fatal(err.Error())
	}

	pendingUpdateCount := prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: config.Prometheus.Namespace,
		Subsystem: config.Prometheus.Subsystem,
		Name:      config.Prometheus.Name,
	})
	prometheus.MustRegister(pendingUpdateCount)

	go func() {
		var winfo *model.WebhookInfo
		var logPraefix string

		if config.Prometheus.Namespace != "" {
			logPraefix += config.Prometheus.Namespace + "_"
		}
		if config.Prometheus.Subsystem != "" {
			logPraefix += config.Prometheus.Subsystem + "_"
		}
		logPraefix += config.Prometheus.Name // set by default

		for {
			winfo, err = telegram.LoadWebhookInfo(config)
			if err != nil {
				panic(err)
			}
			pendingUpdateCount.Set(float64(winfo.Result.PendingUpdateCount))
			log.Print(logPraefix, ": ", winfo.Result.PendingUpdateCount)

			time.Sleep(time.Duration(config.App.Interval) * time.Second)
		}
	}()

	log.Printf("Starting http server on %s:%d", config.HTTP.Addr, config.HTTP.Port)
	http.Handle("/metrics", promhttp.Handler())
	err = http.ListenAndServe(fmt.Sprintf("%s:%d", config.HTTP.Addr, config.HTTP.Port), nil)
	if err != nil {
		panic(err)
	}
}
