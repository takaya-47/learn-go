package main

import (
	"errors"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"takaya-47/learn-go/ex15/internal/processor"
	"time"
)

func newController(out chan []byte) http.Handler {
	var numSent int
	var numRejected int
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		numSent++
		// take in data
		data, err := io.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Bad Input"))
			return
		}
		// write it to the queue in raw format
		select {
		case out <- data:
			// success!
		default:
			// if the channel is backed up, return an error
			numRejected++
			w.WriteHeader(http.StatusServiceUnavailable)
			w.Write([]byte("Too Busy: " + strconv.Itoa(numRejected)))
			return
		}
		w.WriteHeader(http.StatusAccepted)
		w.Write([]byte("OK: " + strconv.Itoa(numSent)))
	})
}

func run() error {
	// set everything up
	ch1 := make(chan []byte, 100)
	ch2 := make(chan processor.Result, 100)
	go processor.DataProcessor(ch1, ch2)

	f, err := os.Create("results.txt")
	if err != nil {
		return err
	}
	defer f.Close()

	go processor.WriteData(ch2, f)

	s := http.Server{
		Addr:         ":8080",
		Handler:      newController(ch1),
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 90 * time.Second,
		IdleTimeout:  120 * time.Second,
	}
	err = s.ListenAndServe()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	return nil
}

func main() {
	if err := run(); err != nil {
		log.Fatalf("Error when starting server: %v", err)
	}
}
