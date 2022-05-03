package main

import "fmt"

func main() {
	map1 := map[string]int{
		"zero":  0,
		"one":   1,
		"two":   2,
		"three": 3,
	}
	for key, _ := range map1 {
		fmt.Println(map1[key])
	}

	//var m map[string]int // ошибка, не проициализирована

	var m map[string]int = make(map[string]int, 10)
	m["one"] = 1
	fmt.Println(m)
}
