package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

func AllHandlerWithContext(w http.ResponseWriter, r *http.Request) {    
	ctx := context.Background()
	childCtx, cancel := context.WithCancel(ctx)
	defer cancel()
	parseResponse := ParseResponse(ctx, cancel)
	data := <-parseResponse
	w.Write([]byte(data))

	for {
		select {
		case <-time.After(4 * time.Second):
			fmt.Println("Time")
			return
		case <-childCtx.Done():
			fmt.Println("It worked childCtx.Done")
			return
		// case <-parseResponse:
		// 	data, ok := <-parseResponse
		// 	fmt.Println("parse responce in select", ok)
		// 	fmt.Println("parse responce in select", data)
		// 	w.Write([]byte(data))
		// 	return
		// Почему тут я не могу получить данные?????
		// остутсвие return  порождает бесконгечный цикл
		}
	}


	
	
	
}

func ParseResponse(ctx context.Context, cancel context.CancelFunc) <-chan string {
	out := make(chan string)

	// go func ()  {
	// 	resultArrChans := GetAllData(ctx, cancel)
	// }()
	
	go func() {
		resultArrChans := GetAllData(ctx, cancel)
		mergeRusult := MergeParallel(resultArrChans[0], resultArrChans[1]) 
		defer close(out)

		var data string
		for val := range mergeRusult {
			data += val
		}
		out <-data
	}()
	return out
}

func GetAllData(ctx context.Context,cancel context.CancelFunc) [] <-chan string {
	var wg sync.WaitGroup
	out := make([] <-chan string, 2)
	
	wg.Add(1)
	go func() {
		defer wg.Done()
		out[0] = GetDataWithCanceled(cancel,firstReq)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		out[1] = GetDataWithCanceled(cancel,secondRequest)
	}()

	wg.Wait()
	return out
}

func GetDataWithCanceled(cancel context.CancelFunc,url string) <-chan string {
	out := make(chan string)

	go func() {
		defer close(out)
		resp, err := http.Get(url)

		if err != nil {
			fmt.Println("error", err)
			cancel() //Выше все отмениться
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
				fmt.Println(err)
				cancel() //Выше все отмениться
		}
		out <-string(body[:])
	}()
		return out
}