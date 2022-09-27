package main

import (
	"fmt"
	"math/rand"
	RL "petProjects/SolutionRL/RateLimiter"
	"sync"
	"time"
)

func main() {

	var wg sync.WaitGroup
	jobs := make(chan RL.Job, 4)

	rand.Seed(time.Now().UnixNano() + rand.Int63())
	sliceJobs := make([]RL.Job, 3) // Слайс функций для нагрузки
	sliceJobs[0] = func() {        // одну переменный типа job
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

	current := time.Now()
	wg.Add(1)
	go RL.RateLimiter(jobs, 5, 3000, &wg)
	for i := 0; i < 1000; i++ {
		jobs <- sliceJobs[rand.Intn(3)]
	} // передаем функции
	close(jobs) // закрытие канала - сигнал для завершения
	wg.Wait()   // ждем когда все "рабочие" закончат
	fmt.Println("Total time: ", time.Since(current))
}
