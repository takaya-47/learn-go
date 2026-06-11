package main

import (
	"encoding/json"
	"log"
	"log/slog"
	"net"
	"net/http"
	"os"
	"strings"
	"time"
)

type fullDateTime struct {
	dayOfWeek  string `json:"day_of_week"`
	dayOfMonth int    `json:"day_of_month"`
	month      string `json:"month"`
	year       int    `json:"year"`
	hour       int    `json:"hour"`
	minute     int    `json:"minute"`
	second     int    `json:"second"`
}

func main() {

	s := http.Server{
		Addr:         ":8080",
		Handler:      newServeMux(),
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 90 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	err := s.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func newServeMux() *http.ServeMux {
	mux := http.NewServeMux()

	mux.Handle("GET /", withClientIPLogging(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Header.Get("Accept") {
		case "application/json":
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			now := time.Now()
			err := json.NewEncoder(w).Encode(fullDateTime{
				dayOfWeek:  now.Weekday().String(),
				dayOfMonth: now.Day(),
				month:      now.Month().String(),
				year:       now.Year(),
				hour:       now.Hour(),
				minute:     now.Minute(),
				second:     now.Second(),
			})
			if err != nil {
				http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
				return
			}
		case "text/plain":
			w.Header().Set("Content-Type", "text/plain; charset=utf-8")
			w.Write([]byte(time.Now().Format(time.RFC3339)))
		default:
			w.Write([]byte(time.Now().Format(time.RFC3339)))
		}
	})))

	return mux
}

func withClientIPLogging(next http.Handler) http.Handler {
	logger := newLogger()

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Info("Received request",
			slog.String("client_ip", clientIP(r)),
			slog.String("method", r.Method),
			slog.String("path", r.URL.Path),
		)

		next.ServeHTTP(w, r)
	})
}

func newLogger() *slog.Logger {
	jsonHandler := slog.NewJSONHandler(os.Stderr, nil)
	return slog.New(jsonHandler)
}

func clientIP(r *http.Request) string {
	// X-Forwarded-For: "client, proxy1, proxy2" の形式。先頭が元のクライアント。
	if xff := r.Header.Get("X-Forwarded-For"); xff != "" {
		first, _, _ := strings.Cut(xff, ",")
		return strings.TrimSpace(first)
	}

	if xrip := r.Header.Get("X-Real-IP"); xrip != "" {
		return strings.TrimSpace(xrip)
	}

	// プロキシが無い場合は RemoteAddr（IP:port 形式）から IP 部分だけ取り出す。
	if ip, _, err := net.SplitHostPort(r.RemoteAddr); err == nil {
		return ip
	}

	// フォールバック用
	return r.RemoteAddr
}
