package main

import "fmt"

var x int
var y float64

func main() {
	x = 10
	y = 10.9
	fmt.Println("x=", x)
	fmt.Println("x=", y)
	fmt.Printf("x= %d  y= %g", x, y)
}
