package main

import "fmt"

func main(){
	var beratBadan,tinggiBadan,bmi float32
	
	fmt.Print("Masukkan berat badan anda: ")
	fmt.Scan(&beratBadan)
	
	fmt.Print("Masukkan tinggi badan anda: ")
	fmt.Scan(&tinggiBadan)
	
	bmi = beratBadan / (tinggiBadan * tinggiBadan)
	
	fmt.Printf("Nilai BMI anda adalah:%.2f \n", bmi)
}