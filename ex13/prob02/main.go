package main

import (
	"log"
	"log/slog"
	"net"
	"net/http"
	"os"
	"strings"
	"time"
)

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

	mux.Handle("/", clientIpLogger(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}

		w.Write([]byte(time.Now().Format(time.RFC3339)))
	})))

	// 以下のようにGETを明示的に指定することもでき、その場合はGoが自動的にHTTPメソッドをチェックし、不一致なら405エラーを返してくれる。
	// mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte(time.Now().Format(time.RFC3339)))
	// })

	return mux
}

func clientIpLogger(next http.Handler) http.Handler {
	logger := newLogger()

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Info("Received request", slog.String("client_ip", clientIP(r)))

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
	return r.RemoteAddr
}
