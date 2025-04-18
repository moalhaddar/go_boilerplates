package main

import (
	"net/http"
	_ "net/http/pprof"

	"alhaddar.dev/graceful_http_shutdown/server"
)

const (
	BYTE     = 1
	KILOBYTE = 1024 * BYTE
	MEGABYTE = 1024 * KILOBYTE
	GIGABYTE = 1024 * MEGABYTE
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var body []byte
		for range GIGABYTE { // one gig of data for the lulz
			body = append(body, []byte("A")...)
		}
		w.Write(body)
	})
	server := server.NewServer("0.0.0.0:8080", nil)

	server.StartServer()
}
