package main

import (
	"fmt"
	"net/http"
)



func main () {
	fmt.Println("run server")

	http.HandleFunc("/", HandlerRequestEasy)
	http.HandleFunc("/first", UsdHandlerRequest)
	http.HandleFunc("/second", EuroHandlerRequest)
	http.HandleFunc("/all", AllHandlerWithContext)
	http.ListenAndServe(":8000", nil)
	// HandlerMakeRequest()
}




