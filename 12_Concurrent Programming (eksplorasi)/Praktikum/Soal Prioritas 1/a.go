package main

import (
	"fmt"
	"time"
)

func printMultiples(x int) {
	for i := 1; ; i++ {
		if i%x == 0 {
			fmt.Printf("%d ", i)
		}
		time.Sleep(3 * time.Second)
	}
}

func main() {

	//goroutine
	go printMultiples(3)
	var input string
	fmt.Scanln(&input) // berhenti ketika ada input dari keyboard

}
