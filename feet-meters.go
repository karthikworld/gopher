package main

import "fmt"

func main() {
	
	fmt.Println("Enter Feet")
	var feet float64
	fmt.Scanf("%f", &feet)
	c := (feet * 0.3048)
	fmt.Println(c)
}