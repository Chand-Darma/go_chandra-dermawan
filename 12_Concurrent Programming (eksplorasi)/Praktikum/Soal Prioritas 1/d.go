package main

import (
	"fmt"
	"sync"
)

func factorial(n int, wg *sync.WaitGroup, mu *sync.Mutex, res *int) {
	defer wg.Done()

	fact := 1
	for i := 2; i <= n; i++ {
		fact *= i
	}

	mu.Lock()
	*res *= fact
	mu.Unlock()
}

func main() {
	n := 29 //
	res := 1

	var wg sync.WaitGroup
	var mu sync.Mutex

	for i := n; i > 0; i-- {
		wg.Add(1)
		go factorial(i, &wg, &mu, &res)
	}

	wg.Wait()

	fmt.Printf("%d! = %d\n", n, res)
}
