package main

import (
	"fmt"
	"math"
)

type Student struct {
	Name  string
	Score []int
}

func main() {
	// inisialisasi data para siswa
	students := []Student{
		{"Rizky", []int{80}},
		{"Kobar", []int{75}},
		{"Ismail", []int{70}},
		{"Umam", []int{75}},
		{"Yopan", []int{60}},
	}

	// hitung rata-rata skor para siswa
	var total int
	for _, student := range students {
		for _, score := range student.Score {
			total += score
		}
	}
	avg := math.Round(float64(total) / float64(len(students)))

	// cari siswa dengan skor minimum dan maksimum
	var minScore int = math.MaxInt32
	var maxScore int = math.MinInt32
	var minName string
	var maxName string
	for _, student := range students {
		for _, score := range student.Score {
			if score < minScore {
				minScore = score
				minName = student.Name
			}
			if score > maxScore {
				maxScore = score
				maxName = student.Name
			}
		}
	}

	// tampilkan hasil
	fmt.Printf("Average score: %v\n", int(avg))
	fmt.Printf("Min score of students: %v (%v)\n", minName, minScore)
	fmt.Printf("Max score of students: %v (%v)\n", maxName, maxScore)
}
