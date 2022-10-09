package ratelimiter

import (
	RL "RateLimiter"
	"sync"
	"testing"
	"time"
)

func TestSingleW(t *testing.T) {
	// One worker 3 functions, 1 sec each
	finished := time.Millisecond * 3005 // 5 msec spare
	var wg sync.WaitGroup
	jobs := make(chan RL.Job, 4)

	sliceJobs := func() { // одну переменный типа job
		time.Sleep(time.Second)
	}

	current := time.Now()
	wg.Add(1)
	go RL.RateLimiter(jobs, 1, 3000, &wg)
	for i := 0; i < 3; i++ {
		jobs <- sliceJobs
	} // передаем функции
	close(jobs) // закрытие канала - сигнал для завершения
	wg.Wait()   // ждем когда все "рабочие" закончат
	total := time.Since(current)

	if total > finished {
		t.Errorf("%d != %d", total, finished)
	}
}

func TestWorkersEqFunc(t *testing.T) {
	// Same number of workers and functions
	finished := time.Millisecond * 1005 // 5 msec spare
	var wg sync.WaitGroup
	jobs := make(chan RL.Job, 4)

	sliceJobs := func() { // одну переменный типа job
		time.Sleep(time.Second)
	}

	current := time.Now()
	wg.Add(1)
	go RL.RateLimiter(jobs, 50, 3000, &wg)
	for i := 0; i < 50; i++ {
		jobs <- sliceJobs
	} // передаем функции
	close(jobs) // закрытие канала - сигнал для завершения
	wg.Wait()   // ждем когда все "рабочие" закончат
	total := time.Since(current)

	if total > finished {
		t.Errorf("%d != %d", total, finished)
	}
}

func TestWPool10W20F(t *testing.T) {
	// 10 workers and 20 functions
	finished := time.Millisecond * 2005 // 5 msec spare
	var wg sync.WaitGroup
	jobs := make(chan RL.Job, 4)

	sliceJobs := func() { // одну переменный типа job
		time.Sleep(time.Second)
	}

	current := time.Now()
	wg.Add(1)
	go RL.RateLimiter(jobs, 10, 3000, &wg)
	for i := 0; i < 20; i++ {
		jobs <- sliceJobs
	} // передаем функции
	close(jobs) // закрытие канала - сигнал для завершения
	wg.Wait()   // ждем когда все "рабочие" закончат
	total := time.Since(current)

	if total > finished {
		t.Errorf("%d != %d", total, finished)
	}
}
