package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type HelloRequest struct {
	Name string `json:"fullName"`
}

type HelloResponse struct {
	Hello string `json:"dasd"`
}

type JSONHandler[req any, res any] func(r req) (res, error)

func wrapJsonHandler[req any, res any](handler JSONHandler[req, res]) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var request req
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			w.Write([]byte("failed to decode json"))
			return
		}

		response, err := handler(request)
		if err != nil {
			w.Write([]byte("request failed for some reason.."))
			return
		}

		err = json.NewEncoder(w).Encode(response)
		if err != nil {
			log.Printf("Failed to encode payload\n")
			return
		}
	})
}

func handler(request HelloRequest) (HelloResponse, error) {
	var response HelloResponse
	response.Hello = request.Name
	return response, nil
}

func main() {
	http.Handle("/", wrapJsonHandler(handler))
	http.ListenAndServe("0.0.0.0:8080", nil)
	log.Printf("Hello world\n")
}
