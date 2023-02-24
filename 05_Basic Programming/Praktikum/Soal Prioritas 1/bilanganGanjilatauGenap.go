package main

import "fmt"

func main() {
	var bilangan int

	///inputan bilangan
	fmt.Print("Masukkan bilangan: ")

	fmt.Scanln(&bilangan)

	if bilangan%2 == 0 {
		//menampilkan bilangan genap
		fmt.Printf("%d adalah bilangan genap", bilangan)
	} else {
		//menampilan bilangan ganjil
		fmt.Printf("%d adalah bilangan ganjil", bilangan)
	}
}
