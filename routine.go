package main

import (
	"time"
	"log"
	"sync"
)

func main(){
	start := time.Now()
	var counter = 0 
	var wg sync.WaitGroup

	for i := 1; i <= 10; i++ {
		wg.Add(1) // new thread-->delta argu 
		counter++ 
		go func(thisCounter int){
			log.Println("starting...", thisCounter)
			wg.Done()  // decrement by one, the amount of tasks the waitgroup will need to synchronize at the end, so wg.Add increases it and wg.Done decreases it,
		}(counter)
	}
	wg.Wait()//time.Sleep(2* time.Second)
	elapsed := time.Since(start)
	log.Println(elapsed)
}