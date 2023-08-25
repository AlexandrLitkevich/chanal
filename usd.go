package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func UsdHandlerRequest(w http.ResponseWriter, r *http.Request) {
	for {
		select {
			case <-time.After(10 * time.Second):
				fmt.Println("This time after")
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte("time is up"))

				return
			case data := <-GetUsd(firstReq):
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(data))
				return
		}
	}
}

func GetUsd(url string) <-chan string{
	out := make(chan string)
	go func() {
		defer close(out)
		resp, err := http.Get(url)
		if err != nil {
			fmt.Println("error")
			return
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
				fmt.Println(err)
		}
		out <- string(body[:])
	}()
	
	return out
}