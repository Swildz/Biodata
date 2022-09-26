package main

import (
	"fmt"
)

type DataDiri struct {
	id        int64
	nama      string
	alamat    string
	pekerjaan string
	alasan    string
}

func main() {
	biodata()

}
func biodata() {
	var data = []DataDiri{
		{id: 1, nama: "Siddiq", alamat: "Silaping", pekerjaan: "Mahasiswa", alasan: "Menambah pengalaman"},
		{id: 2, nama: "Dandi Irwanto", alamat: "Jawa Timur", pekerjaan: "Mahasiswa", alasan: "Menambah jam terbang"},
		{id: 3, nama: "Charles Prabowo", alamat: "Jawa Barat", pekerjaan: "Software Enginering", alasan: "Menambah relasi"},
		{id: 4, nama: "Maulana Dwi Wahyudi", alamat: "Jakarta", pekerjaan: "Software Developer", alasan: "Meningkatkan Pengetahuan"},
		{id: 5, nama: "Musthafa Kamal Faisal", alamat: "Sulawesi", pekerjaan: "Enginerr", alasan: "Mengisi waktu luang"},
	}

	for _, v := range data {
		fmt.Printf("Masukkan Nomor Absen : ")
		fmt.Scan(&v.id)
		switch v.id {
		case 1:
			fmt.Printf("%+v\n", v)
			break
		case 2:
			fmt.Printf("%+v\n", v)
			break
		case 3:
			fmt.Printf("%+v\n", v)
			break
		case 4:
			fmt.Printf("%+v\n", v)
		case 5:
			fmt.Printf("%+v\n", v)
			break
		default:
			fmt.Printf("Kembali ke awal")
		}
	}

}
