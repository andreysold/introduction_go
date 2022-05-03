package main

import "fmt"

func display(slice []int) {
	//var pointer *int
	//if len(slice) > 0 {
	//	pointer = &slice[0]
	//}
	//fmt.Println("Slice Ptr = ", pointer)
	//fmt.Println(desc)
	fmt.Println("len = ", len(slice))
	fmt.Println("cap = ", cap(slice))
	//if len(slice) > 0 {
	//fmt.Println("Slice: ", slice)
	//}
}

func main() {
	slice := []int{}

	for i := 0; i < 6; i++ {
		var pointer *int
		if len(slice) > 0 {
			pointer = &slice[0]
		}
		fmt.Println("address = ", pointer, i)
		display(slice)
		slice = append(slice, i)
	}
	//display(slice)
	fmt.Println(slice)
	//display("\t[0]\n", slice)
	//slice = append(slice, 1)
	//display("\t[1]\n", slice)
	//slice = append(slice, 2)
	//display("\t[2]\n", slice)
	//slice = append(slice, 3)
	//display("\t[3]\n", slice)
	//slice = append(slice, 4)
	//display("\t[4]\n", slice)
	//slice = append(slice, 5)
	//display("\t[5]\n", slice)
	//display(slice)
	//fmt.Printf("address = %p\n", slice)
	//fmt.Println(slice)
	//fmt.Println("len = ", len(slice))
	//fmt.Println("cap = ", cap(slice))
	//
	//fmt.Println("From massiv to slice")

	//slice1 := array[:]
	//fmt.Println(slice1)
	//fmt.Println("len = ", len(slice1))
	//fmt.Println("cap = ", cap(slice1))
	//slice1 = append(slice1, 1)
	//fmt.Println(slice1)
	//fmt.Printf("address = %p\n", slice1)
	//fmt.Println("len = ", len(slice1))
	//fmt.Println("cap = ", cap(slice1))
	//fmt.Println(array)
}
