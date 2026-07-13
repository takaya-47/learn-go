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

func NewHandler(out chan []byte) http.Handler {
	c := &counter{}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c.incrementSent()
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
			c.incrementRejected()
			w.WriteHeader(http.StatusServiceUnavailable)
			w.Write([]byte("Too Busy: " + strconv.Itoa(c.numRejected)))
			return
		}
		w.WriteHeader(http.StatusAccepted)
		w.Write([]byte("OK: " + strconv.Itoa(c.numSent)))
	})
}
