package main

import "fmt"

func main(){
	var x,y,z int
	var status bool
	
	fmt.Print("Masukkan tiga digit angka: ")
	fmt.Scan(&x,&y,&z)
	
	status = x < y && y < z 
	
	fmt.Println("Apakah tiga digit tersebut terurut membesar?\n",status )
}