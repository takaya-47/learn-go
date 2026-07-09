package handler

import (
	"io"
	"net/http"
	"strconv"
	"sync"
)

type counter struct {
	l           sync.Mutex
	numSent     int
	numRejected int
}

func (c *counter) incrementSent() {
	c.l.Lock()
	defer c.l.Unlock()
	c.numSent++
}
func (c *counter) incrementRejected() {
	c.l.Lock()
	defer c.l.Unlock()
	c.numRejected++
}

// TODO: counter構造体のメソッドを使うように修正する
func NewHandler(out chan []byte) http.Handler {
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
