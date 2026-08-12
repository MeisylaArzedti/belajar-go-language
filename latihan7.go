package main

import "fmt"

func main(){
	var x,y int
	var status string
	fmt.Scan(&x, &y)
	
	if x > 0 && y > 0 {
		status = "Kuadran I"
	}else if x < 0 && y > 0 {
		status = "Kuadran II"
	}else if x < 0 && y > 0 {
		status = "Kuadran III"
	}else if x > 0 && y < 0 {
		status = "Kuadran IV"
	} else {
		status = "Pada sumbu"
	}
	
	fmt.Println(status)
}