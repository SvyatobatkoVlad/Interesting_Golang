package main

import "fmt"

func main() {

	m := make(chan string, 3)
	cnt := 5

	for i := 0; i < cnt; i++ {
		go func() {
			m <- fmt.Sprintf("Goroutine %d", i)
		}()
	}

	for i := 0; i < cnt; i++ {
		ReceiveFromCh(m)
	}

}
func ReceiveFromCh(ch chan string) {
	fmt.Println(<-ch)
}

// output 5 5 5 5 5
