package main

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/takaya-47/learn-go/ex14/prob03/log"
)

func main() {
	s := http.Server{
		Addr:         ":8080",
		Handler:      log.Middleware(http.HandlerFunc(message)),
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 90 * time.Second,
		IdleTimeout:  120 * time.Second,
	}
	err := s.ListenAndServe()
	if err != nil {
		slog.Error(err.Error())
	}
}

func message(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log.Log(ctx, log.Debug, "This is a debug message")
	log.Log(ctx, log.Info, "This is an info message")

	w.Write([]byte("Done"))
}
