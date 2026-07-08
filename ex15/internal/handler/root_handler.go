package handler

import (
	"io"
	"net/http"
	"strconv"
)

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
