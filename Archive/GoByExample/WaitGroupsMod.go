// package ratelimiter
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

type job func() // тип функция

func worker(id int, jobs <-chan job, wg *sync.WaitGroup, cPM *uint64) {
	defer wg.Done()
	for fn := range jobs {
		atomic.AddUint64(cPM, 1)
		//fmt.Println("worker", id, "started de job")
		fn()
		//fmt.Println("worker", id, "finished da job")
	}
	//fmt.Println("Done - worker", id )
}

func main() {

	var wg sync.WaitGroup
	jobs := make(chan job, 4)

	rand.Seed(time.Now().UnixNano() + rand.Int63())
	sliceJobs := make([]job, 3) // Слайс функций для нагрузки
	sliceJobs[0] = func() {     // одну переменный типа job
		//current := time.Now()
		time.Sleep(time.Duration(rand.Intn(1e2)) * time.Millisecond)
		//fmt.Println("Hello", time.Since(current))
	}
	sliceJobs[1] = func() { // вторую
		//current := time.Now()
		time.Sleep(time.Duration(rand.Intn(2e2)) * time.Millisecond)
		//fmt.Println("Golang", time.Since(current))
	}
	sliceJobs[2] = func() { // и еще одну
		//current := time.Now()
		time.Sleep(time.Duration(rand.Intn(3e2)) * time.Millisecond)
		//fmt.Println("World", time.Since(current))
	}

	current := time.Now()
	wg.Add(1)
	go rateLimiter(jobs, 5, 220, &wg)
	for i := 0; i < 1000; i++ {
		jobs <- sliceJobs[rand.Intn(3)]
	} // передаем функции
	close(jobs) // закрытие канала - сигнал для завершения
	wg.Wait()   // ждем когда все "рабочие" закончат
	fmt.Println("Total time: ", time.Since(current))
}

func rateLimiter(jobs chan job, totalWorkers int, ratePerM int, wg *sync.WaitGroup) {
	defer wg.Done()

	var cpM uint64 // counterPerM

	go func() {
		for {
			time.Sleep(time.Second)
			fmt.Println(cpM)
			atomic.CompareAndSwapUint64(&cpM, cpM, 0)
		}
	}() // resetting the counter

	for w := 1; w <= totalWorkers; w++ {
		wg.Add(1)
		go worker(w, jobs, wg, &cpM)
	}
}
