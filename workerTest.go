package main

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"strconv"
	"sync"
	"time"
)

type Post struct {
	PostId int
	Id     int
	Name   string
	Email  string
	Body   string
}

func worker(id int, jobs <-chan int, results chan<- []Post, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := range jobs {

		output := make([]Post, 0)
		script := exec.Command("python", "query.py", strconv.Itoa(i))

		stdout, err := script.StdoutPipe()
		if err != nil {
			panic(err)
		}

		if err := script.Start(); err != nil {
			panic(err)
		}

		if err := json.NewDecoder(stdout).Decode(&output); err != nil {
			panic(err)
		}

		if err := script.Wait(); err != nil {
			panic(err)
		}

		results <- output
	}
}

func main() {
	totalJobs := 500
	totalWorkers := 150

	start := time.Now()
	jobs := make(chan int, totalJobs)
	results := make(chan []Post, totalJobs)
	var wg sync.WaitGroup

	for i := 1; i <= totalJobs; i++ {
		jobs <- i
	}
	close(jobs)

	for i := 0; i < totalWorkers; i++ {
		wg.Add(1)
		go worker(i, jobs, results, &wg)
	}

	wg.Wait()
	close(results)

	finish := time.Since(start)
	fmt.Println("Finished at ", finish)

	for result := range results {
		for _, p := range result {
			fmt.Printf("Id: %v, Email: %v", p.Id, p.Email)
		}
		fmt.Println("\n")
	}

}
