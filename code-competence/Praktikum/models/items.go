package models

import "gorm.io/gorm"

type Items struct {
	gorm.Model
	//ID       int    `json:"id" form:"id"`
	Name        string `json:"name" form:"name"`
	Description string `json:"description" form:"description"`
	Stock       string `json:"stock" form:"stock"`
	Price       string `json:"price" form:"price"`
	Category_Id string `json:"category_id" form:"category_id"`
	Token       string `json: "token" form: "token"`
}
