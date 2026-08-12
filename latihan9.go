package main

import "fmt"

func main(){
	var nilai int
	var status string
	
	fmt.Print("Masukkan nilai siswa: ")
	fmt.Scan(&nilai)
	
	if nilai >= 85 {
		status = "pujian"
	}else if nilai >= 70 {
		status = "baik"
	}else if nilai >= 55 {
		status = "cukup"
	}else { 
	status = "tidak lulus"
	}
	
	fmt.Println(status)
	
}