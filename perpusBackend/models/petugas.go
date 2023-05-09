package models

import "gorm.io/gorm"

type Petugas struct {
	gorm.Model
	Username   string `json:"username" form:"username"`
	Password   string `json:"password" form:"password"`
	Nama       string `json:"nama" form:"nama"`
	Level      string `json: "level" form: "level"`
	No_petugas string `json: "no_petugas" form: "no_petugas"`
}
