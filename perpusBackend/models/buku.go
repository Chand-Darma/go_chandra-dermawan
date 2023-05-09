package models

import "gorm.io/gorm"

type Buku struct {
	gorm.Model
	//ID       int    `json:"id" form:"id"`
	Pengarang string `json:"pengarang" form:"pengarang"`
	Judul     string `json:"judul" form:"judul"`
	Kategori  string `json:"kategori" form:"kategori"`
	Penerbit  string `json:"penerbit" form:"penerbit"`
	Tahun     string `json:"tahun" form:"tahun"`
	ID_Buku   string `json:"id_buku" form:"id_buku"`
	Stok      string `json:"stok" form:"stok"`
}
