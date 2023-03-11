package main

import "fmt"

type Kendaraan struct {
	TotalRoda       int
	KecepatanPerJam int
}

type Mobil struct {
	Kendaraan
}

func (m *Mobil) Berjalan() {
	m.TambahKecepatan(10)
}

func (m *Mobil) TambahKecepatan(kecepatanBaru int) {
	m.KecepatanPerJam += kecepatanBaru
}

func main() {
	mobilCepat := Mobil{}
	mobilCepat.Berjalan()
	mobilCepat.Berjalan()
	mobilCepat.Berjalan()

	mobilLamban := Mobil{}
	mobilLamban.Berjalan()

	fmt.Println("Kecepatan mobil cepat:", mobilCepat.KecepatanPerJam)
	fmt.Println("Kecepatan mobil lamban:", mobilLamban.KecepatanPerJam)
}
