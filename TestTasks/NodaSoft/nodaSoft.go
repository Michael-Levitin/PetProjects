//======================================================================================
// nodaSoft - Задание
package main

import (
	"fmt"
	"time"
)

const (
	timeRun  = 3000  // total run in milliseconds
	coolDown = 150   // task processing "cool down" in milliseconds > 0
	timeout  = 20000 // task timeout in milliseconds
)

// ЗАДАНИЕ:
// * сделать из плохого кода хороший;
// * важно сохранить логику появления ошибочных тасков;
// * сделать правильную мультипоточность обработки заданий.
// Обновленный код отправить через merge-request.

// приложение эмулирует получение и обработку тасков, пытается и получать и обрабатывать в многопоточном режиме
// В конце должно выводить успешные таски и ошибки выполнены остальных тасков

// A tType represents a meaninglessness of our life
type tType struct {
	id         int64
	startTime  string // время создания
	endTime    string // время завершения
	taskError  error
	taskResult []byte
}

func main() {

	taskCreater := func(a chan tType) {
		var err error
		var startTime string
		var timeTemp time.Time
		for {
			err = nil
			//timeTemp = time.Now().UnixNano()
			timeTemp = time.Now()
			startTime = timeTemp.Format(time.RFC3339Nano)
			if timeTemp.UnixNano()%2 > 0 { // вот такое условие появления ошибочных тасков
				err = fmt.Errorf("task creation failed")
			}
			a <- tType{id: timeTemp.UnixNano(), startTime: startTime, taskError: err} // передаем таск на выполнение
		}
	}

	taskWorker := func(a tType) tType {
		if a.taskError != nil {
			a.taskResult = []byte("task is erroneous")
		} else {
			tt, _ := time.Parse(time.RFC3339Nano, a.startTime)

			if tt.After(time.Now().Add(-time.Millisecond * timeout)) {
				a.taskResult = []byte("task has been completed successfully")
			} else {
				a.taskError = fmt.Errorf("timeout failure")
				a.taskResult = []byte("task failed to execute")
			}
		}

		time.Sleep(time.Millisecond * coolDown)
		a.endTime = time.Now().Format(time.RFC3339Nano)
		return a
	}

	taskSorter := func(t tType, doneTasks, undoneTasks chan tType) {
		if t.taskError == nil {
			doneTasks <- t
		} else {
			undoneTasks <- t
		}
	}

	superChan := make(chan tType, 10)
	defer close(superChan)
	go taskCreater(superChan)

	doneTasks := make(chan tType)
	undoneTasks := make(chan tType)
	defer close(doneTasks)
	defer close(undoneTasks)
	go func() {
		for t := range superChan {
			t = taskWorker(t)
			taskSorter(t, doneTasks, undoneTasks)
		}
	}()

	var result []int64
	go func() {
		for r := range doneTasks {
			result = append(result, r.id)
		}
	}()

	var err []string
	go func() {
		for r := range undoneTasks {
			err = append(err, fmt.Sprintf("Task id: %d, start time: %s, error: %s\n", r.id, r.startTime, r.taskError))
		}
	}()

	time.Sleep(time.Millisecond * timeRun)

	fmt.Println("Finished processing tasks")
	fmt.Printf("Summary - Done tasks: %v, Errors: %v\n", len(result), len(err))
	fmt.Printf("Done tasks:\n %v\n\n", result)
	fmt.Printf("Errors:\n %v\n", err)
}
