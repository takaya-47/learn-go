package main

import (
	"errors"
	"io"
	"log"
	"net/http"
	"os"
	"takaya-47/learn-go/ex15/internal/handler"
	"takaya-47/learn-go/ex15/internal/processor"
	"time"
)

func run(w io.Writer) error {
	ch1 := make(chan []byte, 100)
	ch2 := make(chan processor.Result, 100)
	go processor.DataProcessor(ch1, ch2)

	go processor.WriteData(ch2, w)

	s := http.Server{
		Addr:         ":8080",
		Handler:      handler.NewHandler(ch1),
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 90 * time.Second,
		IdleTimeout:  120 * time.Second,
	}
	err := s.ListenAndServe()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	return nil
}

func main() {
	f, err := os.Create("results.txt")
	if err != nil {
		log.Fatalf("Error when creating results.txt: %v", err)
	}
	defer f.Close()

	if err := run(f); err != nil {
		log.Fatalf("Error when starting server: %v", err)
	}
}
