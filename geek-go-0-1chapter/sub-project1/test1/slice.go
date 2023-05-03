package test

import "fmt"

func s() {
	var a []int = make([]int, 0, 10)
	a = append(a, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println(a)
	b := a[2:5]
	fmt.Printf("%v,%v,%d\n", a, b, b[0])
	a = append(a, 11)
	b[0] = 100
	fmt.Printf("%v,%v,%d", a, b, b[0])
}
