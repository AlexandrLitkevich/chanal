package main

import (
	"fmt"
	"io"
	"net/http"
	"sync"
)


const (
	firstReq  = "https://api.open-meteo.com/v1/forecast?latitude=60&longitude=100&hourly=temperature_2m,precipitation_probability"

	secondRequest = "https://api.open-meteo.com/v1/forecast?latitude=9.4844&longitude=-2.4579&hourly=dewpoint_2m,pressure_msl,cloudcover_high,visibility"
)

func HandlerRequestEasy(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello, new request w.write"))
}

func AllCurrencyHandler(w http.ResponseWriter, r *http.Request) {
                        
	resultArrChans := GetAllCurrency()
	mergeRusult := MergeParallel(resultArrChans[0], resultArrChans[1])
	
	var data string
	for val := range mergeRusult {
		data += val
	}
	w.Write([]byte(data))
}

func GetUsdWithCanceled(cancel chan struct{},url string) <-chan string{
	out := make(chan string)

	go func() {
		defer close(out)
		resp, err := http.Get(url)


		if err != nil {
			fmt.Println("error")
			close(cancel)
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
				fmt.Println(err)
				close(cancel)
		}
		out <- string(body[:])
	}()

	
		return out
}

func GetAllCurrency() [] <-chan string {
	//тут добавить гоурутину

	var wg sync.WaitGroup
	cancel := make(chan struct{})


	out := make([] <-chan string, 2)
	
	wg.Add(1)
	go func() {
		defer wg.Done()
		out[0] = GetUsdWithCanceled(cancel,firstReq)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		out[1] = GetUsdWithCanceled(cancel,secondRequest)
	}()

	wg.Wait()
	return out
}

func MergeParallel(in1, in2 <-chan string) <-chan string {
	var wg sync.WaitGroup
	wg.Add(2)

  out := make(chan string)
  go func() {
		defer wg.Done()
		for val := range in1 {
			out <- val
		}
	}()

	go func() {
		defer wg.Done()
		for val := range in2 {
			out <- val
		}
	}()

	// ждем, пока исчерпаются оба входных канала,
  // после чего закрываем выходной
	go func() {
		wg.Wait()
		close(out)
	}()

  return out 
}
