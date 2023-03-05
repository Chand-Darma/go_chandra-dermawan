package main

import "fmt"

func caesar(offset int, input string) string {
	var result string
	for _, char := range input {
		if char >= 'a' && char <= 'z' {
			// Shift character
			char = 'a' + (char-'a'+rune(offset))%26
		} else if char >= 'A' && char <= 'Z' {
			// Shift character
			char = 'A' + (char-'A'+rune(offset))%26
		}
		// Append to result
		result += string(char)
	}
	return result
}

func main() {
	fmt.Println(caesar(3, "abc"))                           // def
	fmt.Println(caesar(1, "abcdefghijklmnopqrstuvwxyz"))    // bcdefghijklmnopqrstuvwxyza
	fmt.Println(caesar(2, "alta"))                          // cnvc
	fmt.Println(caesar(10, "alterraacademy"))               // kvdobbkkmknowi
	fmt.Println(caesar(1, "abcdefghijklmnopqrstuvwxyz"))    // bcdefghijklmnopqrstuvwxyza
	fmt.Println(caesar(1000, "abcdefghijklmnopqrstuvwxyz")) // mnopqrstuvwxyzabcdefghijkl
}
