package main

import "fmt"

func main() {

	var x [5]float64
	x[0] = 5.3
	x[0] = 8
	x[0] = 4.2
	x[0] = 2.1
	x[0] = 7.8

	var total float64 = 0
	for i := 0; i < 5; i++ {
		total += x[i]

	}
	fmt.Println(total / 5)
}
