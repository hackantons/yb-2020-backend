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

func main() {

	in, err := os.Open("msft.csv")
	if err != nil {
		panic(err)
	}
	defer in.Close()

	weekData := []*WeekData{}

	if err := gocsv.UnmarshalFile(in, &weekData); err != nil {
		panic(err)
	}

	value, _ := json.Marshal(weekData)

	http.HandleFunc("/shares", func(w http.ResponseWriter, r *http.Request) {

		enableCors(&w)

		fmt.Fprint(w, string(value))

	})

	log.Fatal(http.ListenAndServe(":5080", nil))
}

// Just for testing purposes
func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}
