package log

import (
	"context"
	"fmt"
	"net/http"
)

type Level string

const (
	Debug Level = "debug"
	Info  Level = "info"
)

// コンテキストのキーとして使用するため、非公開。
type logKey int

const (
	_        logKey = iota // 0は使用しない
	logLevel               // logLevelは1で型はlogKey
)

// contextWithLogLevel は、与えられたログレベルを持つ新しいコンテキストを返します。
func ContextWithLogLevel(ctx context.Context, level Level) context.Context {
	return context.WithValue(ctx, logLevel, level)
}

// LogLevelFromContext は、コンテキストからログレベルを取得します。
// ログレベルが存在する場合は、その値とtrueを返し、存在しない場合は空の文字列とfalseを返します。
func LogLevelFromContext(ctx context.Context) (Level, bool) {
	// Valueがnilを返し、nilに型アサーションした場合、levelにはLevel型のゼロ値（空文字）が入る
	level, ok := ctx.Value(logLevel).(Level)
	return level, ok
}

func Middleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		level := r.URL.Query().Get("log_level")
		if Level(level) != Debug && Level(level) != Info {
			http.Error(w, "Invalid log level", http.StatusBadRequest)
			return
		}

		// コンテキストにログレベルをセット
		ctx := ContextWithLogLevel(r.Context(), Level(level))

		// ログレベルを追加したコンテキストを持つリクエストを生成
		r = r.WithContext(ctx)

		h.ServeHTTP(w, r)
	})
}

func Log(ctx context.Context, level Level, message string) {
	inLevel, ok := LogLevelFromContext(ctx)
	if !ok {
		return
	}

	if level == Debug && inLevel == Debug {
		fmt.Println(message)
	}
	if level == Info && (inLevel == Debug || inLevel == Info) {
		fmt.Println(message)
	}
}
