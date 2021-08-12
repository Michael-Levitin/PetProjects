// package ratelimiter
package main

import (
	"fmt"
	"math/rand"
	"time"
)

type job func() // тип функция

// WorkerPool is a contract for Worker Pool implementation
type WorkerPool interface {
	Run()
	AddTask(task func())
}

type workerPool struct {
	maxWorker   int
	queuedTaskC chan job
}

func main() {
	var sleep time.Duration
	ch := make(chan job)             // создаем канал функций
	wait := make(chan time.Duration) // и канал времени
	done := make(chan bool)

	sliceJobs := make([]job, 3)

	// Creating in for loop doesnt work
	//greetings:= []string{"Hello", "World", "And", "Welcome", "To", "Golang"}
	//fmt.Println(greetings)
	//sliceJobs := make([]job, len(greetings))
	//for i:=0; i<len(greetings)-1;i++{
	//	sliceJobs[i] = func() { // Создаем функции
	//		current := time.Now()
	//		time.Sleep(time.Duration(rand.Intn(2e2)) * time.Millisecond)
	//		fmt.Println(greetings[i], time.Since(current))
	//	}
	//}
	//for i:=0; i<len(greetings)-1;i++{
	//	sliceJobs[i]()
	//}

	sliceJobs[0] = func() { // одну переменный типа job
		current := time.Now()
		time.Sleep(time.Duration(rand.Intn(1e2)) * time.Millisecond)
		fmt.Println("Hello", time.Since(current))
	}
	sliceJobs[1] = func() { // вторую
		current := time.Now()
		time.Sleep(time.Duration(rand.Intn(2e2)) * time.Millisecond)
		fmt.Println("Golang", time.Since(current))
	}
	sliceJobs[2] = func() { // и еще одну
		current := time.Now()
		time.Sleep(time.Duration(rand.Intn(3e2)) * time.Millisecond)
		fmt.Println("World", time.Since(current))
	}

	ratelimiter(ch, wait, done, 5, 200) //запускаем функцию которая принимает 2 канала

	current := time.Now()
	rand.Seed(time.Now().UnixNano() + rand.Int63())
	for i := 0; i < 10; i++ {
		select {
		case sleep = <-wait:
			fmt.Println("Sleeping for ", sleep)
			time.Sleep(sleep)
		case ch <- sliceJobs[rand.Intn(3)]:
		}
	}
	done <- true
	fmt.Println("\nRun:", time.Since(current))
}

func ratelimiter(ch chan job, wait chan time.Duration, done chan bool, totalWorker int, ratePerM int) {
	var fn job
	var quit bool
	//waitC := make(chan bool)
	//sleep := time.Second

	// Start Worker Pool.
	wp := workerPool{totalWorker, ch}
	wp.Run()

	go func() {
		for {
			select {
			case fn = <-ch:
				wp.AddTask(fn)
			case quit = <-done:
				_ = quit
				return
			}
		}
	}()
	//<-waitC
}

func (wp *workerPool) Run() {
	for i := 0; i < wp.maxWorker; i++ {
		go func(workerID int) {
			for task := range wp.queuedTaskC {
				task()
			}
		}(i + 1)
	}
}

func (wp *workerPool) AddTask(task func()) {
	wp.queuedTaskC <- task
}
