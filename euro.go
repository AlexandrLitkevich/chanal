package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)


func EuroHandlerRequest(w http.ResponseWriter, r *http.Request) {
	// go GetUsd(usdRequest, ch)
	for {
		select {
			case <-time.After(10 * time.Second): //Завершение по таймауту
				fmt.Println("This time after")
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte("time is up"))
				return
			case data := <-GetEuro(secondRequest):
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(data))
				return
		}
	}
}

func GetEuro(url string) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		resp, err := http.Get(url)
		if err != nil {
			fmt.Println("error")
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