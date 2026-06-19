package main

import (
	"context"
	"errors"
	"log/slog"
	"math/rand"
	"net/http"
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
		slog.Error(err.Error())
	}
}

func newServeMux() *http.ServeMux {
	mux := http.NewServeMux()

	timeout := timeoutMiddleware(100) // タイムアウト値 = 100ミリ秒（0.1秒）

	mux.Handle("GET /work", timeout(http.HandlerFunc(handleWork)))

	return mux
}

func handleWork(w http.ResponseWriter, r *http.Request) {
	msg, err := doRandomWork(r.Context()) // コンテキストを引き継ぎ、橋渡しする
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			w.WriteHeader(http.StatusGatewayTimeout)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
	} else {
		w.WriteHeader(http.StatusOK)
	}

	w.Write([]byte(msg + "\n"))
}

func timeoutMiddleware(ms int) func(h http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			ctx, cancelFunc := context.WithTimeout(ctx, time.Duration(ms)*time.Millisecond)
			defer cancelFunc()
			r = r.WithContext(ctx)

			h.ServeHTTP(w, r)
		})
	}
}

func doRandomWork(ctx context.Context) (string, error) {
	wait := rand.Intn(200) // 0~200までのランダムな整数を生成
	select {
	case <-time.After(time.Duration(wait) * time.Millisecond):
		// タイムアウト値の0.1秒未満で仕事を終えた場合、こちらが選択される
		return "制限時間内に仕事が終わりました", nil
	case <-ctx.Done():
		// タイムアウト値の0.1秒を超過した場合、チャネルクローズによりこちらが選択される
		return "制限時間内に仕事が終わりませんでした", ctx.Err()
	}
}
