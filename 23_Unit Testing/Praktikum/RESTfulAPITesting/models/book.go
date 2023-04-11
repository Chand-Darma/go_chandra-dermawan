package models

import "gorm.io/gorm"

type Books struct {
	gorm.Model
	Judul    string `json:"judul" form:"judul"`
	Penulis  string `json:"penulis" form:"penulis"`
	Penerbit string `json:"penerbit" form:"penerbit"`
	Token    string `json: "token" form: "token"`
}
