package main

import (
	"fmt"
	"math"
)

func main() {
	// contoh matriks 3x3
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{9, 8, 9},
	}

	// inisialisasi variabel untuk menyimpan jumlah diagonal
	sum1 := 0
	sum2 := 0

	// hitung jumlah diagonal dari kiri ke kanan
	for i := 0; i < len(matrix); i++ {
		sum1 += matrix[i][i]
	}

	// hitung jumlah diagonal dari kanan ke kiri
	for i := 0; i < len(matrix); i++ {
		sum2 += matrix[i][len(matrix)-i-1]
	}

	// hitung selisih absolut antara jumlah diagonal
	diff := int(math.Abs(float64(sum1 - sum2)))

	// tampilkan hasil
	fmt.Printf("Matriks:\n%v\nPerbedaan mutlak adalah : %d\n", matrix, diff)
}
