package main

import (
	"fmt"
)

type student struct {
	name       string
	nameEncode string
	score      int
}

type Chiper interface {
	Encode() string
	Decode() string
}

func (s *student) Encode() string {
	var nameEncode string

	for _, c := range s.name {
		if c >= 'a' && c <= 'z' {
			nameEncode += string('a' + ((c - 'a' + 3) % 26))
		} else if c >= 'A' && c <= 'Z' {
			nameEncode += string('A' + ((c - 'A' + 3) % 26))
		} else {
			nameEncode += string(c)
		}
	}

	s.nameEncode = nameEncode

	return nameEncode
}

func (s *student) Decode() string {
	var nameDecode string

	for _, c := range s.nameEncode {
		if c >= 'a' && c <= 'z' {
			nameDecode += string('a' + ((c - 'a' - 3 + 26) % 26))
		} else if c >= 'A' && c <= 'Z' {
			nameDecode += string('A' + ((c - 'A' - 3 + 26) % 26))
		} else {
			nameDecode += string(c)
		}
	}

	s.name = nameDecode

	return nameDecode
}

func main() {
	var menu int
	var a student = student{}
	var c Chiper = &a

	fmt.Print("[1] Encrypt \n[2] Decrypt \nChoose your menu? ")
	fmt.Scan(&menu)

	if menu == 1 {
		fmt.Print("\nInput Student Name: ")
		fmt.Scan(&a.name)
		fmt.Print("\nEncode of student’s name " + a.name + " is : " + c.Encode())
	} else if menu == 2 {
		fmt.Print("\nInput Student Name: ")
		fmt.Scan(&a.nameEncode)
		fmt.Print("\nDecode of student’s name " + a.nameEncode + " is : " + c.Decode())
	}
}
