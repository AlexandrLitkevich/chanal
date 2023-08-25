package main


 /* 
 	chan: для чтения и записи (по умолчанию);
	chan<- : только для записи (send-only);
	<-chan: только для чтения (receive-only).
 */

import (
	"fmt"
)

func main () {
	fmt.Println("run main")
	FirstMyChan()
}

func FirstMyChan () {
	fmt.Println("run first my chan")
	cancelChan := make(chan struct{})
	defer close(cancelChan)

	gen := Gener(cancelChan, 10, 100)
	for v := range gen {
		if v == 50 {
			break
		}
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
					// Прекратиться этот кейс после того как мы перестанем читать их канала
					fmt.Println("This write in chanal")
				case <- cancel:
					fmt.Println("This cancel chan")
					return
			}
		}
	 }()

	 return out
}