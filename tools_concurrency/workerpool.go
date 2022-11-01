package main

import (
	"fmt"
	"runtime"
	"time"
)

const goroutinesNum = 3

func startWorker(workerNum int, in <-chan string) {
	for input := range in {
		fmt.Println("workNum: ", workerNum, " input: ", input)
		runtime.Gosched()
	}
	fmt.Println("Finished: ", workerNum)
}

func main() {
	workerInput := make(chan string, 2)
	for i := 0; i < goroutinesNum; i++ {
		go startWorker(i, workerInput)
	}

	months := []string{"Ян", "фе", "Ма", "Ап", "Май", "Июнь", "Июль", "Аф", "Се", "Ок", "Ноя", "Декабрь"}

	for _, monthName := range months {
		workerInput <- monthName
	}

	close(workerInput)

	time.Sleep(time.Microsecond)

}
