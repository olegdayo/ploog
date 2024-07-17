package ploog

import (
	"log/slog"
	"sync"
)

type Task func() error

type Ploog struct {
	tasks <-chan Task
	maxGs uint
}

func New(maxSimultaniousTasks uint) (*Ploog, chan<- Task) {
	tasks := make(chan Task)
	return &Ploog{
		tasks: tasks,
		maxGs: maxSimultaniousTasks,
	}, tasks
}

func execute(task Task, wg *sync.WaitGroup, sema chan struct{}) {
	wg.Add(1)
	defer func() {
		wg.Done()
		<-sema
	}()
	sema <- struct{}{}
	err := task()
	if err != nil {
		slog.Error(err.Error())
	}
}

func (p *Ploog) Start() {
	wg := sync.WaitGroup{}
	sema := make(chan struct{}, p.maxGs)
	for task := range p.tasks {
		go execute(task, &wg, sema)
	}
	wg.Wait()
}
