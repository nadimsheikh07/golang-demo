package main

import (
	"fmt"
	"sync"
	"time"
)

var images = []string{
	"https://example.com/image1.jpg",
	"https://example.com/image2.jpg",
	"https://example.com/image3.jpg",
	"https://example.com/image4.jpg",
	"https://example.com/image5.jpg",
	"https://example.com/image6.jpg",
	"https://example.com/image7.jpg",
	"https://example.com/image8.jpg",
	"https://example.com/image9.jpg",
	"https://example.com/image10.jpg",
}

func worker(jobsChan chan string, wg *sync.WaitGroup, resultChan chan string) {
	defer wg.Done()

	for job := range jobsChan {
		time.Sleep(time.Millisecond * 50)
		resultChan <- job
	}
}

func main() {
	var wg sync.WaitGroup
	startTime := time.Now()
	resultChan := make(chan string, len(images))
	jobsChan := make(chan string, len(images))

	var totalWorkers = 5

	for i := 1; i < totalWorkers; i++ {
		wg.Add(1)
		go worker(jobsChan, &wg, resultChan)

	}

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	for _, image := range images {
		jobsChan <- image
	}

	close(jobsChan)

	for result := range resultChan {
		fmt.Println("job completed: ", result)
	}

	fmt.Println("it took", time.Since(startTime), "ms.")
}
