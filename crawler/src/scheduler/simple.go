package scheduler

import (
	"engine"
)

type SimpleScheduler struct {
	workerChan chan engine.Request
}

func (s *SimpleScheduler) WorkerChan() chan engine.Request {
	return s.workerChan
}

func (s *SimpleScheduler) ConfigureMasterWorkerChan(c chan engine.Request) {
	s.workerChan = c
}

func (s *SimpleScheduler) Submit(r engine.Request) {
	go func() { s.workerChan <- r }()
}

func (s *SimpleScheduler) workerReady(r engine.Request) {
	panic("aaaaa")
}

func (s *SimpleScheduler) Run() {
	s.workerChan = make(chan engine.Request)
}
