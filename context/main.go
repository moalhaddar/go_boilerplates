package main

import (
	"context"
	"errors"
	"io"
	"log"
	"net/http"
	"time"
)

func main() {
	ctx := context.Background()
	requestContext, cancelRequest := context.WithTimeout(ctx, 1*time.Second)
	defer cancelRequest()

	req, _ := http.NewRequestWithContext(requestContext, "GET", "https://google.com", nil)

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			log.Println("Ayo bruv we got a deadline over here")
		}
		return
	}

	body, _ := io.ReadAll(res.Body)

	log.Printf("Read: %s", body)
}
