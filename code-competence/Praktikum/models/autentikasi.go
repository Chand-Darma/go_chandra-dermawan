package models

import "gorm.io/gorm"

type Users struct {
	gorm.Model
	//ID       int    `json:"id" form:"id"`
	Username string `json:"username" form:"username"`
	Email    string `json:"email" form:"email"`
	Password string `json: "password" form: "password"`
	Token    string `json: "token" form: "token"`
}
