package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Print("Masukkan sebuah angka: ")
	fmt.Scan(&n)

	factors := []int{}
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			factors = append(factors, i)
		}
	}

	sort.Ints(factors)

	//menampilkan hasil faktor sebuah angka
	fmt.Printf("Faktor bilangan dari %d adalah:\n", n)
	for _, factor := range factors {
		fmt.Println(factor)
	}
}
