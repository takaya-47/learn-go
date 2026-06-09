package main

import (
	"log"
	"net/http"
	"time"
)

func main() {

	s := http.Server{
		Addr:         ":8080",
		Handler:      createServerMux(),
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 90 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	err := s.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func createServerMux() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}

		w.Write([]byte(time.Now().Format(time.RFC3339)))
	})

	// 以下のようにGETを明示的に指定することもでき、その場合はGoが自動的にHTTPメソッドをチェックし、不一致なら405エラーを返してくれる。
	// mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte(time.Now().Format(time.RFC3339)))
	// })

	return mux
}
