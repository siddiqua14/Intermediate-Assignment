package main

import (
	"log"
	"time"
)


func Worker(id int, done chan bool) {
	log.Printf("Worker %d: Starting", id) 
	time.Sleep(2 * time.Second) 
	log.Printf("Worker %d: Finished", id) 
	done <- true 
}

func main() {
	log.Println("Starting the main program ")
	numWorkers := 5
	done := make(chan bool, numWorkers) 

	for i := 1; i <= numWorkers; i++ {
		go Worker(i, done)
	}

	log.Printf("All %d workers have started. Waiting for them to complete", numWorkers)

	for i := 1; i <= numWorkers; i++ {
		<-done 
		log.Printf("Worker %d completed signal received.", i)
	}

	log.Println("All workers have completed their tasks.")
	log.Println("Exiting the main program.")
}
