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
	DayOfWeek  string `json:"day_of_week"`
	DayOfMonth int    `json:"day_of_month"`
	Month      string `json:"month"`
	Year       int    `json:"year"`
	Hour       int    `json:"hour"`
	Minute     int    `json:"minute"`
	Second     int    `json:"second"`
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
		now := time.Now()
		switch r.Header.Get("Accept") {
		case "application/json":
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			err := json.NewEncoder(w).Encode(fullDateTime{
				DayOfWeek:  now.Weekday().String(),
				DayOfMonth: now.Day(),
				Month:      now.Month().String(),
				Year:       now.Year(),
				Hour:       now.Hour(),
				Minute:     now.Minute(),
				Second:     now.Second(),
			})
			if err != nil {
				http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
				return
			}
		case "text/plain":
			w.Header().Set("Content-Type", "text/plain; charset=utf-8")
			w.Write([]byte(now.Format(time.RFC3339) + "\n"))
		default:
			w.Header().Set("Content-Type", "text/plain; charset=utf-8")
			w.Write([]byte(now.Format(time.RFC3339) + "\n"))
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
