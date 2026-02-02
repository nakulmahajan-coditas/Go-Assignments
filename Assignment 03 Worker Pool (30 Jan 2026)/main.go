// Your task is to create a worker pool with N workers. Jobs should be sent through a shared channel.
// Each worker should pick a job, wait for some time to simulate work, and print which worker processed which job.
// The program should make sure that only N jobs are processed at the same time and should stop after all jobs are finished.

package main

import (
	"fmt"
	"sync"
	"time"
)

func workerPool(workerId uint8, jobsCh <-chan uint8, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobsCh {
		fmt.Printf("worker %d is performing task %d\n", workerId, job)
		time.Sleep(time.Second)
		fmt.Printf("worker %d finished task %d\n", workerId, job)

	}

}

func main() {
	fmt.Println("worker pool starting...")

	fmt.Println("enter number of jobs to be executed: ")
	var totalJobs uint8
	fmt.Scan(&totalJobs)

	fmt.Println("enter number of workers: ")
	var workers uint8
	fmt.Scan(&workers)

	jobsCh := make(chan uint8)
	var wg sync.WaitGroup

	// var w uint8
	for w := uint8(1); w <= workers; w++ {
		wg.Add(1)
		go workerPool(w, jobsCh, &wg)
	}

	// var j uint8
	for j := uint8(1); j <= totalJobs; j++ {
		jobsCh <- j
	}

	close(jobsCh)
	wg.Wait()

	fmt.Println("all jobs completed!")
}

//worker pool helps in res mgmnt by limiting no. of gorouting and reusing them for multiple tasks
//they're useful for cntrling resc usage and improving the efficiency of concurrent task processing
