package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"log"
)

var (
	totalRequests = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_total_requests",
			Help: "Total requests receieved",
		},
		[]string{"path"},
	)

	promSentPacketsLifetime = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "sent_packets_lifetime",
			Help: "Number of packets sent during a lifetime",
		},
		[]string{"port"},
	)

	promReceivedPacketsLifetime = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "received_packets_lifetime",
			Help: "Number of packets received during a lifetime",
		},
		[]string{"port"},
	)

	promCollisions = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "packet_collisions",
			Help: "Number of packet collisions",
		},
		[]string{"port"},
	)

	promSentPacketsBytesPerSecond = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "sent_packets_bytes_second",
			Help: "Number of packets sent, bytes per second",
		},
		[]string{"port"},
	)

	promReceivedPacketsBytesPerSecond = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "received_packets_bytes_second",
			Help: "Number of packets sent, bytes per second",
		},
		[]string{"port"},
	)
)

func initPrometheus() {
	if err := prometheus.Register(totalRequests); err != nil {
		log.Fatalln(err)
	}

	if err := prometheus.Register(promSentPacketsLifetime); err != nil {
		log.Fatalln(err)
	}

	if err := prometheus.Register(promReceivedPacketsLifetime); err != nil {
		log.Fatalln(err)
	}

	if err := prometheus.Register(promCollisions); err != nil {
		log.Fatalln(err)
	}

	if err := prometheus.Register(promSentPacketsBytesPerSecond); err != nil {
		log.Fatalln(err)
	}

	if err := prometheus.Register(promReceivedPacketsBytesPerSecond); err != nil {
		log.Fatalln(err)
	}
}
