package main

import (
	"fmt"
	"strings"
)

func main() {
	var input1, input2 string
	var option int

	//menampilkan kalimat
	fmt.Println("Apakah sebuah Palindrome?")
	//menampilkan inputan jumlah
	fmt.Print("Pilih jumlah inputan (1 atau 2): ")
	fmt.Scan(&option)

	if option == 1 {
		fmt.Print("Masukkan sebuah kata: ")
		fmt.Scan(&input1)

		if isPalindrome(input1) {
			fmt.Printf("captured: %s adalah sebuah palindrome.\n", input1)
		} else {
			fmt.Printf("captured: %s bukan sebuah palindrome.\n", input1)
		}
	} else if option == 2 {
		fmt.Print("Masukkan dua kata dipisahkan dengan spasi: ")
		fmt.Scan(&input1, &input2)

		input1 = strings.ToLower(input1)
		input2 = strings.ToLower(input2)

		if isPalindrome(input1) && isPalindrome(input2) {
			fmt.Printf("captured: %s dan %s keduanya palindrome.\n", input1, input2)
		} else if isPalindrome(input1) {
			fmt.Printf("captured: %s adalah sebuah palindrome, tetapi %s bukan.\n", input1, input2)
		} else if isPalindrome(input2) {
			fmt.Printf("captured: %s bukan sebuah palindrome, tetapi %s adalah.\n", input1, input2)
		} else {
			fmt.Printf("captured: %s dan %s keduanya bukan palindrome.\n", input1, input2)
		}
	} else {
		fmt.Println("Pilihan tidak valid.")
	}
}

func isPalindrome(input string) bool {
	length := len(input)

	for i := 0; i < length/2; i++ {
		if input[i] != input[length-1-i] {
			return false
		}
	}

	return true
}
