package main

// Job is a struct the a worker can process
type Job struct {
	ID      int
	message string
}

var workers chan chan Job

// Jobs receives job objects to be handled by the workers
var Jobs = make(chan Job, 100)

// CreateWorkers instantiate n workers and returns a common
// output chan, and input chan and a function to shut the workers
// down.
func CreateWorkers(n int) (chan string, func()) {
	out := make(chan string)
	workers = make(chan chan Job, n)
	var quits []chan bool
	for i := 0; i < n; i++ {
		w := newWorker(i, out)
		w.start()
		quits = append(quits, w.Quit)
	}
	go accept()
	return out, func() {
		for i := 0; i < n; i++ {
			quits[i] <- true
		}
	}
}

func accept() {
	for {
		select {
		case job := <-Jobs:
			go assign(job)
		}
	}
}

func assign(job Job) {
	worker := <-workers
	worker <- job
}
