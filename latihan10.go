package main

import "fmt"

func main(){
	var x,y,selisih float32
	var status string
	
	fmt.Scan(&x,&y)
	
	if x > y{
		status = "peningkatan sebesar"
		selisih = x - y
	}else if y > x{
		status = "penurunan sebesar"
		selisih = y - x
	}else{
		status = "tetap"
		selisih = 0
	}
	
	fmt.Printf("%s %.2f\n",status,selisih)
	
}