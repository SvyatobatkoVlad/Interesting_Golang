package main

import (
	"fmt"
	"time"
)

func getComments() chan string {
	result := make(chan string, 1)
	go func(out chan<- string) {
		time.Sleep(2 * time.Second)
		fmt.Println("async operation ready, return comments")
		out <- "32 comment"
	}(result)
	return result
}

func getPage() {
	resultCh := getComments()

	time.Sleep(1 * time.Second)
	fmt.Println("get related articles")

	commentsData := <-resultCh
	fmt.Println("Main goroutine: ", commentsData)
}

func main() {
	getPage()
	time.Sleep(5 * time.Second)
}
