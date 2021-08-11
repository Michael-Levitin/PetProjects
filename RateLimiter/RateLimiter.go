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

	pingpong(ch, wait, done) //запускаем функцию которая принимает 2 канала

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
