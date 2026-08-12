package main 

import "fmt"

func main(){
	var rupiah int
	
	fmt.Print("Masukkan nilai mata uang dalam rupiah : ")
	fmt.Scan(&rupiah)
	
	var dollar int = rupiah / 15000
	
	fmt.Println("mata uang anda telah dikonversi dari ", rupiah, "Rp ke", dollar, "USD", )
}