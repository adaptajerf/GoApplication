package main

import "fmt"

func main() {
	fmt.Print("Hello World")
	var a [10]int
	for i := 0; i < 10; i++ {
		for j := 0; j < len(a); j++ {
			a[i] += i * j
		}
	}
	fmt.Println(a)
}
