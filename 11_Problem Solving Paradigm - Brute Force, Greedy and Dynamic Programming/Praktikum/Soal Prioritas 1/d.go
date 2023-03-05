package main

import "fmt"

func SimpleEquations(a, b, c int) {
	// Inisialisasi variabel
	var x, y, z int

	// Loop untuk mencari solusi
	for i := 1; i <= 1000; i++ {
		for j := 1; j <= 1000; j++ {
			for k := 1; k <= 1000; k++ {
				if i+j+k == a && i*j*k == b && i*i+j*j+k*k == c {
					x, y, z = i, j, k
					break
				}
			}
		}
	}

	// Cek apakah solusi ditemukan
	if x == 0 && y == 0 && z == 0 {
		fmt.Println("no solution")
	} else {
		fmt.Println(x, y, z)
	}
}

func main() {
	// Contoh penggunaan program
	SimpleEquations(1, 2, 3)  // no solution
	SimpleEquations(6, 6, 14) // 1 2 3
}
