package main

import (
	"fmt"
	"runtime"
)

const (
	iterationsNum = 7
	goroutinesNum = 5
)

func doSomeWork(in int) {
	for j := 0; j < iterationsNum; j++ {
		fmt.Printf(formatWork(in, j))
		runtime.Gosched()
	}
}

func main() {
	for i := 0; i < goroutinesNum; i++ {
		go doSomeWork(i)
	}
	fmt.Scanln()
}

func formatWork(in, j int) string {
	return fmt.Sprintf("GoroutinNum: %d iterarions: %d \n", in, j)
}
