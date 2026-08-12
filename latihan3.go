package main

import "fmt"

func main(){
	var x,y float32
	fmt.Print("Diketahui terdapat persamaan f(x)= 2x+5+5. \nMasukkan nilai x untuk persamaan tersebut: ")
	fmt.Scan(&x)
	
	y = 2*x+5+5
	
	fmt.Printf("hasil yang memenuhi persamaan tersebut adalah:%.2f\n",y )
}