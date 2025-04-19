package main

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
)

type HelloRequest struct {
	Name string `json:"fullName"`
}

type HelloResponse struct {
	Hello string `json:"dasd"`
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var request HelloRequest
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			if errors.Is(err, io.ErrUnexpectedEOF) {
				w.Write([]byte("Yo bro this is unexpected eof.."))
				return
			}
		}

		var resposne HelloResponse
		resposne.Hello = request.Name
		err = json.NewEncoder(w).Encode(resposne)
		if err != nil {
			log.Printf("Failed to encode payload\n")
		}
	})
	http.ListenAndServe("0.0.0.0:8080", nil)
	log.Printf("Hello world\n")
}
