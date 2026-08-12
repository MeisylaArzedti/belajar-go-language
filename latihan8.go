package main 

import "fmt"

func main(){
	var a,b,c,d,e int
	var status string
	
	fmt.Scan(&a, &b, &c, &d, &e)
	
	if a > b && b > c && c > d && d > e{
		status = "Stabil turun"
	} else if a < b && b < c && c < d && d < e{
		status = "Stabil naik"
	} else {
		status = "Tidak stabil"
	}
	
	fmt.Println(status)
}