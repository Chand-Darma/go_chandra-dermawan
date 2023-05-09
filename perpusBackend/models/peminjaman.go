package models

import "gorm.io/gorm"

type Peminjaman struct {
	gorm.Model
	//ID       int    `json:"id" form:"id"`
	ID_Peminjaman   string `json:"id_peminjaman" form:"id_peminjaman"`
	Tagggal_Pinjam  string `json:"tanggal_pinjam" form:"tanggal_pinjam"`
	Tanggal_Kembali string `json: "tanggal_kembali" form: "tanggal_kembali"`
	Nama_Peminjam   string `json: "nama_peminjam" form: "nama_peminjam"`
	Nomor_Identitas string `json: "nomor_identitas" form: "nomor_identitas"`
	Kode_Buku       string `json: "kode_buku" form: "kode_buku"`
}
