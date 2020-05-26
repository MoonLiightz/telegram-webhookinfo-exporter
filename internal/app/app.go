package app

import (
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
	config, err := config.Load()
	if err != nil {
		panic(err)
	}

	pendingUpdateCount := prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: config.Prometheus.Namespace,
		Subsystem: config.Prometheus.Subsystem,
		Name:      config.Prometheus.Name,
	})
	prometheus.MustRegister(pendingUpdateCount)

	go func() {
		var winfo *model.WebhookInfo
		for {
			winfo, err = telegram.LoadWebhookInfo(config)
			if err != nil {
				panic(err)
			}
			pendingUpdateCount.Set(float64(winfo.Result.PendingUpdateCount))
			log.Print("pending_update_count: ", winfo.Result.PendingUpdateCount)

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
