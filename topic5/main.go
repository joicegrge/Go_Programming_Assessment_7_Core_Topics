package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"
)

func DownloadURLs(urls []string) {
	var waitGroup sync.WaitGroup
	for _, url := range urls {
		url := url // Capture this iteration's value for the goroutine.
		waitGroup.Add(1)
		go func() {
			defer waitGroup.Done()
			fmt.Println("Starting simulated download:", url)
			time.Sleep(700 * time.Millisecond)
			fmt.Println("Finished simulated download:", url)
		}()
	}
	waitGroup.Wait()
	fmt.Println("All download goroutines have finished.")
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("TOPIC 5: GOROUTINES & WAITGROUP")
	fmt.Print("Enter URL labels separated by commas (example: https://a.com, https://b.com): ")
	line, _ := reader.ReadString('\n')
	parts := strings.Split(strings.TrimSpace(line), ",")
	urls := make([]string, 0, len(parts))
	for _, part := range parts {
		if value := strings.TrimSpace(part); value != "" {
			urls = append(urls, value)
		}
	}
	if len(urls) == 0 {
		fmt.Println("Please enter at least one URL.")
		return
	}
	fmt.Printf("Launching %d goroutines...\n", len(urls))
	DownloadURLs(urls)
	fmt.Println("Note: this is a simulation; it does not make network requests.")
}
