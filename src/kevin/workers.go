package kevin

import (
	"fmt"
	"math/rand"
	"time"
)

type Jobs struct {
	jobId int
	done  chan bool
}

func worker(id int, jobs <-chan Jobs) {
	for j := range jobs {
		fmt.Println("worker : ", id, " =====> start job", j.jobId)
		s := rand.Intn(10)
		for t := 0; t < s; t++ {
			time.Sleep(time.Millisecond * 10)
		}
		fmt.Println("worker ", id, "finished job ", j.jobId)
		j.done <- true
	}
}

func checkJob(job Jobs, done chan<- bool) {
	for {
		if <-job.done {
			done <- true
		}
	}
}

func mainWorkers() {
	starttime := time.Now()
	jobs := make(chan Jobs)
	done := make(chan bool)
	allDone := make(chan bool)
	for w := 0; w < 10; w++ {
		//done <- false
		go worker(w, jobs)
	}

	for j := 0; j < 100; j++ {
		job := Jobs{j, done}
		jobs <- job
		go checkJob(job, allDone)
	}

	//close(jobs)

	for {
		if <-allDone {
			break
		}
	}

	fmt.Println(starttime)
	fmt.Println(time.Now())
	time.Sleep(time.Second * 2)
}
