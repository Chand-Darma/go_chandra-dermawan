package models

import "gorm.io/gorm"

type Users struct {
	gorm.Model

	Username string `json:"username" form:"username"`
	Password string `json:"username" form:"username"`
	Token    string `json: "token" form: "token"`
}
