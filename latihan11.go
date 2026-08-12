package main

import "fmt"

func main(){
	var kode string
	fmt.Scan(&kode)
	
	switch kode{
		case "A" :
		fmt.Print("ojek motor, Rp10.000")
		case "B" :
		fmt.Print("mobil ekonomi,Rp25.000")
		case "C" :
		fmt.Print("mobil premium, Rp40.000")
		case "D" :
		fmt.Print("kirim barang, Rp15.000")
		case "E" :
		fmt.Print("Antar makanan, Rp12.000")
		default :
		fmt.Print("kode tidak dikenali")
	}
}