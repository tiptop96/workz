package main

import (
	"fmt"
)

type worker struct {
	ID   int
	Quit chan bool
	Jobs chan Job
	Out  chan string
}

func newWorker(id int, out chan string) *worker {
	return &worker{
		ID:   id,
		Quit: make(chan bool),
		Jobs: make(chan Job),
		Out:  out}
}

func (w *worker) start() {
	go func() {
		for {
			workers <- w.Jobs

			select {
			case work := <-w.Jobs:
				w.Out <- fmt.Sprintf("Echo from worker [%d], Job [%d]: %s", w.ID, work.ID, work.message)
			case <-w.Quit:
				return
			}
		}
	}()
}
