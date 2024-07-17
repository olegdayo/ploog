package main

import (
	"log/slog"
	"sync"
)

type Task func() error

type Pool struct {
	tasks chan Task
	maxGs uint
}

func New(maxSimultaniousTasks uint) *Pool {
	return &Pool{
		tasks: make(chan Task),
		maxGs: maxSimultaniousTasks,
	}
}

func (p *Pool) AddTasks(tasks ...Task) {
	for _, task := range tasks {
		p.tasks <- task
	}
}

func (p *Pool) FinishInput() {
	close(p.tasks)
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

func (p *Pool) Start() {
	wg := sync.WaitGroup{}
	sema := make(chan struct{}, p.maxGs)
	for task := range p.tasks {
		go execute(task, &wg, sema)
	}
	wg.Wait()
}
