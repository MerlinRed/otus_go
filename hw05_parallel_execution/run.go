package hw05parallelexecution

import (
	"errors"
	"sync"
	"sync/atomic"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

func Run(tasks []Task, n, m int) error {
	wg := sync.WaitGroup{}

	errorCount := int32(0)
	tasksChannel := make(chan Task, len(tasks))

	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for task := range tasksChannel {
				if int32(m) < atomic.LoadInt32(&errorCount) {
					return
				}

				if task() != nil {
					atomic.AddInt32(&errorCount, 1)
				}
			}
		}()
	}

	for _, task := range tasks {
		tasksChannel <- task
	}

	close(tasksChannel)
	wg.Wait()

	if int32(m) < atomic.LoadInt32(&errorCount) {
		return ErrErrorsLimitExceeded
	}

	return nil
}
