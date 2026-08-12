package main

import "fmt"

func main(){
	fmt.Println("Halo! Nama saya Meisyla. Ini hari pertama saya latihan golang.")

	var (
		nama = "Meisyla"
		umur = 18
		kota = "Medan"
	)

	fmt.Printf("Nama: %v \nUmur: %v \nKota: %v\n", nama, umur, kota)

	tahunLahir := 2008
	umur = 2026 - tahunLahir
	fmt.Printf("Tahun lahir: %v \nUmur: %v\n", tahunLahir, umur)

	panjang := 10
	lebar := 5
	luas := panjang * lebar
	fmt.Printf("Diketahui luas persegi panjang dengan panjang %v dan lebar %v adalah %v\n", panjang, lebar, luas)

} 
