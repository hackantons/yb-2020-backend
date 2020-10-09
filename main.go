package main

import (
	"encoding/json"
	"fmt"
	"github.com/gocarina/gocsv"
	"log"
	"net/http"
	"os"
)

type WeekData struct {
	Timestamp string  `csv:"timestamp"`
	Open      float32 `csv:"open"`
	High      float32 `csv:"high"`
	Low       float32 `csv:"low"`
	Close     float32 `csv:"close"`
	Volume    float32 `csv:"volume"`
}

type Share struct {
	name   string
	symbol string
}

var (
	shares []Share
)

func main() {
	generateShares()

	in, err := os.Open("msft.csv")
	if err != nil {
		panic(err)
	}
	defer in.Close()

	weekData := []*WeekData{}

	if err := gocsv.UnmarshalFile(in, &weekData); err != nil {
		panic(err)
	}

	http.HandleFunc("/shares", func(w http.ResponseWriter, r *http.Request) {

		enableCors(&w)
		value, _ := json.Marshal(weekData)
		fmt.Fprint(w, string(value))

	})

	http.HandleFunc("/health/ready", func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)
		w.WriteHeader(http.StatusOK)
	})

	http.HandleFunc("/shares", func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)

		value, _ := json.Marshal(shares)
		fmt.Fprint(w, string(value))
	})

	log.Fatal(http.ListenAndServe(":5080", nil))
}

// Just for testing purposes
func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func generateShares() {
	shares = append(shares, Share{name: "Amazon", symbol: "AMZN"})
	shares = append(shares, Share{name: "Facebook", symbol: "FB"})
	shares = append(shares, Share{name: "JP Morgan", symbol: "JPM"})
	shares = append(shares, Share{name: "Coca Cola", symbol: "KO"})
	shares = append(shares, Share{name: "Mastercard", symbol: "MA"})
	shares = append(shares, Share{name: "Mc Donalds", symbol: "MCD"})
	shares = append(shares, Share{name: "Microsoft", symbol: "MSFT"})
	shares = append(shares, Share{name: "Philip Morris", symbol: "PM"})
}
