package main

import "fmt"

//Слайсы

func main() {
	array := [6]float64{1, 2, 3, 4, 5, 6}

	slice := array[0:1]
	//slice := array[0:1]
	// info slice
	// ptr = <array>
	// len = 2
	// cap = 5
	fmt.Println(slice)
	fmt.Println("len(slice)", len(slice))
	fmt.Println("cap(slice)", cap(slice))
}
