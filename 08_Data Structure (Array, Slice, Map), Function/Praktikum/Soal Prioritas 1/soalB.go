package main

import "fmt"

func Mapping(slice []string) map[string]int {
	counts := make(map[string]int)
	for _, elemen := range slice {
		if _, exist := counts[elemen]; exist {
			counts[elemen]++
		} else {
			counts[elemen] = 1
		}
	}
	return counts
}

func main() {
	fmt.Println(Mapping([]string{"asd", "qwe", "asd", "adi", "qwe", "qwe"}))
	fmt.Println(Mapping([]string{"asd", "qwe", "asd"}))
	fmt.Println(Mapping([]string{}))
}
