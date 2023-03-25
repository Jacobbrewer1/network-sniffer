package main

import (
	"encoding/json"
	"github.com/Jacobbrewer1/network-sniffer/processor/src/entities"
	"github.com/Jacobbrewer1/network-sniffer/processor/src/response"
	"io"
	"log"
	"net/http"
)

func metricsController(w http.ResponseWriter, r *http.Request) {
	resp := response.NewResponse(w, r)

	got, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		return
	}

	var records []entities.Record
	if err := json.Unmarshal(got, &records); err != nil {
		log.Println(err)
		resp.BadRequest(err)
		return
	}

	sendToPrometheus(records)

	resp.Message(http.StatusOK, "Processed")
}

func sendToPrometheus(records []entities.Record) {
	const linkDown = "Link Down"
	for _, record := range records {
		if record.Status == linkDown {
			continue
		}

		promSentPacketsLifetime.WithLabelValues(record.Port).Set(float64(record.SentPacketsLifetime))
		promReceivedPacketsLifetime.WithLabelValues(record.Port).Set(float64(record.ReceivedPacketsLifetime))
		promCollisions.WithLabelValues(record.Port).Set(float64(record.SentBytesPerSecond))
		promSentPacketsBytesPerSecond.WithLabelValues(record.Port).Set(float64(record.SentBytesPerSecond))
		promReceivedPacketsBytesPerSecond.WithLabelValues(record.Port).Set(float64(record.ReceivedBytesPerSecond))
	}
}
