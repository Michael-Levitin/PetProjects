// package ratelimiter
package main

import (
	"fmt"
	"math/rand"
	"time"
)

type job func() // тип функция

func main() {
	var sleep time.Duration

	ch := make(chan job)             // создаем канал функций
	wait := make(chan time.Duration) // и канал времени
	done := make(chan bool)

	sliceJobs := make([]job, 2)
	sliceJobs[0] = func() { // одну переменный типа job
		current := time.Now()
		time.Sleep(time.Duration(rand.Intn(2e2)) * time.Millisecond)
		fmt.Println("Hello", time.Since(current))
		//fmt.Println("Hello")
	}
	sliceJobs[1] = func() { // и вторую
		current := time.Now()
		time.Sleep(time.Duration(rand.Intn(2e2)) * time.Millisecond)
		fmt.Println("world", time.Since(current))
		//fmt.Println("world")
	}

	go pingpong(ch, wait, done) //запускаем функцию которая принимает 2 канала

	for i := 0; i < 10; i++ {
		select {
		case sleep = <-wait:
			fmt.Println("Sleeping for ", sleep)
			time.Sleep(sleep)
		case ch <- sliceJobs[rand.Intn(2)]:
		}

		//if sleep = <-wait; sleep != 0{
		//	fmt.Println("Sleeping for ", sleep)
		//	time.Sleep(sleep)
		//	sleep = 0
		//} else {
		//	ch <- sliceJobs[rand.Intn(2)]
		//}

		//case
		//	ch <- sliceJobs[rand.Intn(2)]
		//}
		//fmt.Println("Iteration", i)
		//duration := time.Duration(rand.Intn(1e3)) * time.Millisecond
		//fmt.Println("Sleeping for", duration )
		//time.Sleep(duration)

	}
	done <- true
}

//quantity int, ratePerM int
func pingpong(ch <-chan job, wait chan time.Duration, done chan bool) {
	var fn job
	var quit bool
	sleep := time.Second
	//var counter int
	go func() {
		for i := 0; ; i++ {
			if i%5 == 0 {
				fmt.Println("Sleeping")
				wait <- sleep
				time.Sleep(sleep)
			}
			select {
			case fn = <-ch:
				fn()
			case quit = <-done:
				_ = quit
				return
			}
		}
	}()
}
