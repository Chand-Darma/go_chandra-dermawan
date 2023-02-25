package main

import (
	"fmt"
	"math"
)

func primeNumbering(number int) []bool {
	// Buat array yang berisi n+1 elemen dengan nilai awal true
	isPrime := make([]bool, number+1)
	for i := 2; i <= number; i++ {
		isPrime[i] = true
	}

	// Lakukan eliminasi pada array
	for i := 2; i <= int(math.Sqrt(float64(number))); i++ {
		if isPrime[i] {
			for j := i * i; j <= number; j += i {
				isPrime[j] = false
			}
		}
	}

	return isPrime
}

func primeNumber(number int) string {
	if number <= 1 {
		return "Bilangan tidak prima"
	}

	// Cek apakah n adalah bilangan prima menggunakan array yang dihasilkan dari sieveOfEratosthenes
	primes := primeNumbering(number)
	if primes[number] {
		return "Bilangan prima"
	} else {
		return "Bilangan tidak prima"
	}
}

func main() {
	fmt.Println(primeNumber(1000000007)) // true
	fmt.Println(primeNumber(1500450271)) // true
	fmt.Println(primeNumber(13))         // true
	fmt.Println(primeNumber(17))         // true
	fmt.Println(primeNumber(20))         // false
	fmt.Println(primeNumber(35))         // false
}
