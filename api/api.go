package api

import (
	"fmt"
)

func FirstMyChan () {
	fmt.Println("run first my chan")
	cancelChan := make(chan struct{})
	defer close(cancelChan)

	gen := Gener(cancelChan, 10, 100)
	for v := range gen {
		fmt.Println("this data in chanal", v)
		
	}
}

func Gener(cancel <-chan struct{}, start int, stop int) <-chan int {
   out := make(chan int)
	 go func() {
		defer close(out)
		for i := start; i < stop; i++ {
			select {
				case out <- i:
				case <- cancel:
					return
			}
		}
	 }()

	 return out
}