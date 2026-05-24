package main

import (
	"context"
	"sync"
)

func ParallelMapCtx(ctx context.Context, inputs []int, fn func(int) int, workers int) ([]int, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	if workers < 1 {
		workers = 1
	}

	type task struct {
		index int
		value int
	}

	results := make([]int, len(inputs))
	jobs := make(chan task)
	var workersWG sync.WaitGroup

	for i := 0; i < workers; i++ {
		workersWG.Add(1)
		go func() {
			defer workersWG.Done()
			for {
				select {
				case <-ctx.Done():
					return
				case job, ok := <-jobs:
					if !ok {
						return
					}
					results[job.index] = fn(job.value)
				}
			}
		}()
	}

	go func() {
		defer close(jobs)

		for i, value := range inputs {
			select {
			case <-ctx.Done():
				return
			case jobs <- task{index: i, value: value}:
			}
		}
	}()

	workersDone := make(chan struct{})
	go func() {
		workersWG.Wait()
		close(workersDone)
	}()

	select {
	case <-workersDone:
		if err := ctx.Err(); err != nil {
			return nil, err
		}
		return results, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}