package main

import (
	"fmt"
)

func main() {
	// membuat buffered channel dengan kapasitas 10
	c := make(chan int, 10)

	// membuat goroutine untuk mengirimkan bilangan kelipatan 3 ke channel
	go func() {
		for i := 1; i <= 100; i++ {
			if i%3 == 0 {
				c <- i // mengirimkan bilangan ke channel
			}
		}
		close(c) // menutup channel setelah selesai mengirimkan bilangan
	}()

	// mengambil dan mencetak bilangan dari channel
	for n := range c {
		fmt.Println(n)
	}
}
