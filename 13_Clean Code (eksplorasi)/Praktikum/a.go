package main

import "fmt"

// type user haruslah menggunakan huruf kapital untuk bisa diakses oleh package lain
type User struct {
	ID       int
	Username string // tipe data string lebih tepat untuk username
	Password string // tipe data string lebih tepat untuk password
}

// gunakan nama yang lebih deskriptif dan bermakna
type UserService struct {
	users []User // ubah variabel t menjadi users dan ubah tipe data menjadi User
}

// method getallusers tidak perlu menggunakan receiver u, cukup dengan menggunakan variabel us
func (us *UserService) GetAllUsers() []User {
	return us.users
}

func (us *UserService) GetUserByID(id int) (User, bool) { // tambahkan return boolean untuk menandakan apakah user ditemukan atau tidak
	for _, u := range us.users {
		if u.ID == id {
			return u, true
		}
	}
	// tambahkan nilai false untuk menandakan user tidak ditemukan
	return User{}, false
}

func main() {
	// buat variabel userservice dengan menggunakan struct UserService dan inisialisasi dengan slice kosong
	userService := UserService{users: []User{}}

	// tambahkan user baru
	userService.users = append(userService.users, User{ID: 1, Username: "john_doe", Password: "password123"})

	// panggil method GetAllUsers dan GetUserByID
	users := userService.GetAllUsers()
	user, found := userService.GetUserByID(1)

	// output hasil cetak
	fmt.Println(users)
	fmt.Println(user, found)
}
