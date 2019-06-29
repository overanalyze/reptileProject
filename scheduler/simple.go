package scheduler

import "reptileProject/engine"

// 实现Scheduler接口
type SimpleScheduler struct {
	workerChan chan engine.Request
}

func (s *SimpleScheduler) ConfigureMasterWorkerChan(in chan engine.Request) {
	s.workerChan = in
}

func (s *SimpleScheduler) Submit(r engine.Request) {
	go func() { s.workerChan <- r }()
}
