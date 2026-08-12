package main

import "fmt"

func main(){
	var nama,kelulusan string
	var rataRata,kehadiran float64
	
	fmt.Println("Pengecekan kelulusan siswa")
	
	fmt.Print("Masukkan nama siswa: ")
	fmt.Scan(&nama)
	
	fmt.Print("Masukkan rata-rata nilai siswa: ")
	fmt.Scan(&rataRata)
	
	fmt.Print("Masukkan nilai kehadiran siswa: ")
	fmt.Scan(&kehadiran)
	
	if rataRata >= 75 && kehadiran >= 80 {
		kelulusan = "lulus"
	} else {
		kelulusan = "tidak lulus"
	}
	
	fmt.Println("Ananda",nama,"dinyatakan",kelulusan)
	
	
}