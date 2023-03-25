package main

import (
	"errors"
	"flag"
	"log"
	"scraper/src/config"
	"scraper/src/entities"
	"strconv"
	"time"
)

func start() {
	for {
		if err := run(); err != nil {
			log.Println(err)
		}

		log.Println("waiting")
		time.Sleep(time.Second * 5)
	}
}

func run() error {
	parsedRecords, err := scrape()
	if err != nil {
		return err
	}

	if parsedRecords == nil {
		return errors.New("records was empty")
	}

	records := constructRecords(parsedRecords)

	return sendRecords(records)
}

func constructRecords(parsedRecords [][]string) []entities.Record {
	var records []entities.Record
	for i, r := range parsedRecords {
		if i == 0 {
			// ignore headers of table
			continue
		}

		record := entities.Record{
			Port:   r[0],
			Status: r[1],
			UpTime: r[7],
		}

		if sentPacketsLife, err := strconv.Atoi(r[2]); err == nil {
			record.SentPacketsLifetime = int64(sentPacketsLife)
		}

		if receivedPacketsLife, err := strconv.Atoi(r[3]); err == nil {
			record.SentPacketsLifetime = int64(receivedPacketsLife)
		}

		if collisions, err := strconv.Atoi(r[4]); err == nil {
			record.Collisions = int64(collisions)
		}

		if sentBytesPS, err := strconv.Atoi(r[5]); err == nil {
			record.SentBytesPerSecond = int64(sentBytesPS)
		}

		if receivedBytesPS, err := strconv.Atoi(r[5]); err == nil {
			record.ReceivedBytesPerSecond = int64(receivedBytesPS)
		}

		records = append(records, record)
	}
	return records
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

func init() {
	initLogging()
	initFlags()
}

func initLogging() {
	log.SetFlags(log.LstdFlags)
}

func initConfig() {
	if !isFlagProvided(config.FlagName) {
		log.Fatalln("no config flag provided")
	}

	if err := config.CreateConfig(); err != nil {
		log.Println(err)
	}
}

func initFlags() {
	configPath := flag.String(config.FlagName, "./config/config.yml", "Provides the location of the config file (required)")

	flag.Parse()

	config.Location = *configPath
}

func main() {
	log.Println("setup")
	initConfig()

	log.Println("running")
	start()
}
