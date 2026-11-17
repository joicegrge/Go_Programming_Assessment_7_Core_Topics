package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

type SafeCounter struct {
	mutex sync.Mutex
	value int
}

func (counter *SafeCounter) Increment() {
	counter.mutex.Lock()
	defer counter.mutex.Unlock()
	counter.value++
}

func (counter *SafeCounter) Value() int {
	counter.mutex.Lock()
	defer counter.mutex.Unlock()
	return counter.value
}

func RunConcurrentIncrements(counter *SafeCounter, workers, incrementsPerWorker int) {
	var waitGroup sync.WaitGroup
	waitGroup.Add(workers)
	for worker := 0; worker < workers; worker++ {
		go func() {
			defer waitGroup.Done()
			for i := 0; i < incrementsPerWorker; i++ {
				counter.Increment()
			}
		}()
	}
	waitGroup.Wait()
}

func readPositiveInt(reader *bufio.Reader, prompt string) int {
	for {
		fmt.Print(prompt)
		line, _ := reader.ReadString('\n')
		value, err := strconv.Atoi(strings.TrimSpace(line))
		if err == nil && value > 0 {
			return value
		}
		fmt.Println("Enter a whole number greater than zero.")
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("TOPIC 6: MUTEX & RACE CONDITIONS")
	workers := readPositiveInt(reader, "Enter number of goroutines/workers: ")
	increments := readPositiveInt(reader, "Enter increments per goroutine: ")

	counter := &SafeCounter{}
	RunConcurrentIncrements(counter, workers, increments)
	expected := workers * increments
	fmt.Printf("Workers: %d\nIncrements per worker: %d\n", workers, increments)
	fmt.Printf("Expected count: %d\nActual count:   %d\n", expected, counter.Value())
}
