package main

import "fmt"

func main() {

	//sebuah variabel
	var n int
	// menampilkan teks untuk input angka
	fmt.Print("Masukkan angka : ")
	fmt.Scanln(&n)

	//looping untuk menghasilkan data
	for i := 1; i <= n; i++ {
		for j := 1; j <= n-i; j++ {
			fmt.Print(" ")
		}
		for k := 1; k <= i; k++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}
