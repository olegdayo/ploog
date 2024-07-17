package main

import (
	"fmt"
	"time"

	"github.com/olegdayo/ploog"
)

func main() {
	task := func() error {
		time.Sleep(2 * time.Second)
		fmt.Println("bruh")
		return nil
	}

	p, tasksCh := ploog.New(5)

	go func() {
		for i := 0; i < 20; i++ {
			tasksCh <- task
		}
		close(tasksCh)
	}()

	p.Start()
}
