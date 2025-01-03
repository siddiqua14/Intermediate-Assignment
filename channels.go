package main

import (
	"time"
	"log"
)

func main(){
	counterChan := make (chan int)

	go func(){
		for i := 1; i <= 8; i++ {
			counterChan <-1
			time.Sleep(100 * time.Millisecond)
		}
	}()

	go func(){
		for i := 1; i <= 8; i++ {
			counterChan <-1
			time.Sleep(100 * time.Millisecond)
		}
	}()
	{
		var counter = 0
		for counter < 16 {
			counter += <-counterChan
			log.Println(counter)
		}
	}
}