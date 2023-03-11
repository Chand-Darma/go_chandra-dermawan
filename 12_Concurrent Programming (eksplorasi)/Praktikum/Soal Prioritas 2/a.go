package main

import (
	"fmt"
	"sync"
)

func countLetters(text string, ch chan map[rune]int, wg *sync.WaitGroup) {
	freq := make(map[rune]int)
	for _, char := range text {
		if char >= 'a' && char <= 'z' || char >= 'A' && char <= 'Z' {
			char = rune(toLower(byte(char)))
			if _, ok := freq[char]; ok {
				freq[char]++
			} else {
				freq[char] = 1
			}
		}
	}
	ch <- freq
	wg.Done()
}

func toLower(char byte) byte {
	if char >= 'A' && char <= 'Z' {
		return char + 'a' - 'A'
	}
	return char
}

func main() {
	text := "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua."
	numProcesses := 4

	// Bagi teks menjadi beberapa bagian
	ch := make(chan map[rune]int)
	wg := sync.WaitGroup{}
	for i := 0; i < numProcesses; i++ {
		start := i * len(text) / numProcesses
		end := (i + 1) * len(text) / numProcesses
		wg.Add(1)
		go countLetters(text[start:end], ch, &wg)
	}

	// Gabungkan hasil perhitungan frekuensi huruf dari setiap bagian
	finalFreq := make(map[rune]int)
	for i := 0; i < numProcesses; i++ {
		freq := <-ch
		for char, count := range freq {
			finalFreq[char] += count
		}
	}
	wg.Wait()

	// Cetak frekuensi huruf keseluruhan
	for char, count := range finalFreq {
		fmt.Printf("%c: %d\n", char, count)
	}
}
