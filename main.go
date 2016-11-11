package main

import (
	"fmt"
)

func main() {
	m := make(map[int]int)
	m[1] = 1
	m[2] = 2
	m1 := map[int]int{m}
	m[3] = 3
	fmt.Println(m)
	fmt.Println(m1)
}
