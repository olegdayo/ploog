package main

import (
	"fmt"
	"time"
)

func main() {
	task := func() error {
		time.Sleep(2 * time.Second)
		fmt.Println("bruh")
		return nil
	}

	p := New(5)

	go func() {
		for i := 0; i < 20; i++ {
			p.AddTasks(task)
		}
		p.FinishInput()
	}()

	p.Start()
}
