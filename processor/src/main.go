package main

import (
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
	"processor/src/config"
	"processor/src/response"
)

type controller func(w http.ResponseWriter, r *http.Request)

func notFound(w http.ResponseWriter, r *http.Request) {
	response.NewResponse(w, r).NotFound()
}

func methodNotAllowed(w http.ResponseWriter, r *http.Request) {
	response.NewResponse(w, r).MethodNotAllowed()
}

func auth(c controller) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		totalRequests.WithLabelValues(r.URL.Path).Inc()
		c(w, r)
	}
}

func prometheusAuth() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Authorization") == "t1e2s3t4" {
			log.Println("prometheus authorised")
		} else {
			log.Println("prometheus not authorised")
		}

		promhttp.Handler().ServeHTTP(w, r)
	}
}

func init() {
	initLogging()
	initFlags()
	initPrometheus()
}

func isFlagProvided(flagName string) (isProvided bool) {
	isProvided = false

	flag.Visit(func(f *flag.Flag) {
		if f.Name == flagName {
			isProvided = true
		}
	})

	return isProvided
}

func initLogging() {
	log.SetFlags(log.LstdFlags)
}

func initFlags() {
	configPath := flag.String(config.FlagName, "./config/config.yml", "Provides the location of the config file (required)")

	flag.Parse()

	config.Location = *configPath
}

func initConfig() {
	if !isFlagProvided(config.FlagName) {
		log.Fatalln("no config flag provided")
	}

	if err := config.CreateConfig(); err != nil {
		log.Println(err)
	}
}

func main() {
	initConfig()

	r := mux.NewRouter()

	r.HandleFunc("/network", auth(metricsController)).Methods(http.MethodPost)

	scrape := r.PathPrefix("/scrape").Subrouter()
	scrape.HandleFunc("/prometheus", prometheusAuth())

	r.NotFoundHandler = http.HandlerFunc(notFound)
	r.MethodNotAllowedHandler = http.HandlerFunc(methodNotAllowed)

	http.Handle("/", r)

	log.Printf("listening at %s...\n", config.Cfg.Setup.ListeningPort)
	if err := http.ListenAndServeTLS(fmt.Sprintf(":%s", config.Cfg.Setup.ListeningPort),
		config.Cfg.Setup.CertPath, config.Cfg.Setup.KeyPath, nil); err != nil {
		log.Fatalln(err)
	}
}
