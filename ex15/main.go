package main

import (
	"context"
	"errors"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"takaya-47/learn-go/ex15/internal/handler"
	"takaya-47/learn-go/ex15/internal/processor"
	"time"
)

func run(ctx context.Context, w io.Writer) error {
	ch1 := make(chan []byte, 100)
	ch2 := make(chan processor.Result, 100)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		processor.DataProcessor(ch1, ch2)
	}()
	go func() {
		defer wg.Done()
		processor.WriteData(ch2, w)
	}()

	s := http.Server{
		Addr:         ":8080",
		Handler:      handler.NewHandler(ch1),
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 90 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	serverShutdown := make(chan struct{})
	go func() {
		// シグナルを受け取るまでここで待機
		<-ctx.Done()

		// シグナルを受け取ったらグレースフルシャットダウンを開始。
		log.Println("Shutting down server...")
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := s.Shutdown(shutdownCtx); err != nil {
			log.Printf("Error during server shutdown: %v", err)
		}
		close(serverShutdown)
	}()

	err := s.ListenAndServe()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	// ListenAndServeがErrServerClosedを返した時（シャットダウン時）、ここに到達する。
	// serverShutdownチャネルが閉じられるまで待機。
	<-serverShutdown
	close(ch1)
	wg.Wait()
	return nil
}

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	f, err := os.Create("results.txt")
	if err != nil {
		log.Fatalf("Error when creating results.txt: %v", err)
	}
	defer f.Close()

	if err := run(ctx, f); err != nil {
		log.Fatalf("Error when starting server: %v", err)
	}
}
