package main

import "fmt"

func inc(i *int) {
	*i++
}

func zerofy(sl []int) {
	for i := range sl {
		sl[i] = 1
	}
	sl = append(sl, 4)
	fmt.Println(sl)
}
func main() {
	var x int = 0
	inc(&x)
	fmt.Println(x)

	sl := make([]int, 10)
	fmt.Println(sl)
	zerofy(sl)
	fmt.Println(sl)

}
