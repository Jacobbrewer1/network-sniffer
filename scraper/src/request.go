package main

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/Jacobbrewer1/network-sniffer/scraper/src/entities"
	"io"
	"net/http"
	"strings"
	"time"
)

const apiUrl = "https://localhost:8443/network"

func sendRecords(records []entities.Record) error {
	// Remove when production certs are in place
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	body, err := json.Marshal(records)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, apiUrl, strings.NewReader(string(body)))
	if err != nil {
		return err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode == http.StatusOK {
		return nil
	}

	return processApiError(resp)
}

func processApiError(resp *http.Response) error {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return fmt.Errorf("api error: %s", string(body))
}
