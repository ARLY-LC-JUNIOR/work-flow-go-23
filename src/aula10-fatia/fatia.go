package main

import "fmt"

func main() {

	fatia1 := []int{1, 2, 3}
	fatia2 := make([]int, 2)
	copy(fatia2, fatia1)
	fmt.Println(fatia1, fatia2)
}

//var x []float64
//fatia:=make([]float64,4)
//fatia:=[low;hight]
//fatia:=arr[0:5]
//acrescentar elemento Append
//copiar uma fatia para outra
