package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5, 6}
	a = append(a[:3])
	fmt.Println(a)
}
