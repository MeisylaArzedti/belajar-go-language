package main

import "fmt"

func main(){
	var suhu float32
	
	fmt.Print("Masukkan suhu anda:")
	fmt.Scan(&suhu)
	
	
	if suhu >= 37.5 {
		fmt.Println("PERINGATAN!:suhu tinggi\nSilahkan periksa ke klinik")
	}
}