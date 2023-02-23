package main

import "fmt"

func main() {
	var sisiAtas, sisiBawah, tinggi float64

	fmt.Println("Program Menghitung Luas Trapesium")
	fmt.Println("---------------------------------")
	fmt.Print("Masukkan sisi atas trapesium: ")
	fmt.Scan(&sisiAtas)
	fmt.Print("Masukkan sisi bawah trapesium: ")
	fmt.Scan(&sisiBawah)
	fmt.Print("Masukkan tinggi trapesium: ")
	fmt.Scan(&tinggi)

	luas := (sisiAtas + sisiBawah) / 2 * tinggi

	fmt.Printf("Luas trapesium adalah: %.2f\n", luas)
}
