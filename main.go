package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(url string, wg *sync.WaitGroup, resultChan chan string) {
	defer wg.Done()
	time.Sleep(time.Millisecond * 50)
	fmt.Println("Worker for URL:", url)
	resultChan <- url
}

func main() {
	var wg sync.WaitGroup
	startTime := time.Now()
	resultChan := make(chan string, 2)
	wg.Add(2)
	go worker("https://example.com", &wg, resultChan)
	go worker("https://example.com", &wg, resultChan)
	wg.Wait()

	close(resultChan)

	for result := range resultChan {
		fmt.Println("result ", result)
	}

	fmt.Println("it took", time.Since(startTime), "ms.")
}
