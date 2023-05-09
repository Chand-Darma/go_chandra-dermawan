package models

import "gorm.io/gorm"

type Pengembalian struct {
	gorm.Model
	//ID       int    `json:"id" form:"id"`
	Nama_Pengembalian string `json:"nama_pengembalian" form:"nama_pengembalian"`
	Tagggal_Pinjam    string `json:"tanggal_pinjam" form:"tanggal_pinjam"`
	Tanggal_Kembali   string `json: "tanggal_kembali" form: "tanggal_kembali"`
	Nomor_Identitas   string `json: "nomor_identitas" form: "nomor_identitas"`
	Kode_Buku         string `json: "kode_buku" form: "kode_buku"`
	Kondisi_Buku      string `json: "kondisi_buku" form: "kondisi_buku"`
}
